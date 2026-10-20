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
		peers: make(map[net.Addr]peer),
	}
}

func (t* TCPtransport) ListenAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.listenAddr)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %v", t.listenAddr, err.Error())
	}
	go t.startAcceptLoop()

	return nil
}

func (t *TCPtransport) startAcceptLoop()  {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCPTransport error: %s\n", err.Error())
		}
		go t.handleConn(conn)
	}
}

func (t *TCPtransport) handleConn(conn net.Conn){
	peer := NewTCPPeer(conn, Inbound)
	t.peers[conn.RemoteAddr()] = peer
	fmt.Printf("New incoming connection {conn:%s, peerID: %s}\n", conn.RemoteAddr().String(), conn.RemoteAddr().String())
}
