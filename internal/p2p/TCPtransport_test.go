package p2p

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTCPtransport(t *testing.T) {
	listenAddr := ":4000"
	tr := NewTCPtransport(
		listenAddr,
		GOBDecoder{},
	)

	assert.Equal(t,tr.listenAddr,listenAddr)
	assert.Nil(t,tr.ListenAndAccept())
}
