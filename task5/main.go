package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RequestBody struct {
	Name string `json:"name"`
}

type ResponseBody struct {
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Method check
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode JSON
	var body RequestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil || body.Name == "" {
		http.Error(w, "Invalid JSON format or missing name", http.StatusBadRequest)
		return
	}

	// Response
	resp := ResponseBody{
		Message: fmt.Sprintf("Hello %s", body.Name),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	fmt.Println("--- Task 5: HTTP API ---")
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server on :8080...")
	fmt.Println("Use POST /hello with JSON {\"name\": \"YourName\"}")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Server failed: %v\n", err)
	}
}
