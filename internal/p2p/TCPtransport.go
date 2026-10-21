package p2p

import (
	"fmt"
	"net"
	"sync"
)

type TCPtransport struct {
	listenAddr string
	listener net.Listener
	shakeHands func(peer peer) error
	decoder Decoder

	mu sync.RWMutex
	peers map[net.Addr]peer
}

func NewTCPtransport(listenAddr string,decoder Decoder) *TCPtransport {
	return &TCPtransport{
		listenAddr: listenAddr,
		peers: make(map[net.Addr]peer),
		shakeHands: Handshake,
		decoder: decoder,
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


type temp struct {}
func (t *TCPtransport) handleConn(conn net.Conn){
	peer := NewTCPPeer(conn, Inbound)

	if err := t.shakeHands(peer); err != nil {
		conn.Close()
		fmt.Printf("TCP handshake error: %s\n", err)
		return
	}
	fmt.Printf("New incoming connection {conn:%s}\n", conn.RemoteAddr().String())

	msg := &message{}
	for {
		if err := t.decoder.Decode(conn, msg); err != nil {
			conn.Close()
			fmt.Printf("TCP Decode Error: %s\n", err)
			return
		}
		// Process the decoded message (msg) here
		fmt.Printf("message: %+v\n",msg)
	}
}
