package main

import (
	"log"

	"github.com/xaliz06/dfs/p2p"
)

func main() {
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.BurpHandshakeFunc,
		Decoder:       p2p.BurpDecoder{},
	}
	tr := p2p.NewTCPTransport(tcpOpts)

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
	// fmt.Println("Hello world!")
}
