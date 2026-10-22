package storage

import (
	// "io"
	// "os"
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

// func (s *Store) writeStream(key string, r io.Reader) error {
// 	pathName := s.PathTransformer(key)

// 	if err := os.MkdirAll(pathName,os.ModePerm); err != nil {
// 		return err
// 	}


// 	return nil
// }
