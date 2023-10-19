package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/roblkdeboer/banking-app/models"
	db "github.com/roblkdeboer/banking-app/postgres"
	"github.com/roblkdeboer/banking-app/users"
)

func SignUp(w http.ResponseWriter, req *http.Request) {
	// Parse the request body to extract user data
    var user models.User
    err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	if !isValidEmail(user.Email) {
        http.Error(w, "Invalid Email format.", http.StatusBadRequest)
        return
    }

	// Call the InsertUser function from the users package
    if err := users.InsertUser(db.Connection, user); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

	// Return a success response
    w.WriteHeader(http.StatusCreated)
    fmt.Fprintln(w, "User created successfully")
}

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
		// fmt.Println(fullName)
		data += fmt.Sprintf("%s ", fullName)
	}

	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintln(w, data)
}

func isValidEmail(email string) bool {
    return len(email) > 0 && strings.Contains(email, "@")
}