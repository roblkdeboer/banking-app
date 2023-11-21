package utils

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Global DB variable
var Connection *sql.DB

// initDB creates a new instance of DB
func InitDB() {
	err := godotenv.Load(".env") // Load variables from the .env file
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresUser := os.Getenv("POSTGRES_USER")
	
	connStr := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s sslmode=disable",
		postgresUser, postgresPassword, "db", 5432, "postgres",
	)

	var errDB error
	Connection, errDB = sql.Open("postgres", connStr)
	if errDB != nil {
		fmt.Println(err)
	}
}