package p2p

import (
	"fmt"
	"net"
)



type message struct {
	from net.Addr
	payload []byte
}

func (m *message) String() string {
	return fmt.Sprintf("{from:%s payload:%q}", m.from, m.payload)
}