package main

import (
	"log"

	"github.com/KNL110/DistributedCAS/internal/config"
	"github.com/KNL110/DistributedCAS/internal/p2p"
)

func main() {

	config := config.MustLoad()

	tr := p2p.NewTCPtransport(
		config.Protocol.Port,
		p2p.GOBDecoder{},
	)
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatalf("Failed to start TCP transport: %v", err)
	}
	select {} // Block forever to keep the server running
}