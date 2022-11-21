package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", "localhost:8080")

	if err != nil {
		log.Fatal((err))
	}

	fmt.Println("Started tcp server on port 8080")
	conn, err := ln.Accept()

	defer conn.Close()

	if err != nil {
		log.Fatal((err))
	}

	handleConnection(conn)
}

func handleConnection(conn net.Conn) {
	var data = make([]byte, 1024)
	n, err := conn.Read(data)

	if err != nil {
		log.Fatal((err))
	}

	fmt.Println(n, string(data))

	conn.Close()
}
