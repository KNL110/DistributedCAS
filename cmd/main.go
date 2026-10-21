package main

import (
	"fmt"
	"log"

	"github.com/KNL110/DistributedCAS/internal/config"
	"github.com/KNL110/DistributedCAS/internal/p2p"
)

func main() {

	config := config.MustLoad()

	tr := p2p.NewTCPtransport(
		config.Protocol.Port,
		p2p.DefaultDecoder{},
	)
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatalf("Failed to start TCP transport: %v", err)
	}

	go func() {
		msg := <-tr.Consume()
		fmt.Printf("Received message: %+v\n",msg)
	}()
	select {} // Block forever to keep the server running
}