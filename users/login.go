package users

import (
	"database/sql"

	"github.com/roblkdeboer/banking-app/models"
)

func GetUserPassword(db *sql.DB, email string) (*models.User, error) {
	var user models.User
	err := db.QueryRow("SELECT id, email, password FROM users WHERE email=$1", email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}