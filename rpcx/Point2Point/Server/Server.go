package main

import (
	"github.com/smallnest/rpcx"
)

// Args in rpcx
type Args struct {
	A int `msg:"a"`
	B int `msg:"b"`
}

// Reply in rpcx
type Reply struct {
	C int `msg:"c"`
}

// Arith ...
type Arith int

// Mul rpcx args.A * args.B
func (t *Arith) Mul(args *Args, reply *Reply) error {
	reply.C = args.A * args.B
	return nil
}

// Error panic "ERROR"
func (t *Arith) Error(args *Args, reply *Reply) error {
	panic("ERROR")
}

func main() {
	server := rpcx.NewServer()
	server.RegisterName("Arith", new(Arith))
	server.Serve("tcp", "127.0.0.1:7890")
}
