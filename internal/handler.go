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

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		message := scanner.Text()

		ack := strings.ToUpper(strings.TrimSpace(message))

		response := fmt.Sprintf("ACK: %s\n", ack)

		_, err := conn.Write([]byte(response))

		if err != nil {
			log.Printf("Server write error: %v", err)
			return
		}
	}

}
