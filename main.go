package main

import (
	"fmt"
	"log"

	"github.com/xaliz06/dfs/p2p"
)

func OnPeer(peer p2p.Peer) error {
	fmt.Println("Some logc with the Peer outside of TCP Transport")

	return nil
}

func main() {
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.BurpHandshakeFunc,
		Decoder:       p2p.BurpDecoder{},
		OnPeer:        OnPeer,
	}
	tr := p2p.NewTCPTransport(tcpOpts)

	go func() {
		for {
			msg := <-tr.Consume()
			fmt.Printf("%s\n", msg)
		}
	}()

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
	// fmt.Println("Hello world!")
}
