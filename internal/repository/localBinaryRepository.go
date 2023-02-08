package repository

import (
	"fmt"
	"os"
	"path/filepath"
)

type LocalBinaryRepository struct {
	saveDirectory string
}

func NewLocalBinaryRepository(directory string) *LocalBinaryRepository {
	return &LocalBinaryRepository{saveDirectory: directory}
}

func (r LocalBinaryRepository) Save(data []byte, filename string) (string, error) {
	fileLocation := filepath.FromSlash(fmt.Sprintf("%s/%s", r.saveDirectory, filename))
	err := os.WriteFile(fileLocation, data, 0600)
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
