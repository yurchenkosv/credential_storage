package repository

import (
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"path/filepath"
	"testing"
)

func TestLocalBinaryRepository_Load(t *testing.T) {
	type before func(r BinaryRepository, filename string)
	type after func(filename string)
	type fields struct {
		saveDirectory string
	}
	type args struct {
		link string
	}
	tests := []struct {
		name    string
		fields  fields
		before  before
		after   after
		args    args
		want    []byte
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "should load data from file",
			fields: fields{
				saveDirectory: ".",
			},
			before: func(r BinaryRepository, filename string) {
				reader := bytes.NewReader([]byte("test"))
				_, err := r.Save(reader, filename)
				if err != nil {
					log.Error(err)
				}
			},
			after: func(filename string) {
				err := os.Remove(filename)
				if err != nil {
					log.Error(err)
				}
			},
			args: args{
				link: filepath.FromSlash("./tmp_file"),
			},
			want:    []byte("test"),
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := LocalBinaryRepository{
				saveDirectory: tt.fields.saveDirectory,
			}
			tt.before(r, "tmp_file")
			got, err := r.Load(tt.args.link)
			if !tt.wantErr(t, err, fmt.Sprintf("Load(%v)", tt.args.link)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Load(%v)", tt.args.link)
			tt.after(tt.args.link)
		})
	}
}

func TestLocalBinaryRepository_Save(t *testing.T) {
	type after func(filepath string)
	type fields struct {
		saveDirectory string
	}
	type args struct {
		data     io.Reader
		filename string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		after   after
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "should save bytes in file",
			fields: fields{
				saveDirectory: ".",
			},
			args: args{
				data:     bytes.NewReader([]byte("test")),
				filename: "test_file",
			},
			want: filepath.FromSlash("test_file"),
			after: func(filepath string) {
				err := os.Remove(filepath)
				if err != nil {
					log.Error(err)
				}
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := LocalBinaryRepository{
				saveDirectory: tt.fields.saveDirectory,
			}
			got, err := r.Save(tt.args.data, tt.args.filename)
			if !tt.wantErr(t, err, fmt.Sprintf("Save(%v, %v)", tt.args.data, tt.args.filename)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Save(%v, %v)", tt.args.data, tt.args.filename)
			tt.after(tt.want)
		})
	}
}

func TestNewLocalBinaryRepository(t *testing.T) {
	type args struct {
		directory string
	}
	tests := []struct {
		name string
		args args
		want *LocalBinaryRepository
	}{
		{
			name: "should create repo",
			args: args{
				directory: "/tmp/test",
			},
			want: &LocalBinaryRepository{
				saveDirectory: "/tmp/test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewLocalBinaryRepository(tt.args.directory), "NewLocalBinaryRepository(%v)", tt.args.directory)
		})
	}
}
