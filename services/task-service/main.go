package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Task Service is running!")
	})

	log.Printf("Task service starting on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
