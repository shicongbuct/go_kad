package protocol

import "testing"

// RPCProtocol.go
func TestRpcFunc(t *testing.T) {
	a := RPCProtocol{}
	var st []string = []string{"5abc", "9jjj"}
	a.Ping(st)

	//r := Add(1, 2)
	//if r != 3 {
	//	t.Errorf("Add(1, 2) failed. Got %d, expected 3.", r)
	//}
}