package main

import (
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/app/commands"
	"github.com/codecrafters-io/redis-starter-go/app/lib"
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
			handleConn(lib.NewSafeConn(c))
		}()
	}
}

func handleConn(conn *lib.SafeConn) {
	defer func() {
		if err := conn.Close(); err != nil {
			fmt.Println("Error closing connection: ", err.Error())
		}
	}()

	executor := commands.NewExecutor()

	for {
		go func() {
			input := make([]byte, 128)

			if _, err := conn.Read(input); err != nil {
				fmt.Println("Error reading from connection: ", err.Error())
			}

			output, err := executor.Execute(string(input))
			if err != nil {
				fmt.Println("Error executing command: ", err.Error())
			}

			if _, err = conn.Write([]byte(output)); err != nil {
				fmt.Println("Error writing to connection: ", err.Error())
			}
		}()

	}
}
