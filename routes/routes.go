package routes

import (
	"net/http"

	"github.com/Hardik19102003/go-rest-api/handlers"
)

// SetupRoutes defines API routes
func SetupRoutes() {
	http.HandleFunc("/objects", handlers.GetAllObjectsHandler)
	http.HandleFunc("/object", handlers.GetObjectByIDHandler)
	http.HandleFunc("/create", handlers.CreateObjectHandler)
	http.HandleFunc("/update", handlers.UpdateObjectHandler)
	http.HandleFunc("/delete", handlers.DeleteObjectHandler)
}
