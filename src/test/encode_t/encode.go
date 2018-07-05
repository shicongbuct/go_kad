package main
//go:generate msgp
import (
	"reflect"
	"io"
	"net/rpc"
	"net"
	"github.com/ugorji/go/codec"
	"fmt"
)

func main() {
	// create and configure Handle
	var (
		bh codec.BincHandle
		mh codec.MsgpackHandle
		//ch codec.CborHandle
	)

	mh.MapType = reflect.TypeOf(map[string]interface{}(nil))

	fmt.Println(mh.MapType.String())

	// configure extensions
	// e.g. for msgpack, define functions and enable Time support for tag 1
	// mh.SetExt(reflect.TypeOf(time.Time{}), 1, myExt)

	// create and use decoder/encoder
	var (
		r io.Reader
		w io.Writer
		b []byte
		h = &bh // or mh to use msgpack
	)

	dec := codec.NewDecoder(r, h)
	dec = codec.NewDecoderBytes(b, h)
	err := dec.Decode(&v)

	if err != nil {
		fmt.Println(err)
	}

	enc := codec.NewEncoder(w, h)
	enc = codec.NewEncoderBytes(&b, h)
	err = enc.Encode(v)

	//RPC Server
	go func() {
		for {
			conn, _ := listener.Accept()
			rpcCodec := codec.GoRpc.ServerCodec(conn, h)
			//OR rpcCodec := codec.MsgpackSpecRpc.ServerCodec(conn, h)
			rpc.ServeCodec(rpcCodec)
		}
	}()

	//RPC Communication (client side)
	conn, _ := net.Dial("tcp", "localhost:5555")
	rpcCodec := codec.GoRpc.ClientCodec(conn, h)
	//OR rpcCodec := codec.MsgpackSpecRpc.ClientCodec(conn, h)
	client := rpc.NewClientWithCodec(rpcCodec)
}
