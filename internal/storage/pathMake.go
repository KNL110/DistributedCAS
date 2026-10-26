package storage

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"encoding/hex"
)

type pathMaker func(tmp *os.File,r io.Reader) (*pathKey,error)


func CASPathMake(tmp *os.File,r io.Reader) (*pathKey,error) {
	h := sha256.New()
	written, err := io.Copy(io.MultiWriter(tmp,h),r)
	if err != nil {
		tmp.Close()
		return nil,fmt.Errorf("error writing to file %s: %w", tmp.Name(), err)
	}
	
	if err := tmp.Sync();err != nil {
		tmp.Close()
		return nil,err
	}

	if err := tmp.Close(); err != nil {
		return nil,err
	}

	hash := hex.EncodeToString(h.Sum(nil))
	pk := NewpathKey(hash)
	pk.written = written
	return pk,nil
}
