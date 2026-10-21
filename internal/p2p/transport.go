package p2p

type Peer interface {
	Address() string  	// the network address of the peer.
	Close() error		// close the connection to the peer.
}

//handles communication between the nodes, any transport protocol
type Transport interface {
	SendMessage(peer Peer, message []byte) error // SendMessage sends a message to the specified peer.
	ReceiveMessage() ([]byte, error)             // ReceiveMessage receives a message from any connected peer.
	ListenAndAccept() error
	Consume() <-chan RPC
}