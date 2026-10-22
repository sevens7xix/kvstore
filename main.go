package main

import (
	handler "kvstore/internal"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":6380")

	if err != nil {
		log.Fatal("Error listening TCP port 6380:", err)
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Println("error accepting conn:", err)
			continue
		}

		go handler.HandleConn(conn)
	}
}
