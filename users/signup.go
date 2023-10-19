package users

import (
	"database/sql"

	"github.com/roblkdeboer/banking-app/models"
)

func InsertUser(db *sql.DB, user models.User, passwordHash string) error {
    statement := "INSERT INTO users (first_name, last_name, phone, email, password) VALUES ($1, $2, $3, $4, $5)"
    _, err := db.Exec(statement, user.FirstName, user.LastName, user.Phone, user.Email, passwordHash)
    return err
}

func UserExists(db *sql.DB, email string) (bool, error) {
    var exists bool
    err := db.QueryRow("SELECT EXISTS (SELECT 1 FROM users WHERE email=$1)", email).Scan(&exists)
    if err != nil {
        return false, err
    }
    return exists, nil
}