package main

import "net/http"

func setupRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/numbers/", handleNumberRequest)
	return mux
}
