package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type ServerWriter struct {
	serverURL string
}

func NewServerWriter(serverURL string) *ServerWriter {
	return &ServerWriter{serverURL: serverURL}
}

func (sw *ServerWriter) Write(p []byte) (n int, err error) {
	resp, err := http.Post(sw.serverURL, "text/plain", bytes.NewReader(p))
	if err != nil {
		return 0, fmt.Errorf("failed to send log to server %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("server returned non200 status %d", resp.StatusCode)
	}

	return len(p), nil
}

func main() {
	serverUrl := "http://localhost:8080/log"

	serverWriter := NewServerWriter(serverUrl)

	logger := log.New(serverWriter, "Log: ", log.LstdFlags|log.Llongfile)

	go startMockServer()

	names := []string{"hello", "world", "this is good", "the best"}
	for _, nm := range names {
		time.Sleep(time.Millisecond * 1500)
		logger.Println(nm)
	}
	fmt.Println("hello world")
}

func startMockServer() {
	http.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		logData, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "failed to read log data", http.StatusInternalServerError)
		}
		fmt.Printf("received log %s", logData)
		w.WriteHeader(http.StatusOK)
	})

	fmt.Println("mock server running on http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("failed to start server")
	}
}
