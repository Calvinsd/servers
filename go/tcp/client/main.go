package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")

	if err != nil {
		log.Fatal((err))
	}

	log.Println("Sending hello")
	conn.Write([]byte{'h', 'e', 'l', 'l', 'o'})
	conn.Close()

	conn, err = net.Dial("tcp", "127.0.0.1:8080")

	if err != nil {
		log.Fatal((err))
	}

	log.Println("Sending world")

	conn.Write([]byte{'w', 'o', 'r', 'l', 'd'})

	conn.Close()

}
