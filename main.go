package main

import (
	"bytes"
	"fmt"
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

	reader := bytes.NewReader(response)

	// Get the header
	header, err := message.ParseHeader(reader)
	if err != nil {
		log.Fatal(err)
	}

	// Print the header
	fmt.Printf("%+v\n", header)

	// Get the questions
	question, err := message.ParseQuestion(reader)
	if err != nil {
		log.Fatal(err)
	}

	// Print the questions
	fmt.Printf("%+v\n", question)

}
