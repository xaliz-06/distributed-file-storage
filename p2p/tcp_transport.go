package p2p

import (
	"fmt"
	"net"
	"sync"
)

// TCPPeer -> represents the remote node over a TCP established conn.
type TCPPeer struct {
	// underlying connection of the peer
	conn net.Conn

	// dial + retrive a connection => outbound == true
	// accept + retrive a connection => outbound == false
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

type TCPTransport struct {
	listenAddress string
	listener      net.Listener
	shakeHands    HandshakeFunc
	decoder       Decoder

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		shakeHands:    BurpHandshakeFunc,
		listenAddress: listenAddr,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error

	t.listener, err = net.Listen("tcp", t.listenAddress)

	if err != nil {
		return err
	}

	go t.startAcceptLoop()

	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP Accept error: %s\n", err)
		}

		go t.handleConn(conn)
	}
}

type Temp struct{}

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)

	if err := t.shakeHands(peer); err != nil {

	}

	// Read Loop
	msg := &Temp{}
	// lenDecodeErr := 0
	for {
		if err := t.decoder.Decode(conn, msg); err != nil {
			fmt.Println("TCP errorL %s\n", err)
			continue
		}
	}

	fmt.Printf("New incoming connection %+v\n", peer)
}
