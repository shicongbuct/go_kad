package main

import (
	"os"
	"fmt"
	"net"
)

func main() {
	//conn, err := net.Dial("udp", "192.168.1.56:8000")
	//conn, err := net.Dial("udp", "39.106.106.129:8000")
	conn, err := net.Dial("udp", "192.168.1.56:8000")
	defer conn.Close()
	if err != nil {
		os.Exit(1)
	}
	testByte := []byte {0,63,112,51,36,79,49,226,5,142,125,231,172,150,139,50,207,171,38,133,153,146,164,112,105,110,103,217,40,98,49,97,55,53,49,52,56,98,48,48,49,102,99,53,98,99,101,51,50,49,52,54,53,51,97,101,52,56,55,97,100,57,53,56,54,49,101,50,101}
	// nodeid b1a75148b001fc5bce3214653ae487ad95861e2e
	// msgid  3f7033244f31e2058e7de7ac968b32cfab268599
	fmt.Println("XXXX", string(testByte[:]))
	//conn.Write([]byte("Hello shicong111!as"))
	conn.Write(testByte)
	fmt.Println("send msg")

	var msg [20]byte
	conn.Read(msg[0:])

	fmt.Println("msg is", string(msg[0:]))
}
