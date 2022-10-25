package image

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

type Storage struct {
	basePath string
}

func NewStorage(basePath string) (*Storage, error) {
	p, err := filepath.Abs(basePath)
	if err != nil {
		return nil, err
	}

	return &Storage{basePath: p}, nil
}

func (l *Storage) Save(path string, contents io.Reader) error {
	fullPath := l.fullPath(path)

	dir := filepath.Dir(fullPath)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}

	_, err = os.Stat(fullPath)
	if err == nil {
		err = os.Remove(fullPath)
		if err != nil {
			return err
		}
	} else if !os.IsNotExist(err) {
		return err
	}
	f, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, contents)
	if err != nil {

		return err
	}
	log.Println("----------------------------file saved")
	log.Println(l.basePath)
	log.Println(fullPath)
	log.Println("----------------------------file saved")
	return nil
}

func (l *Storage) Get(path string) (*os.File, error) {
	fp := l.fullPath(path)

	f, err := os.Open(fp)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func (l *Storage) fullPath(path1 string) string {
	return filepath.Join(l.basePath, path1)
}
