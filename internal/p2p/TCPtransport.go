package p2p

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"time"
)

type TCPtransport struct {
	listenAddr string
	listener net.Listener

	decoder Decoder
	rpcChan chan RPC

	shakeHands func(peer Peer) error
	onPeer func(peer Peer) error
}

func NewTCPtransport(listenAddr string,decoder Decoder,onPeer func(Peer) error) *TCPtransport {
	return &TCPtransport{
		listenAddr: listenAddr,
		shakeHands: Handshake,
		decoder: decoder,
		rpcChan: make(chan RPC),
		onPeer: onPeer,
	}
}

func (t* TCPtransport) ListenAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.listenAddr)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", t.listenAddr, err)
	}
	go t.startAcceptLoop()
	
	return nil
}

// a read-only channel that returns RPC messages received from peers
func(t* TCPtransport) Consume() <-chan RPC {
	return t.rpcChan
}

// stops accepting new connections
func (t *TCPtransport) Close() error {
	return t.listener.Close()
}

func (t *TCPtransport) startAcceptLoop()  {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				fmt.Printf("TCPTransport listener closed, stopping accept loop\n")
				return
			}
			fmt.Printf("TCPTransport error: %s\n", err.Error())
			time.Sleep(100 * time.Millisecond) // avoid a hot loop on persistent errors like EMFILE
			continue
		}
		go t.handleConn(conn)
	}
}

func (t *TCPtransport) handleConn(conn net.Conn){
	defer conn.Close()
	peer := NewTCPPeer(conn, Inbound)

	if err := t.shakeHands(peer); err != nil {
		fmt.Printf("TCP handshake error: %s\n", err)
		return
	}

	if t.onPeer != nil {
		if err := t.onPeer(peer); err != nil {
			fmt.Printf("TCP onPeer error: %s\n", err)
			return
		}
	}

	r := bufio.NewReader(conn)
	msg := RPC{}
	msg.from = conn.RemoteAddr()
	for {
		if err := t.decoder.Decode(r, &msg); err != nil {
			if !errors.Is(err, io.EOF) {
				fmt.Printf("TCP Decode Error: %s\n", err)
			}
			return
		}
		t.rpcChan <- msg
	}
}
