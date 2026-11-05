package handler

import (
	"bufio"
	"fmt"
	"kvstore/internal/store"
	"log"
	"net"
	"strings"
)

const (
	ErrWrongUseOfCommands = "ERR Wrong use of command"
)

func HandleConn(store store.Store, conn net.Conn) {

	defer conn.Close()

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		message := scanner.Text()

		messageCommands := strings.Split(strings.TrimSpace(message), " ")

		var response string

		if len(messageCommands) > 3 {
			response = fmt.Sprintln(ErrWrongUseOfCommands)
		} else {
			switch messageCommands[0] {
			case "SET":
			case "set":
				if len(messageCommands) != 3 {
					response = fmt.Sprintln(ErrWrongUseOfCommands)
					break
				}
				store.Set(messageCommands[1], messageCommands[2])
				response = fmt.Sprintln("OK")
			case "GET":
			case "get":
				if len(messageCommands) != 2 {
					response = fmt.Sprintln(ErrWrongUseOfCommands)
					break
				}
				value, err := store.Get(messageCommands[1])

				if err != nil {
					response = fmt.Sprintf("ERR %v", err.Error())
					break
				}

				response = fmt.Sprintf("GET: %s\n", value)
			case "DELETE":
			case "delete":
				if len(messageCommands) != 2 {
					response = fmt.Sprintln(ErrWrongUseOfCommands)
					break
				}

				store.Delete(messageCommands[1])
				response = fmt.Sprintln("OK")
			default:
				response = fmt.Sprintln("ERR There is no command for that")
			}
		}

		_, err := conn.Write([]byte(response))

		if err != nil {
			log.Printf("Server write error: %v", err)
			return
		}
	}

}
