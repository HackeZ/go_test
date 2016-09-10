package main

import (
	"fmt"
	"log"
	"time"

	"github.com/smallnest/rpcx"
)

// Args ..
type Args struct {
	A int `msg:"a"`
	B int `msg:"b"`
}

// Reply ..
type Reply struct {
	C int `msg:"c"`
}

func main() {
	// init rpcx selector and client.
	s := &rpcx.DirectClientSelector{Network: "tcp", Address: "127.0.0.1:7890", DialTimeout: 10 * time.Second}
	client := rpcx.NewClient(s)
	defer client.Close()

	args := &Args{A: 7, B: 8}
	var reply Reply
	divCall := client.Go("Arith.Mul", args, &reply, nil)

	replyCall := <-divCall.Done // will be equal to divCall
	if replyCall.Error != nil {
		log.Printf("error for Arith: %d*%d, %v \n", args.A, args.B, replyCall.Error)
	}
	fmt.Printf("Arith: %d*%d=%d \n", args.A, args.B, reply.C)
}
