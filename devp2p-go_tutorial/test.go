
package main

import "github.com/ethereum/go-ethereum/crypto"
import "github.com/ethereum/go-ethereum/p2p"


func main() {
	nodekey, _ := crypto.GenerateKey()
	config := p2p.Config{
		MaxPeers:10,
		PrivateKey: nodekey,
		Name:       "my node name",
		ListenAddr: ":30300",
		Protocols:  []p2p.Protocol{},
	}
	srv := p2p.Server {
		Config:config,
	}
	srv.Start()
}
