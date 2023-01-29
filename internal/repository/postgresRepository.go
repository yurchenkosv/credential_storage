package repository

import (
	"context"
	"database/sql"
	"errors"
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
	err := r.Transactional(ctx, func() error {
		query := `
			INSERT INTO credentials_data(login, password, user_id)
			VALUES ($1, $2, $3);
		`
		_, err := r.Conn.ExecContext(ctx, query, creds.Login, creds.Password, userID)
		if err != nil {
			return err
		}
		query = `
			INSERT INTO data(name, user_id, credentials_data_id)
			VALUES ($1, $2, currval(pg_get_serial_sequence('credentials_data', 'id')));
		`
		_, err = r.Conn.ExecContext(ctx, query, creds.Name, userID)
		if err != nil {
			return err
		}
		err = r.saveMetadata(ctx, creds.Metadata)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (r *PostgresRepository) SaveBankingCardData(ctx context.Context, data *model.BankingCardData, userID int) error {
	err := r.Transactional(ctx, func() error {
		query := `
			INSERT INTO banking_cards_data(user_id,
		                               number,
		                               valid_till,
		                               cardholder_name,
		                               cvv)
			VALUES ($1, $2, $3, $4, $5);
		`
		_, err := r.Conn.ExecContext(ctx,
			query,
			userID,
			data.Number,
			data.ValidUntil,
			data.CardholderName,
			data.CVV)
		if err != nil {
			return err
		}
		query = `
			INSERT INTO data(name, user_id, banking_cards_data_id)
			VALUES ($1, $2, currval(pg_get_serial_sequence('banking_cards_data', 'id')));
		`
		_, err = r.Conn.ExecContext(ctx, query, data.Name, userID)
		if err != nil {
			return err
		}
		err = r.saveMetadata(ctx, data.Metadata)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (r *PostgresRepository) SaveTextData(ctx context.Context, data *model.TextData, userID int) error {
	err := r.Transactional(ctx, func() error {
		query := `
			INSERT INTO text_data(user_id, data)
			VALUES ($1, $2);
		`
		_, err := r.Conn.ExecContext(ctx, query, userID, data.Data)
		if err != nil {
			return err
		}
		query = `
			INSERT INTO data(user_id, text_data_id)
			VALUES ($1, $2, currval(pg_get_serial_sequence('text_data', 'id')));
		`
		_, err = r.Conn.ExecContext(ctx, query, data.Name, userID)
		if err != nil {
			return err
		}
		err = r.saveMetadata(ctx, data.Metadata)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (r *PostgresRepository) SaveBinaryData(ctx context.Context, data *model.BinaryData, userID int, link string) error {

	err := r.Transactional(ctx, func() error {
		query := `
			INSERT INTO binary_data(user_id, link)
			VALUES ($1, $2);
		`
		_, err := r.Conn.ExecContext(ctx, query, userID, link)
		if err != nil {
			return err
		}
		query = `
			INSERT INTO data(name, user_id, binary_data_id)
			VALUES ($1, $2, currval(pg_get_serial_sequence('binary_data', 'id')));		
		`
		_, err = r.Conn.ExecContext(ctx, query, data.Name, userID)
		if err != nil {
			return err
		}
		err = r.saveMetadata(ctx, data.Metadata)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (r *PostgresRepository) GetCredentialsByUserID(ctx context.Context, userID int) ([]model.Credentials, error) {
	query := `
		SELECT data.name, cd.login, cd.password,
			   bcd.cardholder_name, bcd.number, bcd.valid_till, bcd.cvv,
			   td.data, data.id
-- 			   bd.link
		FROM data
				FULL JOIN credentials_data cd ON data.credentials_data_id = cd.id
				FULL JOIN banking_cards_data bcd ON bcd.id = data.banking_cards_data_id
				FULL JOIN text_data td ON td.id = data.text_data_id
-- 				FULL JOIN binary_data bd ON bd.id = data.binary_data_id
		WHERE data.user_id=$1;
	`
	rows, err := r.Conn.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	credentials := []model.Credentials{}
	for rows.Next() {
		var id int
		secrets := model.CredentialsData{}
		bankData := model.BankingCardData{}
		textData := model.TextData{}
		cred := model.Credentials{}
		err = rows.Scan(
			&cred.Name,
			&secrets.Login,
			&secrets.Password,
			&bankData.CardholderName,
			&bankData.Number,
			&bankData.ValidUntil,
			&bankData.CVV,
			&textData.Data,
			&id,
		)
		if err != nil {
			continue
		}

		qry := `
			SELECT meta FROM metadata WHERE data_id=$1;	
		`
		rows, err = r.Conn.QueryContext(ctx, qry, id)
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			meta := model.Metadata{}
			err = rows.Scan(&meta.Value)
			if err != nil {
				continue
			}
			cred.Metadata = append(cred.Metadata, meta)
		}

		credentials = append(credentials, cred)
	}
	return credentials, nil
}

func (r *PostgresRepository) GetCredentialsByName(ctx context.Context, name string, userID int) ([]model.CredentialsData, error) {
	return nil, errors.New("not implemented")
}

func (r *PostgresRepository) UpdateBankingCardData(ctx context.Context, data model.Credentials, userID int) error {
	err := r.Transactional(ctx, func() error {
		query := `
			UPDATE data SET name=$1
			WHERE user_id=$2 AND banking_cards_data_id=$3;
		`
		_, err := r.Conn.ExecContext(ctx, query, data.Name, userID, data.BankingCardData.ID)
		if err != nil {
			return err
		}
		query = `
			UPDATE banking_cards_data SET cvv=$1, valid_till=$2, number=$3, cardholder_name=$4
			FROM banking_cards_data
				JOIN data d ON banking_cards_data.id = d.banking_cards_data_id
			WHERE d.name=$5 AND d.user_id=$6;
		`
		bankingData := data.BankingCardData
		_, err = r.Conn.ExecContext(ctx,
			query,
			bankingData.CVV,
			bankingData.ValidUntil,
			bankingData.Number,
			bankingData.CardholderName,
			bankingData.Name,
			userID)
		if err != nil {
			return err
		}
		err = r.updateMetadata(ctx, data.Metadata, data.ID)
		return nil
	})
	return err
}

func (r *PostgresRepository) UpdateCredentialsData(ctx context.Context, data model.Credentials, userID int) error {
	err := r.Transactional(ctx, func() error {
		query := `
			UPDATE data SET name=$1
			WHERE user_id=$2 AND credentials_data_id=$3 
		`
		_, err := r.Conn.ExecContext(ctx, query, data.Name, userID, data.CredentialsData.ID)
		if err != nil {
			return err
		}
		credentialsData := data.CredentialsData
		query = `
			UPDATE credentials_data set login=$1, password=$2
			WHERE id=$3 AND user_id=$4
		`
		_, err = r.Conn.ExecContext(ctx,
			query,
			credentialsData.Login,
			credentialsData.Password,
			credentialsData.ID,
			userID)
		if err != nil {
			return err
		}
		err = r.updateMetadata(ctx, data.Metadata, data.ID)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (r *PostgresRepository) UpdateTextData(ctx context.Context, data model.Credentials, userID int) error {
	err := r.Transactional(ctx, func() error {
		query := `
			UPDATE data SET name=$1
			WHERE user_id=$2 AND credentials_data_id=$3 
		`
		_, err := r.Conn.ExecContext(ctx, query, data.Name, userID, data.CredentialsData.ID)
		if err != nil {
			return err
		}
		textData := data.TextData
		query = `
			UPDATE text_data SET data = $1
			WHERE id=$2 AND user_id=$3
		`
		_, err = r.Conn.ExecContext(ctx, query, textData.Data, textData.ID, userID)
		if err != nil {
			return err
		}
		err = r.updateMetadata(ctx, data.Metadata, data.ID)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (r *PostgresRepository) UpdateBinaryData(ctx context.Context, data model.Credentials, userID int) error {
	err := r.Transactional(ctx, func() error {
		query := `
			UPDATE data SET name=$1
			WHERE user_id=$2 AND credentials_data_id=$3 
		`
		_, err := r.Conn.ExecContext(ctx, query, data.Name, userID, data.CredentialsData.ID)
		if err != nil {
			return err
		}
		query = `
			UPDATE binary_data SET link=$1
			WHERE id=$2 AND user_id=$3
		`
		_, err = r.Conn.ExecContext(ctx, query, data.BinaryData.Link, data.BinaryData.ID, userID)
		if err != nil {
			return err
		}
		err = r.updateMetadata(ctx, data.Metadata, data.ID)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (r *PostgresRepository) DeleteData(ctx context.Context, data model.Credentials, userID int) error {
	err := r.Transactional(ctx, func() error {
		query := `
			DELETE FROM data WHERE id = $1 AND user_id=$2;
		`
		_, err := r.Conn.ExecContext(ctx, query, data.ID, userID)
		if err != nil {
			return err
		}
		query = `
			DELETE FROM metadata WHERE data_id=$1;
		`
		_, err = r.Conn.ExecContext(ctx, query, data.ID)
		if err != nil {
			return err
		}
		return nil
	})
	return err
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

func (r *PostgresRepository) Transactional(ctx context.Context, do func() error) error {
	tx, err := r.Conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	err = do()
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Error("unable to rollback transaction: %v", rollbackErr)
			return rollbackErr
		}
		return err
	}
	err = tx.Commit()
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Error("unable to rollback transaction: %v", rollbackErr)
			return rollbackErr
		}
		return err
	}
	return nil
}

func (r *PostgresRepository) saveMetadata(ctx context.Context, metadata []model.Metadata) error {
	if len(metadata) == 0 {
		return nil
	}
	query := `
		INSERT INTO metadata(data_id, meta)
		VALUES (currval(pg_get_serial_sequence('data', 'id')), $1);
	`
	for _, meta := range metadata {
		_, err := r.Conn.ExecContext(ctx, query, meta.Value)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *PostgresRepository) updateMetadata(ctx context.Context, metadata []model.Metadata, dataID int) error {
	if len(metadata) == 0 {
		return nil
	}
	query := `
		DELETE FROM metadata
		WHERE data_id=$1
	`
	_, err := r.Conn.ExecContext(ctx, query, dataID)
	if err != nil {
		return err
	}
	query = `
		INSERT INTO metadata(data_id, meta)
		VALUES($1, $2)
	`
	for _, meta := range metadata {
		_, err = r.Conn.ExecContext(ctx, query, dataID, meta.Value)
		if err != nil {
			log.Error("cannot insert metadata ", err)
		}
	}
	return nil
}
