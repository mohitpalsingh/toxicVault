package p2p

// Peer represents the external node.
type Peer interface {
	Close() error
}

// Transport represents anything that handles the communication between the nodes in the network.
type Transport interface {
	ListenAndAccept() error
	Consume() <-chan RPC
}
