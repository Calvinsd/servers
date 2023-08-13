package main

import (
	"log"
	"net/rpc"
)

type Args struct{}

type PingService struct {
	ResponseData string
}

func main() {

	reply := PingService{}

	args := Args{}

	// DialHTTP connects to an HTTP RPC server at the specified network
	client, err := rpc.DialHTTP("tcp", "localhost"+":8080")
	if err != nil {
		log.Fatal("Client connection error: ", err)
	}

	// Invoke the remote function GiveServerTime attached to TimeServer pointer
	// Sending the arguments and reply variable address to the server as well
	err = client.Call("PingService.Ping", args, &reply)
	if err != nil {
		log.Fatal("Client invocation error: ", err)
	}

	// Print the reply from the server
	log.Printf("%v", reply)
}
