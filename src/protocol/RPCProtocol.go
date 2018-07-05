package protocol

import (
	"fmt"
	"bytes"
	"util"
	"net"
	"github.com/vmihailenco/msgpack"
	"encoding/hex"
)

type RPCProtocol struct {

}

var localNodeId string


func (rpc *RPCProtocol) Ping(st []string, args interface{}) {
	rpc.rpc("ping", st, args)
}

func (rpc *RPCProtocol) Stun(st []string) {
	rpc.rpc("stun", st, "")
}

// private
func (rpc *RPCProtocol) rpc(name string, address []string, args interface{}) {
	msgid := util.Digest(util.RandomBytes(20))

	var buffer bytes.Buffer
	method := []byte{00}
	argSlice := []string{name, args.(string)}
	encodeArg, err := msgpack.Marshal(argSlice)
	if err != nil {
		panic(err)
	}
	buffer.Write(method)
	buffer.Write(msgid)
	buffer.Write(encodeArg)
	fmt.Printf("send request %s with mesgId: %s for method: %s\n", address, hex.EncodeToString(msgid), name)
	// send a udp request to peer
	util.UdpRequest(address[0] + ":" + address[1], buffer)

	//fmt.Println("msg is", string(resData))
}

func (rpc *RPCProtocol) DatagramReceived (datagram []byte, raddr *net.UDPAddr, nodeId string) {
	fmt.Printf("received datagram from %s\n", raddr.String())
	localNodeId = nodeId

	if len(datagram) < 20 {
		fmt.Printf("received datagram too small from %s, ignoring\n", string(raddr.IP))
		return
	}

	method := hex.EncodeToString(datagram[0:1])
	msgId := datagram[1:21]
	data := datagram[21:]

	var unpackData []string = []string{}
	err := msgpack.Unmarshal(data, &unpackData)
	if err != nil {
		fmt.Println(err)
	}

	if method == "00" {
		AcceptRequest(msgId, unpackData, raddr.String())
	} else if method == "01" {
		AcceptRespond(msgId, unpackData, raddr.String())
	}
}

func AcceptRespond(msgId []byte, data []string, addr string) {
	fmt.Printf("received response %s for message id %s from %s\n", data[0], msgId, addr);
}

func AcceptRequest(msgId []byte, data []string, addr string) {
	var buffer bytes.Buffer
	method := []byte{01}
	//msgid := util.Digest(util.RandomBytes(20))

	buffer.Write(method)
	buffer.Write(msgId)

	// request name: ping stun ...
	fmt.Printf("accept response with act: %s \n", data[0])
	var protocol Protocol = Protocol{}
	switch data[0] {
	case "ping":
		buffer = protocol.PingRpc(buffer, localNodeId)
	case "stun":
		buffer = protocol.StunRpc(buffer, addr)
	default:
		panic("unknown name " + data[0])
	}

	fmt.Printf("sending response  for msgid %s to %s \n", hex.EncodeToString(msgId), addr)
	util.UdpRequest(addr, buffer)  // send udp data
	fmt.Println("after send request")
}
