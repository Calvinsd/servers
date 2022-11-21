package main

import (
	"log"
	"net"
)

func main() {
	s, err := net.ResolveUDPAddr("udp4", ":8080")

	if err != nil {
		log.Fatal((err))
	}

	c, err := net.DialUDP("udp4", nil, s)

	if err != nil {
		log.Fatal((err))
	}

	var data = []byte{'h', 'e', 'l', 'l', 'o'}
	c.Write(data)

}
