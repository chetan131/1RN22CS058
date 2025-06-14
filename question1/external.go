package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func getAPIURL(id string) string {
	switch id {
	case "p":
		return "http://20.244.56.144/evaluation-service/primes"
	case "f":
		return "http://20.244.56.144/evaluation-service/fibo"
	case "e":
		return "http://20.244.56.144/evaluation-service/even"
	case "r":
		return "http://20.244.56.144/evaluation-service/rand"
	default:
		return ""
	}
}

func fetchNumbersFromAPI(client http.Client, url string) ([]int, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// 🔒 Replace this with your actual token (expires after some time)
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJNYXBDbGFpbXMiOnsiZXhwIjoxNzQ5ODc4ODY4LCJpYXQiOjE3NDk4Nzg1NjgsImlzcyI6IkFmZm9yZG1lZCIsImp0aSI6IjA0NTQwYzY4LTNmYjAtNDA1OS04NjBmLWRhM2Y3NDk5NGJjMiIsInN1YiI6IjFybjIyY3MwNTgua2NoZXRhbkBybnNpdC5hYy5pbiJ9LCJlbWFpbCI6IjFybjIyY3MwNTgua2NoZXRhbkBybnNpdC5hYy5pbiIsIm5hbWUiOiJrIGNoZXRhbiIsInJvbGxObyI6IjFybjIyY3MwNTgiLCJhY2Nlc3NDb2RlIjoicG1Wc0VoIiwiY2xpZW50SUQiOiIwNDU0MGM2OC0zZmIwLTQwNTktODYwZi1kYTNmNzQ5OTRiYzIiLCJjbGllbnRTZWNyZXQiOiJmUmVkcmRyQWFzQkptSndaIn0.whGQ08PpdG9XlOPklv4KOwhP0_DoMD1mNY31injv0xw")

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Request failed:", err)
		return nil, err
	}
	defer resp.Body.Close()

	var result NumberResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Println("Decode failed:", err)
		return nil, err
	}

	log.Println("Fetched numbers:", result.Numbers)
	return result.Numbers, nil
}
