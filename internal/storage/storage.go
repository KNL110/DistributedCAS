package storage

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

type StoreOpts struct {
	Root string
	PathKey func (string) *PathKey
}

type Store struct {
	StoreOpts
}

func NewStore(opts StoreOpts) *Store {
	if opts.PathKey == nil {
		opts.PathKey = NewpathKey
	}
	return &Store{
		StoreOpts: opts,
	}
}

func (s *Store) Read(hash string) (io.ReadCloser, error) {
	if !validHash(hash) {
		return nil, fmt.Errorf("invalid hash %q", hash)
	}
	pathKey := s.PathKey(hash)
	f, err := os.Open(filepath.Join(s.Root, pathKey.fullPath()))
	if err != nil {
		return nil, fmt.Errorf("error opening %s: %w", hash, err)
	}
	
	return f,err
}

func (s *Store) Write(r io.Reader) (string, error) {

	if err := os.MkdirAll(s.Root, 0o755); err != nil {
		return "", err
	}

	tmp, err := os.CreateTemp(s.Root,"tmp-*")
	if err != nil {
		return "", fmt.Errorf("error creating temp file: %w",err)
	}
	defer os.Remove(tmp.Name()) //no-op once the rename happens successfully

	h := sha256.New()
	written, err := io.Copy(io.MultiWriter(tmp,h),r)
	if err != nil {
		tmp.Close()
		return "",fmt.Errorf("error writing to file %s: %w", tmp.Name(), err)
	}
	
	if err := tmp.Sync();err != nil {
		tmp.Close()
		return "",err
	}

	if err := tmp.Close(); err != nil {
		return "",err
	}

	hash := hex.EncodeToString(h.Sum(nil))
	if s.Has(hash) {
		log.Printf("content already stored, skipping write: %s", hash)
		return hash, nil
	}
	pathKey := s.PathKey(hash)

	dir := filepath.Join(s.Root,pathKey.PathName)
	if err := os.MkdirAll(dir,0o755); err != nil {
		return "",err
	}

	finalPath := filepath.Join(s.Root,pathKey.fullPath())
	if err := os.Rename(tmp.Name(),finalPath); err != nil {
		return "",fmt.Errorf("error moving %s to %s: %w", tmp.Name(),finalPath,err)
	}

	log.Printf("written (%d) bytes to disk: %s", written,finalPath)

	return pathKey.FileName,nil
}

func (s *Store) Has(hash string) bool {
	if !validHash(hash) {
		return false
	}
	pathKey := s.PathKey(hash)
	_, err := os.Stat(filepath.Join(s.Root, pathKey.fullPath()))
	return err == nil
}

func (s *Store) Delete(hash string) error {
	if !validHash(hash) {
		return fmt.Errorf("invalid hash %q", hash)
	}
	pathKey := s.PathKey(hash)
	filePath := filepath.Join(s.Root, pathKey.fullPath())
	if err := os.Remove(filePath); err != nil {
		return fmt.Errorf("error deleting %s: %w", hash, err)
	}

	log.Printf("deleted from disk: %s", filePath)
	return nil
}


func validHash(hash string) bool {
	if len(hash) != sha256.Size*2 {
		return false
	}
	_, err := hex.DecodeString(hash)
	return err == nil
}
