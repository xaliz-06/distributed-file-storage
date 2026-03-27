package p2p

type HandshakeFunc func(Peer) error

func BurpHandshakeFunc(Peer) error { return nil }
