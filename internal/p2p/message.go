package p2p

import (
	"fmt"
	"net"
)



type RPC struct {
	from net.Addr
	payload []byte
}
//TODO: see if fxn by value is better than by reference
func (m RPC) String() string {
	return fmt.Sprintf("{from:%s payload:%q}", m.from, m.payload)
}