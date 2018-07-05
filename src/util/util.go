package util

import (
	"math/rand"
	"time"
	"crypto/sha1"
	"net"
	"fmt"
	"bytes"
)

func RandomBytes(byteNum int) []byte{
	rNum := rand.New(rand.NewSource(time.Now().UnixNano()))
	var result []byte
	for i := 0; i < byteNum; i++ {
		r := int (rNum.Float32() * 256)
		result = append(result, byte(r))
	}
	return result
}

func Digest(data []byte) []byte{
	sha := sha1.New()
	sha.Write(data)
	result := sha.Sum(nil)
	return result[0:20]
}

func UdpRequest(address string, buffer bytes.Buffer) []byte{
	conn, err := net.Dial("udp", address)
	if err != nil {
		fmt.Printf("request address:%s with error:\n", address)
		fmt.Println(err)
	}
	conn.Write(buffer.Bytes())

	// create a []byte to hold the response data
	var resMsg [100]byte
	conn.SetReadDeadline(time.Now().Add(10 * time.Second)) // timeout
	n, err := conn.Read(resMsg[0:])

	if err != nil {
		fmt.Printf("get confirm reponse from address %s failed with error:\n", address)
		fmt.Println(err)
	}
	return resMsg[0:n]
}