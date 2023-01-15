package storage

import (
	"os"
)

type LocalFS struct {
}

func (fs LocalFS) Write(filename string, content string) error {
	err := os.WriteFile(filename, []byte(content), 0600)
	if err != nil {
		return err
	}

	return nil
}

func (fs LocalFS) Read(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
