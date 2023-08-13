package main

import (
	"log"
	"net"
	"time"
)

func main() {
	ln, err := net.Listen("tcp", "localhost:8080")

	if err != nil {
		log.Fatal((err))
	}

	log.Println("Started tcp server on port 8080")

	for {
		log.Println("Accepting connection.....")
		conn, err := ln.Accept()

		if err != nil {
			log.Fatal((err))
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {

	var data = make([]byte, 1024)

	n, err := conn.Read(data)

	if err != nil {
		log.Fatal((err))
	}

	log.Println(n, string(data))

	time.Sleep(5 * time.Second)

	conn.Close()
}
