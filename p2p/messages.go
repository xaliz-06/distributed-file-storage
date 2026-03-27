package p2p

import "net"

// RPC hold any data tranported between two nodes
// in the network by the transporter
type RPC struct {
	Payload []byte
	From    net.Addr
}
