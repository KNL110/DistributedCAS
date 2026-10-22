package storage

import (
	"bytes"
	"testing"
	"os"
)


func TestStore(t *testing.T) {
	storeOpts := StoreOpts{
		PathTransformer: DefaultPathTransform,
	}
	store := NewStore(storeOpts)

	key := "somekey"
	reader := bytes.NewReader([]byte("fuck"))
	if err := store.writeStream(key,reader); err != nil {
		t.Error(err)
	}
	t.Cleanup(func() {os.RemoveAll(key)})
}