package p2p

import (
	"encoding/gob"
	"io"
)

type Decoder interface {
	Decode(io.Reader,*RPC) error
}


type GOBDecoder struct {}

func (d GOBDecoder) Decode(r io.Reader,msg *RPC) error {
	return gob.NewDecoder(r).Decode(msg)
}

type DefaultDecoder struct {}

func (d DefaultDecoder) Decode(r io.Reader, msg *RPC) error {
	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	if err != nil {
		return err
	}
	if n == 0 {
		return io.EOF
	}

	msg.payload = buf[:n]
	return nil
}