package p2p

import (
	"testing"

	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestTCPtransport(t *testing.T) {
	listenAddr := ":4000"
	tr := NewTCPtransport(listenAddr)

	assert.Equal(t,tr.listenAddr,listenAddr)
	fmt.Println("ok")
}
