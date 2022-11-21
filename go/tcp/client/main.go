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

	var data = []byte{'h', 'e', 'l', 'l', 'o'}
	conn.Write(data)

}
