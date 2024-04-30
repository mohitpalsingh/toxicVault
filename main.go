package main

import (
	"fmt"
	"log"

	"github.com/mohitpalsingh/toxicVault/p2p"
)

func main() {
	fmt.Println("Welcome to ToxicVault!")
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.NOPHandShakeFunc,
		Decoder:       p2p.DefaultDecoder{},
		// OnPeer:        func(p2p.Peer) error { return fmt.Errorf("error") },
	}

	tr := p2p.NewTCPTransport(tcpOpts)

	go func() {
		for {
			msg := <-tr.Consume()
			fmt.Printf("Message recieved -> %v\n", msg)
		}
	}()

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}

}
