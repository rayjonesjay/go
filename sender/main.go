package main

import (
	"fmt"
	"net"
)

func main() {
	message := ""
	for {
		fmt.Scan(&message)
		slice := []byte(message)
		// sending to a TCP server
		netConn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			fmt.Println("Error connecting...", err)
			return
		}
		fmt.Println("listening")
		defer netConn.Close()

		netConn.Write(slice)
		fmt.Printf("sent >> %q\n", message)
	}
}
