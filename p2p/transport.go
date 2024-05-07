package p2p

import "net"

// Peer represents the external node.
type Peer interface {
	net.Conn
	Send([]byte) error
	CloseStream()
}

// Transport represents anything that handles the communication between the nodes in the network.
type Transport interface {
	Addr() string
	Dial(string) error
	ListenAndAccept() error
	Consume() <-chan RPC
	Close() error
}
