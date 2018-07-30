// https://github.com/ethereumproject/go-ethereum/wiki/Peer-to-Peer
package main

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/discover"

)

const messageId = 0

type Message string

func MyProtocol() p2p.Protocol {
	return p2p.Protocol{
		Name:    "MyProtocol",
		Version: 1,
		Length:  1,
		Run:     msgHandler,
	}
}

func main() {
	boot,err := discover.ParseNode("enode://37c31d864b7205a0d2168737f67de1613cea752cb72c4045d84f8410efdc46cda06a71baebd8f36373d6fb53a7765b76e3097ea83a251ffc3909a6c903288b40@45.76.237.136:30301")

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(boot.String())

	nodekey, _ := crypto.GenerateKey()
	config := p2p.Config{
		MaxPeers:   10,
		PrivateKey: nodekey,
		Name:       "my node name",
		ListenAddr: ":30300",
		Protocols:  []p2p.Protocol{MyProtocol()},
		BootstrapNodes: []*discover.Node{boot},
	}

	srv := p2p.Server{
		Config:config,
	}

	if err := srv.Start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		srv.Start()
	}

	select {}
}

func msgHandler(peer *p2p.Peer, ws p2p.MsgReadWriter) error {
	for {
		msg, err := ws.ReadMsg()
		if err != nil {
			return err
		}

		var myMessage Message
		err = msg.Decode(&myMessage)
		if err != nil {
			// handle decode error
			continue
		}

		switch myMessage {
		case "foo":
			err := p2p.SendItems(ws, messageId, "bar")
			if err != nil {
				return err
			}
		default:
			fmt.Println("recv:", myMessage)
		}
	}

	return nil
}
