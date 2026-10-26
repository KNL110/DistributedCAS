package storage

import (
	"crypto/sha1"
	// "crypto/sha256"
	"encoding/hex"
	"strings"
)

type pathTransformer func(string) pathKey


func DefaultPathTransform(key string) pathKey {
	return pathKey{
		pathName: key,
		ogName: key,
	}
}

func SHA1transform(key string) pathKey {
	hash := sha1.Sum([]byte(key))
	hashStr := hex.EncodeToString(hash[:])
	println("the hash of file:", hashStr)
	sliceLen := 5
	blockSize := len(hashStr) / sliceLen
	paths := make([]string, sliceLen)

	for i := 0; i < sliceLen; i++ {
		from, to := i*blockSize,(i+1)*blockSize
		paths[i] = hashStr[from:to]
	}

	return pathKey{
		pathName: strings.Join(paths,"/"),
		ogName: hashStr,
	}
}