Great! Let's start by **creating the MySQL database** and setting up our **Go project to connect to it**. 🚀

---

## 🏗 **Step 1: Create a MySQL Database**

1️⃣ **Start MySQL Server**  
Make sure MySQL is running on your system. If you’re using **Docker**, you can start a MySQL container like this:

```sh
docker run --name mysql-go -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=go_db -p 3306:3306 -d mysql:latest
```

If you're using a **local MySQL installation**, log in to MySQL:

```sh
mysql -u root -p
```

2️⃣ **Create the Database**

Once inside the MySQL shell, create a database:

```sql
CREATE DATABASE go_db;
```

3️⃣ **Create a Table for Objects**

Let's assume we are storing **objects** with an `id`, `name`, and `description`. Run this SQL command:

```sql
USE go_db;

CREATE TABLE objects (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

4️⃣ **Verify the Table**

To check if the table was created successfully:

```sql
SHOW TABLES;
```

---

## 🛠 **Step 2: Set Up Database Connection in Go**

Now, let's write Go code to connect to this database.

1️⃣ **Install the MySQL Driver**

Run this command in your project folder:

```sh
go get -u github.com/go-sql-driver/mysql
```

2️⃣ **Create a `config/config.go` File**

This file will handle our database connection.

📄 **`config/config.go`**

```go
package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// DB is the global database connection
var DB *sql.DB

// InitDB initializes the database connection
func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

	var dbErr error
	DB, dbErr = sql.Open("mysql", dsn)
	if dbErr != nil {
		log.Fatal("❌ Failed to connect to database:", dbErr)
	}

	// Check the connection
	if err = DB.Ping(); err != nil {
		log.Fatal("❌ Database connection failed:", err)
	}

	log.Println("✅ Successfully connected to MySQL database")
}
```

3️⃣ **Add a `.env` File**

This file will store our database credentials.

📄 **`.env`**

```
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=root
DB_NAME=go_db
```

4️⃣ **Initialize Database Connection in `main.go`**

Now, let’s modify our **`cmd/main.go`** file to initialize the database when we start the server.

📄 **`cmd/main.go`**

```go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/your-username/go-rest-api/config"
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
```

---

## 🎯 **Step 3: Run and Test**

1️⃣ **Start the Server**

```sh
go run cmd/main.go
```

If everything is correct, you should see:

```
✅ Successfully connected to MySQL database
🚀 Server started on http://localhost:8080
```

2️⃣ **Test the API**

Open your browser or use `curl` to check the API:

```sh
curl http://localhost:8080
```

You should see this response:

```
🚀 Go REST API is running!
```

---
🔥 **Awesome! Your database connection is working!** 🎉🚀

Now, let's move to the next step: **implementing CRUD operations**!

---

## ✨ **Step 1: Define the Object Model**

Since we are storing objects in MySQL, let's define a **Go struct** for our database table.

📄 **`models/object.go`**

```go
package models

import "time"

// Object represents the database table structure
type Object struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
```

This **matches** our MySQL table:

```sql
CREATE TABLE objects (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

---

## ✨ **Step 2: Create Repository for Database Operations**

A **repository** helps us separate database logic from the main application.

📄 **`repository/object_repository.go`**

```go
package repository

import (
	"database/sql"
	"log"

	"github.com/your-username/go-rest-api/config"
	"github.com/your-username/go-rest-api/models"
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
```

---

## ✨ **Step 3: Create Handlers for API Routes**

Now, let's create **handlers** to process API requests.

📄 **`handlers/object_handler.go`**

```go
package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/your-username/go-rest-api/models"
	"github.com/your-username/go-rest-api/repository"
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
```

---

## ✨ **Step 4: Set Up Routes**

📄 **`routes/routes.go`**

```go
package routes

import (
	"net/http"

	"github.com/your-username/go-rest-api/handlers"
)

// SetupRoutes defines API routes
func SetupRoutes() {
	http.HandleFunc("/objects", handlers.GetAllObjectsHandler)
	http.HandleFunc("/object", handlers.GetObjectByIDHandler)
	http.HandleFunc("/create", handlers.CreateObjectHandler)
	http.HandleFunc("/update", handlers.UpdateObjectHandler)
	http.HandleFunc("/delete", handlers.DeleteObjectHandler)
}
```

---

## ✨ **Step 5: Update `main.go` to Include Routes**

📄 **`cmd/main.go`**

```go
package main

import (
	"log"
	"net/http"

	"github.com/your-username/go-rest-api/config"
	"github.com/your-username/go-rest-api/routes"
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
```

---

## ✅ **Step 6: Run & Test Your API**

### Start the server
```sh
go run cmd/main.go
```

### Test API Endpoints

#### ➕ Create an Object
```sh
curl -X POST http://localhost:8080/create -d '{"name":"Laptop","description":"MacBook Air M1"}' -H "Content-Type: application/json"
```

#### 📜 Get All Objects
```sh
curl http://localhost:8080/objects
```

#### 🔄 Update an Object
```sh
curl -X PUT http://localhost:8080/update -d '{"id":1,"name":"Updated Laptop","description":"Updated Description"}' -H "Content-Type: application/json"
```

#### ❌ Delete an Object
```sh
curl -X DELETE http://localhost:8080/delete?id=1
```

🚀 Let me know if you have any issues! feel free to open a PULL request😃