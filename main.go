package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	defaultPort = "9000"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	log.Printf("Listening on port: %v", port)
	listenInterface := fmt.Sprintf(":%s", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi!\n")
	})

	log.Fatal(http.ListenAndServe(listenInterface, nil))
}
