package storage

import (
	"path/filepath"
)

type PathKey struct {
	PathName string
	FileName string
}

func NewpathKey(hash string) *PathKey {
	return &PathKey{
		PathName: filepath.Join(hash[:2],hash[2:4]),
		FileName: hash,
	}
}
func NewpathKey1(hash string) *PathKey {
	return &PathKey{
		PathName: hash[:2],
		FileName: hash,
	}
}

func (p *PathKey) fullPath() string {
	return filepath.Join(p.PathName,p.FileName)
}

