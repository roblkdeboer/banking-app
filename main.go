package main

import (
	"fmt"
	"net/http"
)

func getData(w http.ResponseWriter, req *http.Request) {
  fmt.Fprintf(w, "Hello world\n")
}

func main() {	
	http.HandleFunc("/data", getData)

	http.ListenAndServeTLS(
		":443",
		"/etc/letsencrypt/live/bankapi.roblkdeboer.com/fullchain.pem",
		"/etc/letsencrypt/live/bankapi.roblkdeboer.com/privkey.pem",
		nil,
	  )
}