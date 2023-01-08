package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5"
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

func (r *PostgresRepository) SaveCredentialsData(ctx context.Context, creds *model.CredentialsData, userID int) error {
	return nil
	//insertBatch := saveMetadata(creds.Metadata)
	//query := `
	//	INSERT INTO credentials_data(name, login, password, user_id)
	//	VALUES ($1, $2, $3, $4);
	//	INSERT INTO data(user_id, credentials_data_id)
	//	VALUES ($4, currval(pg_get_serial_sequence('credentials_data', 'id')));
	//` + insertBatch
	//tx, err := r.Conn.BeginTx(ctx, nil)
	//if err != nil {
	//	return err
	//}
	//_, err = tx.ExecContext(ctx, query, creds.Name, creds.Login, creds.Password, userID)
	//if err != nil {
	//	if rollbackErr := tx.Rollback(); rollbackErr != nil {
	//		log.Error("unable to rollback transaction: %v", rollbackErr)
	//		return rollbackErr
	//	}
	//	return err
	//}
	//
	//err = tx.Commit()
	//if err != nil {
	//	if rollbackErr := tx.Rollback(); rollbackErr != nil {
	//		log.Error("unable to rollback transaction: %v", rollbackErr)
	//		return rollbackErr
	//	}
	//	return err
	//}
	//return nil
}

func (r *PostgresRepository) SaveBankingCardData(ctx context.Context, data *model.BankingCardData, userID int) error {
	//insertBatch := saveMetadata(data.Metadata)
	//query := `
	//	INSERT INTO banking_cards_data(user_id,
	//	                               name,
	//	                               number,
	//	                               valid_till,
	//	                               cardholder_name,
	//	                               cvv)
	//	VALUES ($1, $2, $3, $4, $5, $6);
	//	INSERT INTO data(user_id,banking_cards_data_id)
	//	VALUES ($1, currval(pg_get_serial_sequence('banking_cards_data', 'id')));
	//` + insertBatch
	//tx, err := r.Conn.BeginTx(ctx, nil)
	//if err != nil {
	//	return err
	//}
	//_, err = tx.ExecContext(ctx,
	//	query,
	//	userID,
	//	data.Name,
	//	data.Number,
	//	data.ValidUntil,
	//	data.CardholderName,
	//	data.CVV)
	//if err != nil {
	//	if rollbackErr := tx.Rollback(); rollbackErr != nil {
	//		log.Error("unable to rollback transaction: %v", rollbackErr)
	//		return rollbackErr
	//	}
	//	return err
	//}
	//err = tx.Commit()
	//if err != nil {
	//	if rollbackErr := tx.Rollback(); rollbackErr != nil {
	//		log.Error("unable to rollback transaction: %v", rollbackErr)
	//		return rollbackErr
	//	}
	//	return err
	//}
	return nil
}

func (r *PostgresRepository) SaveTextData(ctx context.Context, data *model.TextData, userID int) error {
	//insertBatch := saveMetadata(data.Metadata)
	//query := `
	//	INSERT INTO text_data(user_id, name, data)
	//	VALUES ($1,$2,$3);
	//	INSERT INTO data(user_id, text_data_id)
	//	VALUES ($1, currval(pg_get_serial_sequence('text_data', 'id')));
	//
	//` + insertBatch
	//tx, err := r.Conn.BeginTx(ctx, nil)
	//if err != nil {
	//	return err
	//}
	//
	//_, err = tx.ExecContext(ctx, query, userID, data.Name, data.Data)
	//if err != nil {
	//	if rollbackErr := tx.Rollback(); rollbackErr != nil {
	//		log.Error("unable to rollback transaction: %v", rollbackErr)
	//		return rollbackErr
	//	}
	//	return err
	//}
	//err = tx.Commit()
	//if err != nil {
	//	if rollbackErr := tx.Rollback(); rollbackErr != nil {
	//		log.Error("unable to rollback transaction: %v", rollbackErr)
	//		return rollbackErr
	//	}
	//	return err
	//}
	return nil
}

func (r *PostgresRepository) SaveBinaryData(ctx context.Context, data *model.BinaryData, userID int, link string) error {
	//query := `
	//	INSERT INTO binary_data(user_id, name, link)
	//	VALUES ($1, $2, $3);
	//	INSERT INTO data(user_id, binary_data_id)
	//	VALUES ($1, currval(pg_get_serial_sequence('binary_data', 'id')));
	//`
	//tx, err := r.Conn.BeginTx(ctx, nil)
	//if err != nil {
	//	return err
	//}
	//
	//_, err = tx.ExecContext(ctx, query, userID, data.Name, link)
	//if err != nil {
	//	if rollbackErr := tx.Rollback(); rollbackErr != nil {
	//		log.Error("unable to rollback transaction: %v", rollbackErr)
	//		return rollbackErr
	//	}
	//	return err
	//}
	//
	//err = tx.Commit()
	//if err != nil {
	//	if rollbackErr := tx.Rollback(); rollbackErr != nil {
	//		log.Error("unable to rollback transaction: %v", rollbackErr)
	//		return rollbackErr
	//	}
	//	return err
	//}
	return nil
}

func (r *PostgresRepository) GetCredentialsByUserID(ctx context.Context, userID int) ([]*model.CredentialsData, error) {
	//query := `
	//	SELECT data.name,
	//	       bcd.cardholder_name,
	//	       bcd.number,
	//	       bcd.valid_till,
	//	       bcd.cvv,
	//			cd.login,
	//			cd.password,
	//			bd.link,
	//			m.meta
	//	FROM data
	//	    JOIN banking_cards_data bcd ON bcd.id = data.banking_cards_data_id
	//		JOIN credentials_data cd ON cd.id = data.credentials_data_id
	//		JOIN binary_data bd on bd.id = data.binary_data_id
	//		JOIN metadata m on data.id = m.data_id
	//	WHERE data.user_id = $1
	//`
	//rows, err := r.Conn.QueryContext(ctx, query, userID)
	//if err != nil {
	//	return nil, err
	//}
	//for rows.Next() {
	//	result := model.CredentialsData{}
	//	bankData := model.BankingCardData{}
	//	meta := model.Metadata{}
	//	binary := model.BinaryData{}
	//
	//	rows.Scan(&result.Name,
	//		&bankData.CardholderName,
	//		&bankData.Number,
	//		&bankData.ValidUntil,
	//		&bankData.CVV,
	//		&result.Login,
	//		&result.Password,
	//		&binary.Data,
	//		&meta.Key,
	//	)
	//}
	return nil, nil
}

func (r *PostgresRepository) GetCredentialsByName(ctx context.Context, name string, userID int) ([]*model.CredentialsData, error) {
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

func saveMetadata(metadata []model.Metadata) string {
	insertBatch := ""
	if len(metadata) != 0 {
		insertBatch = "INSERT INTO metadata(data_id, meta) VALUES "
		for idx, meta := range metadata {
			if idx == len(metadata)-1 {
				insertBatch = insertBatch + fmt.Sprintf("(currval(pg_get_serial_sequence('data', 'id')), '%s');", meta.Value)
			} else {
				insertBatch = insertBatch + fmt.Sprintf("(currval(pg_get_serial_sequence('data', 'id')), '%s'), \n", meta.Value)
			}
		}
	}
	return insertBatch
}
