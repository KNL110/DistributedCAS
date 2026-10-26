package storage

import (
	"bytes"
	"testing"
)


func TestStore(t *testing.T) {
	storeOpts := StoreOpts{
		PathTransformer: SHA256transform,
	}
	store := NewStore(storeOpts)

	key := "somekey"
	reader := bytes.NewReader([]byte("fuck"))
	err := store.writeStream(key,reader)
	if err != nil {
		t.Error(err)
	}
}