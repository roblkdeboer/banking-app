package main

import (
	"fmt"
	"net/http"
	"os"

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

	fmt.Fprint(w, data)
}

func getHello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello world\n")
}

func main() {
	db.InitDB()
	defer db.Connection.Close()

	http.Handle("/", http.FileServer(http.Dir("./static")))
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