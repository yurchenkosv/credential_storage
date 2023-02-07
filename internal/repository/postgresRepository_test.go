package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/yurchenkosv/credential_storage/internal/model"
	"reflect"
	"testing"
)

func TestNewPostgresRepo(t *testing.T) {
	type args struct {
		dbURI string
	}
	tests := []struct {
		name    string
		args    args
		want    *PostgresRepository
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPostgresRepo(tt.args.dbURI)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPostgresRepo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPostgresRepo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostgresRepository_DeleteData(t *testing.T) {
	type fields struct {
		Conn  *sqlx.DB
		DBURI string
	}
	type args struct {
		ctx    context.Context
		data   model.Credentials
		userID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &PostgresRepository{
				Conn:  tt.fields.Conn,
				DBURI: tt.fields.DBURI,
			}
			if err := r.DeleteData(tt.args.ctx, tt.args.data, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("DeleteData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostgresRepository_GetCredentialsByName(t *testing.T) {
	type fields struct {
		Conn  *sqlx.DB
		DBURI string
	}
	type args struct {
		ctx    context.Context
		name   string
		userID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.CredentialsData
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &PostgresRepository{
				Conn:  tt.fields.Conn,
				DBURI: tt.fields.DBURI,
			}
			got, err := r.GetCredentialsByName(tt.args.ctx, tt.args.name, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCredentialsByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCredentialsByName() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostgresRepository_GetCredentialsByUserID(t *testing.T) {
	type fields struct {
		Conn  *sqlx.DB
		DBURI string
	}
	type args struct {
		ctx    context.Context
		userID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.Credentials
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &PostgresRepository{
				Conn:  tt.fields.Conn,
				DBURI: tt.fields.DBURI,
			}
			got, err := r.GetCredentialsByUserID(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCredentialsByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCredentialsByUserID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostgresRepository_GetUser(t *testing.T) {
	type fields struct {
		Conn  *sqlx.DB
		DBURI string
	}
	type args struct {
		ctx  context.Context
		user *model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &PostgresRepository{
				Conn:  tt.fields.Conn,
				DBURI: tt.fields.DBURI,
			}
			got, err := r.GetUser(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostgresRepository_MigrateDB(t *testing.T) {
	type fields struct {
		Conn  *sqlx.DB
		DBURI string
	}
	type args struct {
		migrationsPath string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &PostgresRepository{
				Conn:  tt.fields.Conn,
				DBURI: tt.fields.DBURI,
			}
			if err := r.MigrateDB(tt.args.migrationsPath); (err != nil) != tt.wantErr {
				t.Errorf("MigrateDB() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostgresRepository_SaveBankingCardData(t *testing.T) {
	type fields struct {
		Conn  *sqlx.DB
		DBURI string
	}
	type args struct {
		ctx    context.Context
		data   *model.BankingCardData
		userID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &PostgresRepository{
				Conn:  tt.fields.Conn,
				DBURI: tt.fields.DBURI,
			}
			if err := r.SaveBankingCardData(tt.args.ctx, tt.args.data, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("SaveBankingCardData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostgresRepository_SaveBinaryData(t *testing.T) {
	type fields struct {
		Conn  *sqlx.DB
		DBURI string
	}
	type args struct {
		ctx    context.Context
		data   *model.BinaryData
		userID int
		link   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &PostgresRepository{
				Conn:  tt.fields.Conn,
				DBURI: tt.fields.DBURI,
			}
			if err := r.SaveBinaryData(tt.args.ctx, tt.args.data, tt.args.userID, tt.args.link); (err != nil) != tt.wantErr {
				t.Errorf("SaveBinaryData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostgresRepository_SaveCredentialsData(t *testing.T) {
	type fields struct {
		Conn  *sqlx.DB
		DBURI string
	}
	type args struct {
		ctx    context.Context
		creds  *model.CredentialsData
		userID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &PostgresRepository{
				Conn:  tt.fields.Conn,
				DBURI: tt.fields.DBURI,
			}
			if err := r.SaveCredentialsData(tt.args.ctx, tt.args.creds, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("SaveCredentialsData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostgresRepository_SaveTextData(t *testing.T) {
	type fields struct {
		Conn  *sqlx.DB
		DBURI string
	}
	type args struct {
		ctx    context.Context
		data   *model.TextData
		userID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &PostgresRepository{
				Conn:  tt.fields.Conn,
				DBURI: tt.fields.DBURI,
			}
			if err := r.SaveTextData(tt.args.ctx, tt.args.data, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("SaveTextData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostgresRepository_SaveUser(t *testing.T) {
	type fields struct {
		Conn  *sqlx.DB
		DBURI string
	}
	type args struct {
		ctx  context.Context
		user *model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &PostgresRepository{
				Conn:  tt.fields.Conn,
				DBURI: tt.fields.DBURI,
			}
			if err := r.SaveUser(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("SaveUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostgresRepository_Transactional(t *testing.T) {
	type fields struct {
		Conn  *sqlx.DB
		DBURI string
	}
	type args struct {
		ctx context.Context
		do  func() error
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &PostgresRepository{
				Conn:  tt.fields.Conn,
				DBURI: tt.fields.DBURI,
			}
			if err := r.Transactional(tt.args.ctx, tt.args.do); (err != nil) != tt.wantErr {
				t.Errorf("Transactional() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostgresRepository_UpdateBankingCardData(t *testing.T) {
	type fields struct {
		Conn  *sqlx.DB
		DBURI string
	}
	type args struct {
		ctx    context.Context
		data   model.Credentials
		userID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &PostgresRepository{
				Conn:  tt.fields.Conn,
				DBURI: tt.fields.DBURI,
			}
			if err := r.UpdateBankingCardData(tt.args.ctx, tt.args.data, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("UpdateBankingCardData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostgresRepository_UpdateBinaryData(t *testing.T) {
	type fields struct {
		Conn  *sqlx.DB
		DBURI string
	}
	type args struct {
		ctx    context.Context
		data   model.Credentials
		userID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &PostgresRepository{
				Conn:  tt.fields.Conn,
				DBURI: tt.fields.DBURI,
			}
			if err := r.UpdateBinaryData(tt.args.ctx, tt.args.data, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("UpdateBinaryData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostgresRepository_UpdateCredentialsData(t *testing.T) {
	type fields struct {
		Conn  *sqlx.DB
		DBURI string
	}
	type args struct {
		ctx    context.Context
		data   model.Credentials
		userID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &PostgresRepository{
				Conn:  tt.fields.Conn,
				DBURI: tt.fields.DBURI,
			}
			if err := r.UpdateCredentialsData(tt.args.ctx, tt.args.data, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("UpdateCredentialsData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostgresRepository_UpdateTextData(t *testing.T) {
	type fields struct {
		Conn  *sqlx.DB
		DBURI string
	}
	type args struct {
		ctx    context.Context
		data   model.Credentials
		userID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &PostgresRepository{
				Conn:  tt.fields.Conn,
				DBURI: tt.fields.DBURI,
			}
			if err := r.UpdateTextData(tt.args.ctx, tt.args.data, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("UpdateTextData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostgresRepository_saveMetadata(t *testing.T) {
	type fields struct {
		Conn  *sqlx.DB
		DBURI string
	}
	type args struct {
		ctx      context.Context
		metadata []model.Metadata
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &PostgresRepository{
				Conn:  tt.fields.Conn,
				DBURI: tt.fields.DBURI,
			}
			if err := r.saveMetadata(tt.args.ctx, tt.args.metadata); (err != nil) != tt.wantErr {
				t.Errorf("saveMetadata() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostgresRepository_updateMetadata(t *testing.T) {
	type fields struct {
		Conn  *sqlx.DB
		DBURI string
	}
	type args struct {
		ctx      context.Context
		metadata []model.Metadata
		dataID   int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &PostgresRepository{
				Conn:  tt.fields.Conn,
				DBURI: tt.fields.DBURI,
			}
			if err := r.updateMetadata(tt.args.ctx, tt.args.metadata, tt.args.dataID); (err != nil) != tt.wantErr {
				t.Errorf("updateMetadata() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
