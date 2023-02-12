package repository

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type LocalBinaryRepository struct {
	saveDirectory string
}

func NewLocalBinaryRepository(directory string) *LocalBinaryRepository {
	return &LocalBinaryRepository{saveDirectory: directory}
}

func (r LocalBinaryRepository) Save(data io.Reader, filename string) (string, error) {
	fileLocation := filepath.FromSlash(fmt.Sprintf("%s/%s", r.saveDirectory, filename))
	f, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return "", err
	}
	_, err = io.Copy(f, data)
	if err != nil {
		return "", err
	}
	return fileLocation, nil
}

func (r LocalBinaryRepository) Load(link string) ([]byte, error) {
	fileLocation := filepath.FromSlash(fmt.Sprintf("%s/%s", r.saveDirectory, link))
	data, err := os.ReadFile(fileLocation)
	if err != nil {
		return nil, err
	}
	return data, nil
}
