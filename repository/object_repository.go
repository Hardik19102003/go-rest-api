package repository

import (
	"database/sql"
	"log"

	"github.com/Hardik19102003/go-rest-api/config"
	"github.com/Hardik19102003/go-rest-api/models"
)

// CreateObject inserts a new object into the database
func CreateObject(obj models.Object) (int64, error) {
	query := "INSERT INTO objects (name, description) VALUES (?, ?)"
	result, err := config.DB.Exec(query, obj.Name, obj.Description)
	if err != nil {
		log.Println("❌ Error inserting object:", err)
		return 0, err
	}
	return result.LastInsertId()
}

// GetAllObjects retrieves all objects from the database
func GetAllObjects() ([]models.Object, error) {
	query := "SELECT id, name, description, created_at FROM objects"
	rows, err := config.DB.Query(query)
	if err != nil {
		log.Println("❌ Error fetching objects:", err)
		return nil, err
	}
	defer rows.Close()

	var objects []models.Object
	for rows.Next() {
		var obj models.Object
		err := rows.Scan(&obj.ID, &obj.Name, &obj.Description, &obj.CreatedAt)
		if err != nil {
			log.Println("❌ Error scanning row:", err)
			return nil, err
		}
		objects = append(objects, obj)
	}
	return objects, nil
}

// GetObjectByID retrieves a single object by ID
func GetObjectByID(id int) (models.Object, error) {
	query := "SELECT id, name, description, created_at FROM objects WHERE id = ?"
	var obj models.Object
	err := config.DB.QueryRow(query, id).Scan(&obj.ID, &obj.Name, &obj.Description, &obj.CreatedAt)
	if err != nil {
		log.Println("❌ Error fetching object by ID:", err)
		return models.Object{}, err
	}
	return obj, nil
}

// UpdateObject updates an existing object
func UpdateObject(obj models.Object) error {
	query := "UPDATE objects SET name = ?, description = ? WHERE id = ?"
	_, err := config.DB.Exec(query, obj.Name, obj.Description, obj.ID)
	if err != nil {
		log.Println("❌ Error updating object:", err)
	}
	return err
}

// DeleteObject deletes an object from the database
func DeleteObject(id int) error {
	query := "DELETE FROM objects WHERE id = ?"
	_, err := config.DB.Exec(query, id)
	if err != nil {
		log.Println("❌ Error deleting object:", err)
	}
	return err
}
