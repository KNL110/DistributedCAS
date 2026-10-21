package p2p

import (
	"fmt"
	"net"
)



type RPC struct {
	from net.Addr
	payload []byte
}

func (m *RPC) String() string {
	return fmt.Sprintf("{from:%s payload:%q}", m.from, m.payload)
}