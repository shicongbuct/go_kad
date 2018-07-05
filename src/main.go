package main

import (
	"os"
	"fmt"
	"protocol"
	"time"
	"encoding/hex"
	"util"
	"network"
	"node"
)

const MasterHost string = "39.106.106.129"
const MasterPort string  = "8000"
var pro protocol.RPCProtocol

func main() {
	if len(os.Args) < 2 {
		fmt.Println("we need a SN code")
		return
	}
	SN := os.Args[1]
	fmt.Println(SN)

	nodeId := hex.EncodeToString(util.Digest(util.RandomBytes(255)))
	fmt.Printf("nodeID: %s\n", nodeId)
	server := network.Server{&node.Node{nodeId,"127.0.0.1", "8000"}}
	go server.Listen()

	bootstrap(nodeId)
	keepalive()
}

func bootstrap(nodeId string) {
	pro.Ping([]string{MasterHost, MasterPort}, nodeId)
}

func keepalive() {
	for{
		fmt.Println("one turn")
		time.Sleep(30 * time.Second)
		pro.Stun([]string{MasterHost, MasterPort})
	}
}
