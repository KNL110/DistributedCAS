package p2p

import (
	"net"
	"sync"
)

type TCPtransport struct {
	listenAddr string
	listener net.Listener

	mu sync.RWMutex
	peers map[net.Addr]peer
}

func NewTCPtransport(listenAddr string) *TCPtransport {
	return &TCPtransport{
		listenAddr: listenAddr,
	}
}

