package storage

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"os"
	"path/filepath"
	"testing"
)


func TestStore(t *testing.T) {
	storeOpts := StoreOpts{
		root: t.TempDir(),
		PathMaker: CASPathMake,
	}
	store := NewStore(storeOpts)
	data := []byte("some content")
	reader := bytes.NewReader(data)
	hash,err := store.writeStream(reader)
	if err != nil {
		t.Fatal(err)
	}
	sum := sha256.Sum256(data)
	if want := hex.EncodeToString(sum[:]); hash != want {
		t.Fatalf("hash = %s, want = %s", hash,want)
	}

	pk := NewpathKey(hash)
	if _, err := os.Stat(filepath.Join(store.root, pk.fullPath())); err != nil {
		t.Fatalf("file not at content address: %v", err)
	}
}