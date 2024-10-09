package client

import (
	"fmt"
	"gorpc/server"
	"log"
	"net/rpc"
)

func Client() {
	serverAddress := "localhost"
	client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
	if err != nil {
		log.Fatal("Can not dial to server", err)

	}
	//synchronous call

	args := &server.Args{A: 20, B: 10}

	var reply int
	err = client.Call("Arith.Multiply", args, &reply)

	if err != nil {
		log.Fatal("arith server unavailable", err)
	}
	fmt.Printf("Arith: %d %d %d is", args.A, args.B, reply)

}
