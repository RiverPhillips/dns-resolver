package main

import (
	"log"
	"net"

	"github.com/RiverPhillips/dns-resolver/internal/message"
)

func main() {
	query := message.BuildQuery("google.com", message.TYPE_A)

	// Create a UDP socket
	conn, err := net.Dial("udp", "1.1.1.1:53")
	if err != nil {
		log.Fatal(err)
	}

	// Send the query
	if _, err := conn.Write(query); err != nil {
		log.Fatal(err)
	}

	// Receive the response
	response := make([]byte, 1024)
	if _, err := conn.Read(response); err != nil {
		log.Fatal(err)
	}
}
