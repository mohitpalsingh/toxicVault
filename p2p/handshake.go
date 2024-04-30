package p2p

type HandshakeFunc func(Peer) error

func NOPHandShakeFunc(Peer) error {
	return nil
}
