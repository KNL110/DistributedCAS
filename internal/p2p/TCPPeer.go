package p2p

import "net"


type bound string //weather its an inbound or outbound connection

const (
	Inbound  bound = "inbound"
	Outbound bound = "outbound"
)

type TCPPeer struct {
	conn net.Conn
	bound bound
}

func NewTCPPeer(conn net.Conn, bound bound) *TCPPeer {
	return &TCPPeer{
		conn: conn,
		bound: bound,
	}
}

func (p *TCPPeer) Address() string {
	return p.conn.RemoteAddr().String()
}

func (p *TCPPeer) Close() error {
	return p.conn.Close()
}