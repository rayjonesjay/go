package main

import (
	"fmt"
	"net"
//	"time"
	"log"
	"net/smtp"
)

func sendEmail(notification string) {
	from := "rjmuiruri@gmail.com"
	password := "Nemesis???254"
	to := "rayjaymuiruri@gmail.com"

	smtpHost := "smtp.example.com"
	smtpPort := "587"

	message := []byte("Subject: Connection Alert\n\n" + notification)

	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, message)
	if err != nil {
		log.Printf("Failed to send email: %s\n", err)
	} else {
		fmt.Println("Email sent successfully!")
	}
}

func notifyConnection(addr string) {
	fmt.Printf("Connection detected from: %s\n", addr)
	sendEmail("Connection detected from: " + addr)
}

func main() {
	port := ":8080" // Change to the port you want to monitor
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start listener: %s\n", err)
	}
	defer listener.Close()

	fmt.Printf("Monitoring connections on port %s...\n", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %s\n", err)
			continue
		}
		go func(c net.Conn) {
			addr := c.RemoteAddr().String()
			notifyConnection(addr)
			c.Close()
		}(conn)
	}
}

