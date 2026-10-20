package p2p

import (
	"fmt"
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

func (t* TCPtransport) ListenAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.listenAddr)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %v", t.listenAddr, err)
	}

	t.peers = make(map[net.Addr]peer)

	go t.acceptLoop()

	return nil
}

func (t *TCPtransport) acceptLoop()
