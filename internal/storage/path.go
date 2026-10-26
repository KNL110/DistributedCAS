package storage

import (
	"path/filepath"
)

type pathKey struct {
	pathName string
	fileName string
	written int64
}

func NewpathKey(hash string) *pathKey {
	return &pathKey{
		pathName: filepath.Join(hash[:2],hash[2:4]),
		fileName: hash,
		written: -1,
	}
}
func NewpathKey1(hash string) *pathKey {
	return &pathKey{
		pathName: hash[:2],
		fileName: hash,
		written: -1,
	}
}

func (p *pathKey) fullPath() string {
	return filepath.Join(p.pathName,p.fileName)
}

