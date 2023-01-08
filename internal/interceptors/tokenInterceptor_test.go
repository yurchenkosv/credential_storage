package interceptors

import (
	"github.com/golang-jwt/jwt/v4"
	"testing"
)

func Test_extractUserFromClaims(t *testing.T) {
	type args struct {
		claims jwt.Claims
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "valid claim",
			args: args{
				claims: jwt.MapClaims{
					"user": "user1",
				},
			},
			want:    "user1",
			wantErr: false,
		},
		{
			name: "invalid claim",
			args: args{
				claims: nil,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "missing user claim",
			args: args{
				claims: jwt.MapClaims{
					"foo": "bar",
				},
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := extractUserFromClaims(tt.args.claims)
			if (err != nil) != tt.wantErr {
				t.Errorf("extractUserFromClaims() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("extractUserFromClaims() got = %v, want %v", got, tt.want)
			}
		})
	}
}
