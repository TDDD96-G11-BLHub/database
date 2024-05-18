package peerdb

import (
	"fmt"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-multiaddr"
)

type PeerNode struct {
	PeerInfo []multiaddr.Multiaddr
	Host     host.Host
	Protocol string
	Port     int
}

func InitNode() *PeerNode {

	node, err := libp2p.New()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Listen addresses:", node.Addrs())
	peerinfo := peer.AddrInfo{ID: node.ID(), Addrs: node.Addrs()}
	addrs, err := peer.AddrInfoToP2pAddrs(&peerinfo)
	if err != nil {
		panic(err)
	}

	res := PeerNode{PeerInfo: addrs, Host: node, Protocol: "tcp", Port: 61363}
	// shut the node down
	if err := node.Close(); err != nil {
		panic(err)
	}
	return &res
}
