package main

import (
	"fmt"
	"time"

	"github.com/smallnest/rpcx"
	"github.com/smallnest/rpcx/clientselector"
)

type Args struct {
	A int `msg:"a"`
	B int `msg:"b"`
}

type Reply struct {
	C int `msg:"c"`
}

func main() {
	server1 := clientselector.ServerPair{Network: "tcp", Address: "127.0.0.1:7890"}
	server2 := clientselector.ServerPair{Network: "tcp", Address: "127.0.0.1:7891"}

	servers := []clientselector.ServerPair{server1, server2}

	s := clientselector.NewMultiClientSelector(servers, rpcx.RandomSelect, 10*time.Second)

	for i := 0; i < 10; i++ {
		callServer(s)
	}
}

func callServer(s rpcx.ClientSelector) {
	client := rpcx.NewClient(s)

	args := &Args{7, 8}
	var reply Reply
	err := client.Call("Arith.Mul", args, &reply)
	if err != nil {
		fmt.Printf("error for Arith: %d*%d, %v \n", args.A, args.B, err)
	} else {
		fmt.Printf("Arith: %d*%d=%d \n", args.A, args.B, reply.C)
	}

	client.Close()
}
