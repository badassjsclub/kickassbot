package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	VERSION = "0.0.1"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	log.Printf("Kickass Bot (v%s) is starting on port %s\n", VERSION, port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Kickass Bot (v%s) is in service.", VERSION)
	})
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
