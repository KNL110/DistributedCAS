package p2p

type peer interface {
	Address() string  	// the network address of the peer.
	Disconnect() error	// terminate the connection to the peer.
}

//handles communication between the nodes, any transport protocol
type transport interface {
	SendMessage(peer peer, message []byte) error // SendMessage sends a message to the specified peer.
	ReceiveMessage() ([]byte, error)             // ReceiveMessage receives a message from any connected peer.
	ListenAndAccept() error
}