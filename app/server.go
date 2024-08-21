package main

import (
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/app/commands"
	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	defer func() {
		err = l.Close()
		if err != nil {
			fmt.Println("Failed to close listener: " + err.Error())
		}
	}()
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	for {
		go func() {
			c, err := l.Accept()
			if err != nil {
				fmt.Println("Error accepting connection: ", err.Error())
				os.Exit(1)
			}
			for {
				go handleConn(c)
			}
		}()
	}
}

func handleConn(conn net.Conn) {
	defer func() {
		if err := conn.Close(); err != nil {
			fmt.Println("Error closing connection: ", err.Error())
		}
	}()

	input := make([]byte, 1024)
	if _, err := conn.Read(input); err != nil {
		fmt.Println("Error reading from connection: ", err.Error())
	}

	output, err := commands.Execute(string(input))
	if err != nil {
		fmt.Println("Error executing command: ", err.Error())
	}

	if _, err := conn.Write([]byte(output)); err != nil {
		fmt.Println("Error writing to connection: ", err.Error())
	}
}
