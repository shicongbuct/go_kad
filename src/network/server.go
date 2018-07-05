package network

import (
	"node"
	"fmt"
	"os"
	"net"
	"protocol"
)

type Server struct {
	Node *node.Node
}

func (server *Server) Listen () {
	udp_addr, err := net.ResolveUDPAddr("udp", ":" + server.Node.Port)
	checkError(err)
	for {
		conn, err := net.ListenUDP("udp", udp_addr)
		checkError(err)
		server.recvUDPMsg(conn)
		conn.Close()
	}
}

func checkError(err error){
	if  err != nil {
		fmt.Println("Error: %s", err.Error())
		os.Exit(1)
	}
}

func (server *Server) recvUDPMsg(conn *net.UDPConn){
	var buf [80]byte
	n, raddr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		return
	}

	_, err = conn.WriteToUDP([]byte("take it"), raddr)
	checkError(err)

	protocol := protocol.RPCProtocol{}
	protocol.DatagramReceived(buf[0:n], raddr, server.Node.Id)
}