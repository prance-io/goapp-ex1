package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	defaultListenAddress = ":9000"
)

func main() {
	port := os.Getenv("LISTEN_ADDRESS")
	if port == "" {
		port = defaultListenAddress
	}
	log.Printf("Listening on: %v", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi!\n")
	})

	log.Fatal(http.ListenAndServe(port, nil))
}
