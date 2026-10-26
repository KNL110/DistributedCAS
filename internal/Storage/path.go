package storage

import "fmt"

type pathKey struct {
	pathName string
	ogName string
}

func (p *pathKey) fileName() string {
	return fmt.Sprintf("%s/%s",p.pathName,p.ogName)
}