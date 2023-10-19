package main

import (
	"fmt"
	"net/http"
	"os"

	handlers "github.com/roblkdeboer/banking-app/handlers"
	db "github.com/roblkdeboer/banking-app/postgres"
)


func getHello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello world")
}

func main() {
	db.InitDB()
	defer db.Connection.Close()

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/create-user", handlers.CreateUser)
	http.HandleFunc("/users", handlers.GetUsers);
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