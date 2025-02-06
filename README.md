
# 🚀 Go REST API with MySQL  

A simple REST API built with **Go** and **MySQL**, following best practices and a clean architecture. This API supports **CRUD (Create, Read, Update, Delete) operations** on stored objects.  

---

## 📌 Project Structure  

```
go-rest-api/
│── cmd/                  # Entry point for the application
│   └── main.go           # Main file to start the server
│
│── config/               # Configuration files (e.g., database settings)
│   └── config.go
│
│── internal/             # Internal business logic (not exposed to external apps)
│   ├── models/           # Database models (structs for tables)
│   │   └── object.go     # Object model (struct definition)
│   │
│   ├── repository/       # Database access layer (interacts with MySQL)
│   │   └── object_repo.go
│   │
│   ├── services/         # Business logic (handles operations)
│   │   └── object_service.go
│   │
│   ├── handlers/         # API handlers (controller layer)
│   │   └── object_handler.go
│
│── routes/               # API routes (define endpoints)
│   └── routes.go
│
│── pkg/                  # Utility functions/helpers
│
│── .env                  # Environment variables (e.g., database credentials)
│── go.mod                # Go module file
│── go.sum                # Dependencies lock file
│── Makefile              # Automate commands (optional)
│── README.md             # Project documentation
```

---

## 🏁 Getting Started  

### 1️⃣ Prerequisites  

Before running the project, make sure you have:  
✅ **Go** installed → [Download Go](https://go.dev/dl/)  
✅ **MySQL** installed and running → [Download MySQL](https://dev.mysql.com/downloads/)  
✅ **Postman (optional)** for testing API requests  

---

### 2️⃣ Clone the Repository  

```sh
git clone https://github.com/your-username/go-rest-api.git
cd go-rest-api
```

---

### 3️⃣ Set Up MySQL Database  

Create a **MySQL database** and update the `.env` file with your credentials:  

```ini
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=yourpassword
DB_NAME=yourdatabase
```

---

### 4️⃣ Install Dependencies  

```sh
go mod tidy
```

---

### 5️⃣ Run the Application  

```sh
go run cmd/main.go
```

Your API will be running on **http://localhost:8080** 🎉  

---

## 🛠 API Endpoints  

| Method | Endpoint      | Description                 |
|--------|-------------|-----------------------------|
| GET    | `/objects`  | Get all objects             |
| GET    | `/objects/{id}` | Get object by ID         |
| POST   | `/objects`  | Create a new object         |
| PUT    | `/objects/{id}` | Update an existing object |
| DELETE | `/objects/{id}` | Delete an object         |

Use **Postman** or `curl` to test the API!  

---

## 📜 License  

This project is licensed under the **MIT License**.  

---
