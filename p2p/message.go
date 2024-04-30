package p2p

import "net"

type RPC struct {
	From    net.Addr
	Payload []byte
}
