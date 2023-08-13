package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Args struct{}

type PingService struct {
	ResponseData string
}

func (p PingService) Ping(args *Args, reply *PingService) error {
	*&reply.ResponseData = "pong"

	return nil
}

func main() {
	pingService := new(PingService)

	rpc.Register(pingService)

	rpc.HandleHTTP()

	// Start listening for the requests on port 1234
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Listener error: ", err)
	}

	http.Serve(listener, nil)
}
