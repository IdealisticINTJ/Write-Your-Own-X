package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(os.Stdout, "Listening for connections on 127.0.0.1:8080...")

	go checkHealthyServers()
	acceptRequests(ln)
}

func acceptRequests(ln net.Listener) {
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleConnection(conn)
		fmt.Println()
	}
}
	
