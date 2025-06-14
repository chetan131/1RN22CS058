package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
)

var numberWindow = NewNumberWindow(10)

func handleNumberRequest(w http.ResponseWriter, r *http.Request) {
	numberID := strings.TrimPrefix(r.URL.Path, "/numbers/")
	apiURL := getAPIURL(numberID)

	log.Println("Received request for:", numberID)
	log.Println("Mapped external URL:", apiURL)

	if apiURL == "" {
		http.Error(w, "Invalid number ID", http.StatusBadRequest)
		return
	}

	client := http.Client{Timeout: 500 * time.Millisecond}
	numbers, err := fetchNumbersFromAPI(client, apiURL)
	if err != nil {
		http.Error(w, "Failed to fetch numbers", http.StatusServiceUnavailable)
		return
	}

	prev := numberWindow.CurrentState()
	numberWindow.Update(numbers)
	curr := numberWindow.CurrentState()
	avg := numberWindow.Average()

	response := OutputResponse{
		WindowPrevState: prev,
		WindowCurrState: curr,
		Numbers:         numbers,
		Avg:             avg,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
