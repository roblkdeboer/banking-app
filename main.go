package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	models "github.com/roblkdeboer/banking-app/models"
	db "github.com/roblkdeboer/banking-app/postgres"
)

func getUsers(w http.ResponseWriter, req *http.Request) {
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

func createUser(w http.ResponseWriter, req *http.Request) {
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

	// Insert the user data into the database
    statement := "INSERT INTO users (first_name, last_name, phone, email, password) VALUES ($1, $2, $3, $4, $5)"
    _, err = db.Connection.Exec(statement, user.FirstName, user.LastName, user.Phone, user.Email, user.Password)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

	// Return a success response
    w.WriteHeader(http.StatusCreated)
    fmt.Fprintln(w, "User created successfully")
}

func isValidEmail(email string) bool {
    return len(email) > 0 && strings.Contains(email, "@")
}

func getHello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello world")
}

func main() {
	db.InitDB()
	defer db.Connection.Close()

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/create-user", createUser)
	http.HandleFunc("/users", getUsers);
	http.HandleFunc("/hello", getHello);


	serverEnv := os.Getenv("SERVER_ENV")

	if serverEnv == "DEV" {
		http.ListenAndServe(":8080", nil)
	} else if serverEnv == "PROD" {
		http.ListenAndServeTLS(
			":443",
			"/etc/letsencrypt/live/banking.roblkdeboer.com/fullchain.pem",
			"/etc/letsencrypt/live/banking.roblkdeboer.com/privkey.pem",
			nil,
		)
	}
}