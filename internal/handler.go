package handler

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func HandleConn(conn net.Conn) {

	defer conn.Close()

	reader := bufio.NewReader(conn)
	message, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Bufio reader error: %v", err)
		return
	}

	ACK := strings.ToUpper(strings.TrimSpace(message))

	response := fmt.Sprintf("ACK: %s\n", ACK)

	_, err = conn.Write([]byte(response))

	if err != nil {
		log.Printf("Server write error: %v", err)
		return
	}
}
