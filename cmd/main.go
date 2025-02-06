package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Hardik19102003/go-rest-api/config"
)

func main() {
	// Initialize the database
	config.InitDB()

	// Temporary test route to check if API is running
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "🚀 Go REST API is running!")
	})

	// Start server
	port := ":8080"
	log.Println("🚀 Server started on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
