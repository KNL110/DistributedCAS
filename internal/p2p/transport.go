package p2p

type peer interface {
	ID() string       	// unique identifier of the peer.
	Address() string  	// the network address of the peer.
	Connect() error   	// establishe a connection to the peer.
	Disconnect() error	// terminate the connection to the peer.
}

//handles communication between the nodes, any transport protocol
type transport interface {
	SendMessage(peer peer, message []byte) error // SendMessage sends a message to the specified peer.
	ReceiveMessage() ([]byte, error)             // ReceiveMessage receives a message from any connected peer.
}