package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	c, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	defer func() {
		if err := c.Close(); err != nil {
			fmt.Println("Error closing connection: ", err.Error())
		}
	}()

	buf := make([]byte, 128)
	for {
		if _, err = c.Read(buf); err != nil {
			fmt.Println("Error reading from connection: ", err.Error())
		}

		if _, err = c.Write([]byte("+PONG\r\n")); err != nil {
			fmt.Println("Error writing to connection: ", err.Error())
		}

		clear(buf)
	}
}
