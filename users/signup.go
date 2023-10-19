package users

import (
	"database/sql"

	"github.com/roblkdeboer/banking-app/models"
)

func InsertUser(db *sql.DB, user models.User) error {
    statement := "INSERT INTO users (first_name, last_name, phone, email, password) VALUES ($1, $2, $3, $4, $5)"
    _, err := db.Exec(statement, user.FirstName, user.LastName, user.Phone, user.Email, user.Password)
    return err
}