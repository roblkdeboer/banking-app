package main

import (
	"fmt"
	"log"
	"net/http"
)

func getData(w http.ResponseWriter, req *http.Request) {
  fmt.Fprintf(w, "Hello world\n")
}

func main() {
	addr := ":8080"
	
	http.HandleFunc("/data", getData)

	log.Println("Starting server on port", addr)
    log.Fatal( http.ListenAndServe(addr, nil) )
}