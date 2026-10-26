package storage

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)


func TestStore(t *testing.T) {
	storeOpts := StoreOpts{
		Root: t.TempDir(),
	}
	store := NewStore(storeOpts)
	data := []byte("some content")
	reader := bytes.NewReader(data)
	hash,err := store.Write(reader)
	if err != nil {
		t.Fatal(err)
	}
	sum := sha256.Sum256(data)
	if want := hex.EncodeToString(sum[:]); hash != want {
		t.Fatalf("hash = %s, want = %s", hash,want)
	}

	pk := NewpathKey(hash)
	if _, err := os.Stat(filepath.Join(store.Root, pk.fullPath())); err != nil {
		t.Fatalf("file not at content address: %v", err)
	}
}

func TestStoreRead(t *testing.T) {
	store := NewStore(StoreOpts{Root: t.TempDir()})
	data := []byte("some content")
	hash, err := store.Write(bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}

	rc, err := store.Read(hash)
	if err != nil {
		t.Fatal(err)
	}
	defer rc.Close()

	got, err := io.ReadAll(rc)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(got, data) {
		t.Fatalf("read = %q, want = %q", got, data)
	}
}

func TestStoreReadMissing(t *testing.T) {
	store := NewStore(StoreOpts{Root: t.TempDir()})
	sum := sha256.Sum256([]byte("never written"))

	_, err := store.Read(hex.EncodeToString(sum[:]))
	if !errors.Is(err, fs.ErrNotExist) {
		t.Fatalf("err = %v, want fs.ErrNotExist", err)
	}
}

func TestStoreReadInvalidHash(t *testing.T) {
	store := NewStore(StoreOpts{Root: t.TempDir()})

	for _, hash := range []string{"", "ab", "../../etc/passwd", strings.Repeat("z", 64)} {
		if _, err := store.Read(hash); err == nil {
			t.Errorf("Read(%q) succeeded, want error", hash)
		}
	}
}