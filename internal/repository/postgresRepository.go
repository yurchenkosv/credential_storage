package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v4"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"github.com/yurchenkosv/credential_storage/internal/model"
)

type PostgresRepository struct {
	Conn  *sqlx.DB
	DBURI string
}

func (r *PostgresRepository) Save() error {
	return nil
}

func NewPostgresRepo(dbURI string) (*PostgresRepository, error) {
	conn, err := sqlx.Connect("postgres", dbURI)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	conn.SetMaxOpenConns(100)
	conn.SetMaxIdleConns(5)

	return &PostgresRepository{
		Conn:  conn,
		DBURI: dbURI,
	}, nil
}

func (r *PostgresRepository) GetUser(ctx context.Context, user *model.User) (*model.User, error) {
	var (
		userID   *int
		userName string
	)
	query := `
		SELECT id,name FROM users WHERE username=$1 AND password=$2;
	`
	err := r.Conn.
		QueryRowContext(ctx, query, user.Username, user.Password).
		Scan(&userID, &userName)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Error(err)
		return user, err
	}
	user.ID = userID
	user.Name = userName
	return user, nil
}

func (r *PostgresRepository) SaveUser(ctx context.Context, user *model.User) error {
	query := `
		INSERT INTO users(
		name,
		username,
		password
		)
		VALUES($1, $2, $3);
	`
	_, err := r.Conn.ExecContext(ctx, query, user.Name, user.Username, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresRepository) SaveCredentialsData(ctx context.Context, creds *model.Credentials, userID int) error {
	query := `
		INSERT INTO credentials_data(login, password, user_id)
		VALUES ($1, $2, $3)
	`
	tx, err := r.Conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	res, err := tx.ExecContext(ctx, query, creds.Login, creds.Password, userID)
	if err != nil {
		tx.Rollback()
		return err
	}
	credsId, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}
	query = `
		INSERT INTO data(user_id, credentials_data_id, name) 
		VALUES ($1, $2, $3)
	`
	res, err = tx.ExecContext(ctx, query, userID, credsId, creds.Name)
	if err != nil {
		tx.Rollback()
		return err
	}
	dataID, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}
	err = saveMetadata(ctx, tx, creds.Metadata, dataID)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
func (r *PostgresRepository) SaveBankingCardData(ctx context.Context, creds *model.Credentials, userID int) error {
	query := `
		INSERT INTO banking_cards_data(user_id,
		                               number,
		                               valid_till,
		                               cardholder_name,
		                               cvv)
		VALUES ($1, $2, $3, $4, $5)
	`
	tx, err := r.Conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	res, err := tx.ExecContext(ctx,
		query,
		userID,
		creds.Card.Number,
		creds.Card.ValidUntil,
		creds.Card.CardholderName,
		creds.Card.CVV)
	if err != nil {
		tx.Rollback()
		return err
	}
	bankingCardID, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}
	query = `
		INSERT INTO data(user_id, banking_cards_data_id, name)
		VALUES ($1, $2, $3)
	`
	res, err = tx.ExecContext(ctx, query, userID, bankingCardID, creds.Name)
	if err != nil {
		tx.Rollback()
		return err
	}
	recordID, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}
	err = saveMetadata(ctx, tx, creds.Metadata, recordID)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
func (r *PostgresRepository) SaveTextData(ctx context.Context, creds *model.Credentials, userID int) error {
	query := `
		INSERT INTO text_data(user_id, data) 
		VALUES ($1,$2)
	`
	tx, err := r.Conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	res, err := tx.ExecContext(ctx, query, userID, creds.Text.Data)
	if err != nil {
		tx.Rollback()
		return err
	}
	textID, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}
	query = `
		INSERT INTO data(name, user_id, text_data_id) 
		VALUES ($1, $2, $3)
	`
	res, err = tx.ExecContext(ctx, query, creds.Name, userID, textID)
	if err != nil {
		tx.Rollback()
		return err
	}
	dataID, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}
	err = saveMetadata(ctx, tx, creds.Metadata, dataID)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
func (r *PostgresRepository) SaveBinaryData(ctx context.Context, creds *model.Credentials, userID int, link string) error {
	query := `
		INSERT INTO binary_data(user_id, link) VALUES ($1, $2)
	`
	tx, err := r.Conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	res, err := tx.ExecContext(ctx, query, userID, link)
	if err != nil {
		tx.Rollback()
		return err
	}
	binaryID, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}
	query = `
		INSERT INTO data(name, user_id, binary_data_id) 
		VALUES ($1, $2, $3)
	`
	res, err = tx.ExecContext(ctx, query, creds.Name, userID, binaryID)
	dataID, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}
	err = saveMetadata(ctx, tx, creds.Metadata, dataID)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (r *PostgresRepository) GetCredentialsByUserID(ctx context.Context, userID int) (*model.Credentials, error) {
	return nil, nil
}
func (r *PostgresRepository) GetCredentialsByName(ctx context.Context, name string) (*model.Credentials, error) {
	return nil, nil
}

func (r *PostgresRepository) MigrateDB(migrationsPath string) error {
	m, err := migrate.New(
		"file://"+migrationsPath,
		r.DBURI)
	if err != nil {
		return err
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

func saveMetadata(ctx context.Context, tx *sql.Tx, metadata []model.Metadata, dataID int64) error {
	query := `
		INSERT INTO metadata(data_id, meta) VALUES ($1, $2)
	`
	for _, meta := range metadata {
		_, err := tx.ExecContext(ctx, query, dataID, meta.Data)
		if err != nil {
			return err
		}
	}
	return nil
}
