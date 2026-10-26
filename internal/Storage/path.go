package storage

import "fmt"

type pathKey struct {
	pathName string
	fileName string
}

func (p *pathKey) fullPath() string {
	return fmt.Sprintf("%s/%s",p.pathName,p.fileName)
}