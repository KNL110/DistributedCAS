package p2p

import (
	"fmt"
)

func Handshake(peer peer) error {
	fmt.Printf("Handshake with peer %s\n", peer.Address())
	return nil
}

func NopHandshake(peer) error {
	return nil
}