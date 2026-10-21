package p2p

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func onPeer(peer Peer) error {
	fmt.Printf("New incoming connection {conn:%s}\n", peer.Address())
	return nil
}
func TestTCPtransport(t *testing.T) {
	listenAddr := ":0" // let the OS pick a free port so the test doesn't clash with a running node
	tr := NewTCPtransport(
		listenAddr,
		GOBDecoder{},
		onPeer,
	)

	assert.Equal(t,tr.listenAddr,listenAddr)
	assert.Nil(t,tr.ListenAndAccept())
	t.Cleanup(func() { tr.Close() })
}
