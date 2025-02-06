package main

import (
	"log"
	"net/http"

	"github.com/Hardik19102003/go-rest-api/config"
	"github.com/Hardik19102003/go-rest-api/routes"
)

func main() {
	// Initialize the database
	config.InitDB()

	// Set up routes
	routes.SetupRoutes()

	// Start server
	port := ":8080"
	log.Println("🚀 Server started on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
