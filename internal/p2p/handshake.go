package p2p

import (
	"fmt"
)

func Handshake(peer Peer) error {
	fmt.Printf("Handshake with peer %s\n", peer.Address())
	return nil
}

func NopHandshake(Peer) error {
	return nil
}