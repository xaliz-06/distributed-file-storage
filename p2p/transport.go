package p2p

// PEER -> interface representing a remote node
type Peer interface {
}

// TRANSPORT -> interface representing the communication layer
// between the nodes in the network
type Transport interface {
	ListenAndAccept() error
}
