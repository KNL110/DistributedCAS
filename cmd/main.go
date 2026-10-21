package main

import (
	"fmt"
	"log"

	"github.com/KNL110/DistributedCAS/internal/config"
	"github.com/KNL110/DistributedCAS/internal/p2p"
)

func onPeer(peer p2p.Peer) error {
	fmt.Printf("New incoming connection {conn:%s}\n", peer.Address())
	return nil
}

func main() {

	config := config.MustLoad()

	tr := p2p.NewTCPtransport(
		config.Protocol.Port,
		p2p.DefaultDecoder{},
		onPeer,
	)
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatalf("Failed to start TCP transport: %v", err)
	}

	go func() {
		for {
			msg := <-tr.Consume()
			fmt.Printf("Received message: %+v\n",msg)
		}
	}()
	select {} // Block forever to keep the server running
}