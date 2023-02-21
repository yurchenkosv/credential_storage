package repository

import "io"

type BinaryRepository interface {
	Save(data io.Reader, filename string) (string, error)
	Load(link string) ([]byte, error)
}
