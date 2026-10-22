package storage

import (
	"fmt"
	"io"
	"log"
	"os"
)

type StoreOpts struct {
	PathTransformer pathTransformer
}

type Store struct {
	StoreOpts
}

func NewStore(opts StoreOpts) *Store {
	return &Store{
		StoreOpts: opts,
	}
}

func (s *Store) writeStream(key string, r io.Reader) error {
	pathName := s.PathTransformer(key)

	if err := os.MkdirAll(pathName,os.ModePerm); err != nil {
		return err
	}

	filename := "test.txt"
	filePath := pathName + "/" + filename

	f,err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error opening file %s: %w", filePath, err)
	}
	defer f.Close()

	n, err := io.Copy(f,r)
	if err != nil {
		return fmt.Errorf("error writing to file %s: %w", f.Name(), err)
	}

	log.Printf("written (%d) bytes to disk: %s", n,filePath)

	return nil
}
