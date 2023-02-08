package repository

import (
	"context"
	"fmt"
	"github.com/docker/go-connections/nat"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"github.com/yurchenkosv/credential_storage/internal/model"
	"reflect"
	"testing"
)

func initContainers(t *testing.T, ctx context.Context) testcontainers.Container {

	port, err := nat.NewPort("tcp", "5432")
	if err != nil {
		t.Error(err)
	}

	req := testcontainers.ContainerRequest{
		Image:        "postgres:14",
		ExposedPorts: []string{port.Port() + "/tcp"},
		Env: map[string]string{
			"POSTGRES_PASSWORD": "postgres",
			"POSTGRES_DB":       "goph_keeper",
		},
		WaitingFor: wait.ForListeningPort(port),
		AutoRemove: true,
	}
	postgres, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Error(err)
	}
	return postgres
}

func initDatabase(repo *PostgresRepository, data model.CredentialsData, userID int) {
	ctx := context.Background()
	err := repo.MigrateDB("../migrations")
	if err != nil {
		log.Error(err)
	}
	log.Info("successfully migrate")
	err = repo.SaveCredentialsData(ctx, &data, userID)
	if err != nil {
		log.Error(err)
	}
}

func intPtr(val int) *int {
	return &val
}

func TestNewPostgresRepo(t *testing.T) {
	type args struct {
		dbURI string
	}
	tests := []struct {
		name    string
		args    args
		before  func(t *testing.T, ctx context.Context) testcontainers.Container
		want    *PostgresRepository
		wantErr bool
	}{
		{
			name:    "should success with creating repository",
			args:    args{},
			before:  initContainers,
			want:    &PostgresRepository{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			postgres := tt.before(t, ctx)
			defer func() {
				if err := postgres.Terminate(ctx); err != nil {
					t.Fatalf("failed to terminate container: %s", err.Error())
				}
			}()
			endpoint, err := postgres.Endpoint(ctx, "")
			if err != nil {
				t.Error(err)
			}
			tt.args.dbURI = fmt.Sprintf("postgresql://postgres:postgres@%s/goph_keeper?sslmode=disable", endpoint)
			got, err := NewPostgresRepo(tt.args.dbURI)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPostgresRepo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.IsType(t, tt.want, got)
		})
	}
}

func TestPostgresRepository_DeleteData(t *testing.T) {
	type args struct {
		ctx    context.Context
		data   model.Credentials
		userID int
	}
	tests := []struct {
		name            string
		before          func(t *testing.T, ctx context.Context) testcontainers.Container
		beforeCondition func(repo *PostgresRepository, data model.CredentialsData, userID int)
		args            args
		wantErr         bool
	}{
		{
			name:            "should sucessfuly delete data",
			before:          initContainers,
			beforeCondition: initDatabase,
			args: args{
				ctx:    context.Background(),
				data:   model.Credentials{},
				userID: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			postgres := tt.before(t, ctx)
			defer func() {
				if err := postgres.Terminate(ctx); err != nil {
					t.Fatalf("failed to terminate container: %s", err.Error())
				}
			}()
			endpoint, err := postgres.Endpoint(ctx, "")
			if err != nil {
				t.Error(err)
			}
			DBURI := fmt.Sprintf("postgresql://postgres:postgres@%s/goph_keeper?sslmode=disable", endpoint)
			r, _ := NewPostgresRepo(DBURI)
			credData := model.CredentialsData{
				Name:     "testcreds",
				Login:    "test",
				Password: "test",
			}
			tt.beforeCondition(r, credData, tt.args.userID)
			if err = r.DeleteData(tt.args.ctx, tt.args.data, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("DeleteData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostgresRepository_GetCredentialsByUserID(t *testing.T) {
	type args struct {
		ctx         context.Context
		credentials model.CredentialsData
		userID      int
	}
	tests := []struct {
		name            string
		before          func(t *testing.T, ctx context.Context) testcontainers.Container
		beforeCondition func(repo *PostgresRepository, data model.CredentialsData, userID int)
		args            args
		want            []model.Credentials
		wantErr         bool
	}{
		{
			name:            "should sucessfuly return credentials",
			before:          initContainers,
			beforeCondition: initDatabase,
			args: args{
				ctx: context.Background(),
				credentials: model.CredentialsData{
					ID:       1,
					Name:     "test_creds",
					Login:    "test_login",
					Password: "test_pwd",
				},
				userID: 1,
			},
			want: []model.Credentials{
				{
					ID:   1,
					Name: "test_creds",
					CredentialsData: &model.CredentialsData{
						ID:       1,
						Login:    "test_login",
						Password: "test_pwd",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			postgres := tt.before(t, ctx)
			defer func() {
				if err := postgres.Terminate(ctx); err != nil {
					t.Fatalf("failed to terminate container: %s", err.Error())
				}
			}()
			endpoint, err := postgres.Endpoint(ctx, "")
			if err != nil {
				t.Error(err)
			}
			DBURI := fmt.Sprintf("postgresql://postgres:postgres@%s/goph_keeper?sslmode=disable", endpoint)
			r, _ := NewPostgresRepo(DBURI)
			tt.beforeCondition(r, tt.args.credentials, tt.args.userID)
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
	type beforeCondition func(repo *PostgresRepository, data model.User)
	type args struct {
		ctx  context.Context
		user *model.User
	}
	tests := []struct {
		name            string
		before          func(t *testing.T, ctx context.Context) testcontainers.Container
		beforeCondition beforeCondition
		args            args
		want            *model.User
		wantErr         bool
	}{
		{
			name:   "should successfully get user ",
			before: initContainers,
			beforeCondition: beforeCondition(func(repo *PostgresRepository, data model.User) {
				ctx := context.Background()
				err := repo.MigrateDB("../migrations")
				if err != nil {
					log.Error(err)
				}
				log.Info("successfully migrate")
				err = repo.SaveUser(ctx, &data)
				if err != nil {
					log.Error(err)
				}
			}),
			args: args{
				ctx: context.Background(),
				user: &model.User{
					Username: "test_user",
					Password: "test_password",
					Name:     "testName",
				},
			},
			want: &model.User{
				ID:       intPtr(1),
				Username: "test_user",
				Password: "test_password",
				Name:     "testName",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			postgres := tt.before(t, ctx)
			defer func() {
				if err := postgres.Terminate(ctx); err != nil {
					t.Fatalf("failed to terminate container: %s", err.Error())
				}
			}()
			endpoint, err := postgres.Endpoint(ctx, "")
			if err != nil {
				t.Error(err)
			}
			DBURI := fmt.Sprintf("postgresql://postgres:postgres@%s/goph_keeper?sslmode=disable", endpoint)
			r, _ := NewPostgresRepo(DBURI)
			tt.beforeCondition(r, *tt.args.user)
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
	type args struct {
		migrationsPath string
	}
	tests := []struct {
		name    string
		before  func(t *testing.T, ctx context.Context) testcontainers.Container
		args    args
		wantErr bool
	}{
		{
			name:   "should successfully apply migrations on DB",
			before: initContainers,
			args: args{
				migrationsPath: "../migrations",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			postgres := tt.before(t, ctx)
			defer func() {
				if err := postgres.Terminate(ctx); err != nil {
					t.Fatalf("failed to terminate container: %s", err.Error())
				}
			}()
			endpoint, err := postgres.Endpoint(ctx, "")
			if err != nil {
				t.Error(err)
			}
			DBURI := fmt.Sprintf("postgresql://postgres:postgres@%s/goph_keeper?sslmode=disable", endpoint)
			r, _ := NewPostgresRepo(DBURI)
			if err = r.MigrateDB(tt.args.migrationsPath); (err != nil) != tt.wantErr {
				t.Errorf("MigrateDB() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostgresRepository_SaveBankingCardData(t *testing.T) {
	type args struct {
		ctx    context.Context
		data   *model.BankingCardData
		userID int
	}
	tests := []struct {
		name    string
		before  func(t *testing.T, ctx context.Context) testcontainers.Container
		args    args
		wantErr bool
	}{
		{
			name:   "should successfully save bank data",
			before: initContainers,
			args: args{
				ctx: context.Background(),
				data: &model.BankingCardData{
					Name:           "test_data",
					Number:         "198740198274",
					ValidUntil:     "10/24",
					CardholderName: "test user",
					CVV:            "124",
				},
				userID: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			postgres := tt.before(t, ctx)
			defer func() {
				if err := postgres.Terminate(ctx); err != nil {
					t.Fatalf("failed to terminate container: %s", err.Error())
				}
			}()
			endpoint, err := postgres.Endpoint(ctx, "")
			if err != nil {
				t.Error(err)
			}
			DBURI := fmt.Sprintf("postgresql://postgres:postgres@%s/goph_keeper?sslmode=disable", endpoint)
			r, _ := NewPostgresRepo(DBURI)
			r.MigrateDB("../migrations")
			if err = r.SaveBankingCardData(tt.args.ctx, tt.args.data, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("SaveBankingCardData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostgresRepository_SaveBinaryData(t *testing.T) {
	type args struct {
		ctx    context.Context
		data   *model.BinaryData
		userID int
		link   string
	}
	tests := []struct {
		name    string
		before  func(t *testing.T, ctx context.Context) testcontainers.Container
		args    args
		wantErr bool
	}{
		{
			name:   "should successfully create binary record",
			before: initContainers,
			args: args{
				ctx: context.Background(),
				data: &model.BinaryData{
					Name:     "test_binary",
					Metadata: nil,
				},
				userID: 1,
				link:   "/temp/path/18478347",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			postgres := tt.before(t, ctx)
			defer func() {
				if err := postgres.Terminate(ctx); err != nil {
					t.Fatalf("failed to terminate container: %s", err.Error())
				}
			}()
			endpoint, err := postgres.Endpoint(ctx, "")
			if err != nil {
				t.Error(err)
			}
			DBURI := fmt.Sprintf("postgresql://postgres:postgres@%s/goph_keeper?sslmode=disable", endpoint)
			r, _ := NewPostgresRepo(DBURI)
			r.MigrateDB("../migrations")
			if err = r.SaveBinaryData(tt.args.ctx, tt.args.data, tt.args.userID, tt.args.link); (err != nil) != tt.wantErr {
				t.Errorf("SaveBinaryData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostgresRepository_SaveCredentialsData(t *testing.T) {
	type args struct {
		ctx    context.Context
		creds  *model.CredentialsData
		userID int
	}
	tests := []struct {
		name    string
		before  func(t *testing.T, ctx context.Context) testcontainers.Container
		args    args
		wantErr bool
	}{
		{
			name:   "should successfully save credentials",
			before: initContainers,
			args: args{
				ctx: context.Background(),
				creds: &model.CredentialsData{
					Name:     "test_creds",
					Login:    "test",
					Password: "test_pwd",
				},
				userID: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			postgres := tt.before(t, ctx)
			defer func() {
				if err := postgres.Terminate(ctx); err != nil {
					t.Fatalf("failed to terminate container: %s", err.Error())
				}
			}()
			endpoint, err := postgres.Endpoint(ctx, "")
			if err != nil {
				t.Error(err)
			}
			DBURI := fmt.Sprintf("postgresql://postgres:postgres@%s/goph_keeper?sslmode=disable", endpoint)
			r, _ := NewPostgresRepo(DBURI)
			r.MigrateDB("../migrations")
			if err = r.SaveCredentialsData(tt.args.ctx, tt.args.creds, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("SaveCredentialsData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostgresRepository_SaveTextData(t *testing.T) {
	type args struct {
		ctx    context.Context
		data   *model.TextData
		userID int
	}
	tests := []struct {
		name    string
		before  func(t *testing.T, ctx context.Context) testcontainers.Container
		args    args
		wantErr bool
	}{
		{
			name:   "should successfuly save text data",
			before: initContainers,
			args: args{
				ctx: context.Background(),
				data: &model.TextData{
					Name: "test_text",
					Data: "texttexttexttexttext",
				},
				userID: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			postgres := tt.before(t, ctx)
			defer func() {
				if err := postgres.Terminate(ctx); err != nil {
					t.Fatalf("failed to terminate container: %s", err.Error())
				}
			}()
			endpoint, err := postgres.Endpoint(ctx, "")
			if err != nil {
				t.Error(err)
			}
			DBURI := fmt.Sprintf("postgresql://postgres:postgres@%s/goph_keeper?sslmode=disable", endpoint)
			r, _ := NewPostgresRepo(DBURI)
			r.MigrateDB("../migrations")
			if err = r.SaveTextData(tt.args.ctx, tt.args.data, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("SaveTextData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostgresRepository_SaveUser(t *testing.T) {
	type args struct {
		ctx  context.Context
		user *model.User
	}
	tests := []struct {
		name    string
		before  func(t *testing.T, ctx context.Context) testcontainers.Container
		args    args
		wantErr bool
	}{
		{
			name:   "should successfully save user",
			before: initContainers,
			args: args{
				ctx: context.Background(),
				user: &model.User{
					Username: "test_username",
					Password: "test_password",
					Name:     "Test Name",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			postgres := tt.before(t, ctx)
			defer func() {
				if err := postgres.Terminate(ctx); err != nil {
					t.Fatalf("failed to terminate container: %s", err.Error())
				}
			}()
			endpoint, err := postgres.Endpoint(ctx, "")
			if err != nil {
				t.Error(err)
			}
			DBURI := fmt.Sprintf("postgresql://postgres:postgres@%s/goph_keeper?sslmode=disable", endpoint)
			r, _ := NewPostgresRepo(DBURI)
			r.MigrateDB("../migrations")
			if err = r.SaveUser(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
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
	type beforeCondition func(repo *PostgresRepository, ctx context.Context, data model.BankingCardData, userID int)
	type args struct {
		ctx        context.Context
		data       model.BankingCardData
		beforeData model.BankingCardData
		userID     int
	}
	tests := []struct {
		name            string
		before          func(t *testing.T, ctx context.Context) testcontainers.Container
		beforeCondition beforeCondition
		args            args
		wantErr         bool
	}{
		{
			name:   "should update banking card data",
			before: initContainers,
			beforeCondition: func(repo *PostgresRepository, ctx context.Context, data model.BankingCardData, userID int) {
				err := repo.MigrateDB("../migrations")
				if err != nil {
					log.Error(err)
				}
				err = repo.SaveBankingCardData(ctx, &data, userID)
				if err != nil {
					log.Error(err)
				}
			},
			args: args{
				ctx: context.Background(),
				data: model.BankingCardData{
					ID:             1,
					Name:           "testCard",
					Number:         "123981723",
					ValidUntil:     "10/25",
					CardholderName: "Test Holder",
					CVV:            "213",
					Metadata:       nil,
				},
				beforeData: model.BankingCardData{
					Name:           "testCardAfter",
					Number:         "123981712312",
					ValidUntil:     "10/24",
					CardholderName: "Test User",
					CVV:            "213",
					Metadata: []model.Metadata{
						{
							Value: "test",
						},
					},
				},
				userID: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			postgres := tt.before(t, ctx)
			defer func() {
				if err := postgres.Terminate(ctx); err != nil {
					t.Fatalf("failed to terminate container: %s", err.Error())
				}
			}()
			endpoint, err := postgres.Endpoint(ctx, "")
			if err != nil {
				t.Error(err)
			}
			DBURI := fmt.Sprintf("postgresql://postgres:postgres@%s/goph_keeper?sslmode=disable", endpoint)
			r, _ := NewPostgresRepo(DBURI)
			tt.beforeCondition(r, tt.args.ctx, tt.args.beforeData, tt.args.userID)
			if err = r.UpdateBankingCardData(tt.args.ctx, tt.args.data, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("UpdateBankingCardData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostgresRepository_UpdateBinaryData(t *testing.T) {
	type beforeCondition func(repo *PostgresRepository, ctx context.Context, data model.BinaryData, userID int, link string)
	type args struct {
		ctx        context.Context
		data       model.BinaryData
		beforeData model.BinaryData
		userID     int
	}
	tests := []struct {
		name            string
		before          func(t *testing.T, ctx context.Context) testcontainers.Container
		beforeCondition beforeCondition
		args            args
		wantErr         bool
	}{
		{
			name:   "should update binary data",
			before: initContainers,
			beforeCondition: func(repo *PostgresRepository, ctx context.Context, data model.BinaryData, userID int, link string) {
				err := repo.MigrateDB("../migrations")
				if err != nil {
					log.Error(err)
				}
				err = repo.SaveBinaryData(ctx, &data, userID, link)
				if err != nil {
					log.Error(err)
				}
			},
			args: args{
				ctx: context.Background(),
				data: model.BinaryData{
					ID:   1,
					Name: "test_data",
					Link: "/tmp/test",
				},
				beforeData: model.BinaryData{
					ID:   1,
					Name: "test",
					Link: "/test/tmp",
				},
				userID: 0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			postgres := tt.before(t, ctx)
			defer func() {
				if err := postgres.Terminate(ctx); err != nil {
					t.Fatalf("failed to terminate container: %s", err.Error())
				}
			}()
			endpoint, err := postgres.Endpoint(ctx, "")
			if err != nil {
				t.Error(err)
			}
			DBURI := fmt.Sprintf("postgresql://postgres:postgres@%s/goph_keeper?sslmode=disable", endpoint)
			r, _ := NewPostgresRepo(DBURI)
			tt.beforeCondition(r, tt.args.ctx, tt.args.beforeData, tt.args.userID, tt.args.beforeData.Link)
			if err = r.UpdateBinaryData(tt.args.ctx, tt.args.data, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("UpdateBinaryData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostgresRepository_UpdateCredentialsData(t *testing.T) {
	type beforeCondition func(repo *PostgresRepository, ctx context.Context, data model.CredentialsData, userID int)
	type args struct {
		ctx        context.Context
		data       model.CredentialsData
		beforeData model.CredentialsData
		userID     int
	}
	tests := []struct {
		name            string
		args            args
		before          func(t *testing.T, ctx context.Context) testcontainers.Container
		beforeCondition beforeCondition
		wantErr         bool
	}{
		{
			name: "should update credentials data",
			args: args{
				ctx: context.Background(),
				data: model.CredentialsData{
					ID:       1,
					Name:     "test_cred_updated",
					Login:    "test_updated",
					Password: "test_pwd_updated",
				},
				beforeData: model.CredentialsData{
					ID:       1,
					Name:     "test_cred",
					Login:    "test",
					Password: "test_pwd",
				},
				userID: 1,
			},
			before: initContainers,
			beforeCondition: func(repo *PostgresRepository, ctx context.Context, data model.CredentialsData, userID int) {
				err := repo.MigrateDB("../migrations")
				if err != nil {
					log.Error(err)
				}
				err = repo.SaveCredentialsData(ctx, &data, userID)
				if err != nil {
					log.Error(err)
				}
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			postgres := tt.before(t, ctx)
			defer func() {
				if err := postgres.Terminate(ctx); err != nil {
					t.Fatalf("failed to terminate container: %s", err.Error())
				}
			}()
			endpoint, err := postgres.Endpoint(ctx, "")
			if err != nil {
				t.Error(err)
			}
			DBURI := fmt.Sprintf("postgresql://postgres:postgres@%s/goph_keeper?sslmode=disable", endpoint)
			r, _ := NewPostgresRepo(DBURI)
			tt.beforeCondition(r, tt.args.ctx, tt.args.beforeData, tt.args.userID)
			if err = r.UpdateCredentialsData(tt.args.ctx, tt.args.data, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("UpdateCredentialsData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostgresRepository_UpdateTextData(t *testing.T) {
	type beforeCondition func(repo *PostgresRepository, ctx context.Context, data model.TextData, userID int)
	type args struct {
		ctx        context.Context
		data       model.TextData
		beforeData model.TextData
		userID     int
	}
	tests := []struct {
		name            string
		before          func(t *testing.T, ctx context.Context) testcontainers.Container
		beforeCondition beforeCondition
		args            args
		wantErr         bool
	}{
		{
			name:   "should update text data",
			before: initContainers,
			beforeCondition: func(repo *PostgresRepository, ctx context.Context, data model.TextData, userID int) {
				err := repo.MigrateDB("../migrations")
				if err != nil {
					log.Error(err)
				}
				err = repo.SaveTextData(ctx, &data, userID)
				if err != nil {
					log.Error(err)
				}
			},
			args: args{
				ctx: context.Background(),
				data: model.TextData{
					ID:   1,
					Name: "updated text",
					Data: "updated texttexttext",
				},
				beforeData: model.TextData{
					ID:   1,
					Name: "text",
					Data: "text",
				},
				userID: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			postgres := tt.before(t, ctx)
			defer func() {
				if err := postgres.Terminate(ctx); err != nil {
					t.Fatalf("failed to terminate container: %s", err.Error())
				}
			}()
			endpoint, err := postgres.Endpoint(ctx, "")
			if err != nil {
				t.Error(err)
			}
			DBURI := fmt.Sprintf("postgresql://postgres:postgres@%s/goph_keeper?sslmode=disable", endpoint)
			r, _ := NewPostgresRepo(DBURI)
			tt.beforeCondition(r, tt.args.ctx, tt.args.beforeData, tt.args.userID)
			if err = r.UpdateTextData(tt.args.ctx, tt.args.data, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("UpdateTextData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostgresRepository_saveMetadata(t *testing.T) {
	type beforeCondition func(repo *PostgresRepository, data model.CredentialsData, userID int)
	type args struct {
		ctx        context.Context
		metadata   []model.Metadata
		beforeData model.CredentialsData
	}
	tests := []struct {
		name            string
		before          func(t *testing.T, ctx context.Context) testcontainers.Container
		beforeCondition beforeCondition
		args            args
		wantErr         bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			postgres := tt.before(t, ctx)
			defer func() {
				if err := postgres.Terminate(ctx); err != nil {
					t.Fatalf("failed to terminate container: %s", err.Error())
				}
			}()
			endpoint, err := postgres.Endpoint(ctx, "")
			if err != nil {
				t.Error(err)
			}
			DBURI := fmt.Sprintf("postgresql://postgres:postgres@%s/goph_keeper?sslmode=disable", endpoint)
			r, _ := NewPostgresRepo(DBURI)
			err = r.Transactional(tt.args.ctx, func() error {
				tt.beforeCondition(r, tt.args.beforeData, 1)
				return r.saveMetadata(tt.args.ctx, tt.args.metadata)
			})
			if err != nil != tt.wantErr {
				t.Errorf("saveMetadata() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostgresRepository_updateMetadata(t *testing.T) {
	type args struct {
		ctx        context.Context
		metadata   []model.Metadata
		beforeData model.CredentialsData
		dataID     int
	}
	tests := []struct {
		name            string
		before          func(t *testing.T, ctx context.Context) testcontainers.Container
		beforeCondition func(repo *PostgresRepository, data model.CredentialsData, userID int)
		args            args
		wantErr         bool
	}{
		{
			name:            "should update metadata",
			before:          initContainers,
			beforeCondition: initDatabase,
			args: args{
				ctx: context.Background(),
				metadata: []model.Metadata{
					{
						Value: "meta1",
					},
					{
						Value: "meta2",
					},
				},
				beforeData: model.CredentialsData{
					Name:     "test_cred",
					Login:    "test",
					Password: "test",
					Metadata: []model.Metadata{
						{
							Value: "test_value",
						},
					},
				},
				dataID: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			postgres := tt.before(t, ctx)
			defer func() {
				if err := postgres.Terminate(ctx); err != nil {
					t.Fatalf("failed to terminate container: %s", err.Error())
				}
			}()
			endpoint, err := postgres.Endpoint(ctx, "")
			if err != nil {
				t.Error(err)
			}
			DBURI := fmt.Sprintf("postgresql://postgres:postgres@%s/goph_keeper?sslmode=disable", endpoint)
			r, _ := NewPostgresRepo(DBURI)
			tt.beforeCondition(r, tt.args.beforeData, 1)
			if err = r.updateMetadata(tt.args.ctx, tt.args.metadata, tt.args.dataID); (err != nil) != tt.wantErr {
				t.Errorf("updateMetadata() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
