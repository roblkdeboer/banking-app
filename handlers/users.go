package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/roblkdeboer/banking-app/models"
	db "github.com/roblkdeboer/banking-app/postgres"
	"github.com/roblkdeboer/banking-app/users"
	"github.com/roblkdeboer/banking-app/utils"
)

func GetUsers(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Connection.Query("SELECT first_name, last_name FROM users")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	data := ""
	for rows.Next() {
		var firstName, lastName string // Define variables for first_name and last_name
		err = rows.Scan(&firstName, &lastName)
		if err != nil {
			fmt.Println(err)
		}
		fullName := fmt.Sprintf("%s %s", firstName, lastName)
		data += fmt.Sprintf("%s ", fullName)
	}

	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintln(w, data)
}

func SignIn(w http.ResponseWriter, req *http.Request) {
	var user models.UserSignInRequest
    err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	dbUser, err := users.GetUserPassword(db.Connection, user.Email)
	if err != nil {
        http.Error(w, "incorrect Email/Password", http.StatusInternalServerError)
        return
    }

	if !utils.VerifyPassword(user.Password, dbUser.Password) {
        http.Error(w, "incorrect Email/Password", http.StatusInternalServerError)
        return
	}

	// Generate a JWT for the authenticated user
    tokenString, err := utils.GenerateJWTToken(dbUser.ID)
    if err != nil {
        http.Error(w, "failed to generate JWT", http.StatusInternalServerError)
        return
    }

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Authorization", "Bearer "+tokenString)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successfully logged in"))
}

func SignUp(w http.ResponseWriter, req *http.Request) {
    var user models.User
    err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	if user.FirstName == "" || user.LastName == "" || user.Phone == "" || user.Email == "" || user.Password == "" {
        http.Error(w, "missing user fields.", http.StatusBadRequest)
		return
    }

	if !isValidEmail(user.Email) {
        http.Error(w, "invalid email format.", http.StatusBadRequest)
        return
    }

    exists, err := users.UserExists(db.Connection, user.Email)
    if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
    }
    if exists {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("user already exists"))
        return
    }

	passwordHash, err := utils.GeneratePasswordHash(user.Password)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    err = users.InsertUser(db.Connection, user, passwordHash)
	if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
	w.Write([]byte("user created successfully"))
}

func isValidEmail(email string) bool {
    return len(email) > 0 && strings.Contains(email, "@")
}