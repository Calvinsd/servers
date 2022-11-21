package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	s, err := net.ResolveUDPAddr("udp4", ":8080")

	if err != nil {
		log.Fatal((err))
	}

	conn, err := net.ListenUDP("udp4", s)

	if err != nil {
		log.Fatal((err))
	}

	fmt.Println("Started udp server on port 8080")

	defer conn.Close()

	if err != nil {
		log.Fatal((err))
	}

	handleConnection(conn)
}

func handleConnection(conn *net.UDPConn) {
	var data = make([]byte, 1024)
	n, addr, err := conn.ReadFromUDP(data)

	if err != nil {
		log.Fatal((err))
	}

	fmt.Println(n, addr, string(data))

	conn.Close()
}
