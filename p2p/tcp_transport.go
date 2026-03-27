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

type TCPTransportOpts struct {
	ListenAddr    string
	HandshakeFunc HandshakeFunc
	Decoder       Decoder
}

type TCPTransport struct {
	TCPTransportOpts
	listener net.Listener

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error

	t.listener, err = net.Listen("tcp", t.ListenAddr)

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

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)

	if err := t.HandshakeFunc(peer); err != nil {
		conn.Close()
		fmt.Printf("TCP Handshake Error: %s\n", err)

		return
	}

	// Read Loop
	msg := &RPC{}
	// buf := make([]byte, 2000)
	for {
		// n, err := conn.Read(buf)
		// if err != nil {
		// 	fmt.Printf("TCP Error: %s\n", err)
		// }

		if err := t.Decoder.Decode(conn, msg); err != nil {
			fmt.Printf("TCP error %s\n", err)
			continue
		}

		msg.From = conn.RemoteAddr()
		fmt.Printf("RPC: %s\n", msg)

	}
}
