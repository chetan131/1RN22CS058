package main

import (
	"log"
	"net/http"
)

func main() {
	router := setupRouter()
	log.Println("Server running on port 9876...")
	log.Fatal(http.ListenAndServe(":9876", router))
}

