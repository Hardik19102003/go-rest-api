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

	// Debugging: Print DSN
	fmt.Println("🔍 DSN:", dsn)

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
