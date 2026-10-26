package storage

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

type StoreOpts struct {
	Root string
	PathMaker pathMaker
}

type Store struct {
	StoreOpts
}

func NewStore(opts StoreOpts) *Store {
	if opts.PathMaker == nil {
		opts.PathMaker = CASPathMake
	}
	return &Store{
		StoreOpts: opts,
	}
}

func (s *Store) writeStream(r io.Reader) (string, error) {

	if err := os.MkdirAll(s.Root, 0o755); err != nil {
		return "", err
	}

	tmp, err := os.CreateTemp(s.Root,"tmp-*")
	if err != nil {
		return "", fmt.Errorf("error creating temp file: %w",err)
	}
	defer os.Remove(tmp.Name()) //no-op once the rename happens successfully

	pathKey,err := s.PathMaker(tmp, r)
	if err != nil {
		return "",fmt.Errorf("path transform error: %w",err)
	}

	dir := filepath.Join(s.Root,pathKey.pathName)
	if err := os.MkdirAll(dir,0o755); err != nil {
		return "",err
	}

	finalPath := filepath.Join(s.Root,pathKey.fullPath())
	if err := os.Rename(tmp.Name(),finalPath); err != nil {
		return "",fmt.Errorf("error moving %s to %s: %w", tmp.Name(),finalPath,err)
	}

	log.Printf("written (%d) bytes to disk: %s", pathKey.written,finalPath)

	return pathKey.fileName,nil
}
