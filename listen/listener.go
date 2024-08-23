package main

import (
	"fmt"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server...")
		return
	}

	defer l.Close()

	for {
		conn, err := l.Accept()
		messageBuffer := make([]byte, 1024)
		n, err := conn.Read(messageBuffer)
		if err != nil {
			fmt.Println("error reading from connection:", err)
			conn.Close()
			continue
		}

		received := string(messageBuffer[:n])
		fmt.Println("**********")
		fmt.Println("Received message", received)
		conn.Close()
	}
}
