package store

import (
	"io"
	"os"
)

type File struct {
	path string
}

func NewFile(path string) *File {
	return &File{
		path: path,
	}
}

func (s *File) Get() ([]byte, error) {
	f, err := os.OpenFile(s.path, os.O_RDONLY, 0777)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (s *File) Set(data []byte) error {
	f, err := os.OpenFile(s.path, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		return err
	}

	defer f.Close()

	if _, err := f.Write(data); err != nil {
		return err
	}

	return nil
}
