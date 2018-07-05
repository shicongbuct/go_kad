package main

import (
	"net"
	"os"
	"fmt"
)

func main() {
	udp_addr, err := net.ResolveUDPAddr("udp", "192.168.1.56:8000")
	checkError(err)
	for {
		conn, err := net.ListenUDP("udp", udp_addr)
		checkError(err)
		recvUDPMsg(conn)
		conn.Close()
	}
}

func checkError(err error){
	if  err != nil {
		fmt.Println("Error: %s", err.Error())
		os.Exit(1)
	}
}

func recvUDPMsg(conn *net.UDPConn){
	var buf [80]byte
	n, raddr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		return
	}
	fmt.Println("msg is", buf[0:n], n)

	_, err = conn.WriteToUDP([]byte("nice to see u"), raddr)
	checkError(err)
}

