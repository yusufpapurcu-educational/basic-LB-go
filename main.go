package main

import (
	"io"
	"log"
	"net"
)

var (
	listenAddr       = "localhost:8080"
	availableServers = []string{"localhost:5001", "localhost:5002", "localhost:5003"}
	counter          = 0
)

func main() {
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalf("Failed listening : %s", err)
	}

	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("Failed to accepting conn : %s", err)
		}

		go proxy(conn)
	}
}

func proxy(conn net.Conn) {
	selectedAddr := availableServers[counter%len(availableServers)]
	counter++
	log.Printf("\nSelected Server :%s\n\n", selectedAddr)

	server, err := net.Dial("tcp", selectedAddr)
	if err != nil {
		log.Fatalf("Failed to connecting server : %s", err)
	}
	go io.Copy(conn, server)
	go io.Copy(server, conn)
}
