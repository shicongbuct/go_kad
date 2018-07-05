package protocol

import (
	"github.com/vmihailenco/msgpack"
	"fmt"
	"bytes"
)

type Protocol struct{}


func (protocol *Protocol) PingRpc(buffer bytes.Buffer, localNodeId string) bytes.Buffer{
	resEncode, err := msgpack.Marshal(localNodeId)
	if err != nil {
		fmt.Println("acceptRequest ping error")
		fmt.Println(err)
	}
	buffer.Write(resEncode)
	return buffer
}

func (protocol *Protocol) StunRpc(buffer bytes.Buffer, addr string) bytes.Buffer{
	resEncode, err := msgpack.Marshal(addr)
	if err != nil {
		fmt.Println("acceptRequest stun error")
		fmt.Println(err)
	}
	buffer.Write(resEncode)
	return buffer
}