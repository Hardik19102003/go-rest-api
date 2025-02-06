package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Hardik19102003/go-rest-api/models"
	"github.com/Hardik19102003/go-rest-api/repository"
)

// CreateObjectHandler handles object creation
func CreateObjectHandler(w http.ResponseWriter, r *http.Request) {
	var obj models.Object
	err := json.NewDecoder(r.Body).Decode(&obj)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	id, err := repository.CreateObject(obj)
	if err != nil {
		http.Error(w, "Failed to create object", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int64{"id": id})
}

// GetAllObjectsHandler retrieves all objects
func GetAllObjectsHandler(w http.ResponseWriter, r *http.Request) {
	objects, err := repository.GetAllObjects()
	if err != nil {
		http.Error(w, "Failed to fetch objects", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(objects)
}

// GetObjectByIDHandler retrieves a specific object
func GetObjectByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	obj, err := repository.GetObjectByID(id)
	if err != nil {
		http.Error(w, "Object not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(obj)
}

// UpdateObjectHandler updates an object
func UpdateObjectHandler(w http.ResponseWriter, r *http.Request) {
	var obj models.Object
	err := json.NewDecoder(r.Body).Decode(&obj)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = repository.UpdateObject(obj)
	if err != nil {
		http.Error(w, "Failed to update object", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Object updated successfully"})
}

// DeleteObjectHandler deletes an object
func DeleteObjectHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = repository.DeleteObject(id)
	if err != nil {
		http.Error(w, "Failed to delete object", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Object deleted successfully"})
}
