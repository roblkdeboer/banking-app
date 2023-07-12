package main

import (
	"fmt"
	"net/http"
	"os"
)

func getData(w http.ResponseWriter, req *http.Request) {
  fmt.Fprintf(w, "Hello world\n")
}

func main() {
	http.HandleFunc("/data", getData);

	serverEnv := os.Getenv("SERVER_ENV")

	if serverEnv == "DEV" {
		http.ListenAndServe(":8080", nil)
	} else if serverEnv == "PROD" {
		http.ListenAndServeTLS(
			":443",
			"/etc/letsencrypt/live/bankapi.roblkdeboer.com/fullchain.pem",
			"/etc/letsencrypt/live/bankapi.roblkdeboer.com/privkey.pem",
			nil,
		)
	}
}