package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Response struct {
	Message string `json:"message"`
}

type Request struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserInfo struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Location string `json:"location"`
}

var counter int
var mu sync.Mutex

// API route for GET request (for the initial "Hello from Go API!" message)
func getMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := Response{Message: "Hello from Go API!"}
	json.NewEncoder(w).Encode(resp)
}

// API route for POST request (send a custom message)
func sendCustomMessage(w http.ResponseWriter, r *http.Request) {
	var req Request
	// Parse the incoming JSON body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create a response message
	resp := Response{
		Message: fmt.Sprintf("Hello, %s! Your age is %d. Your message is received!", req.Name, req.Age),
	}

	// Send the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// API route for GET request (return user info based on name and age)
func getUserInfo(w http.ResponseWriter, r *http.Request) {
	// Query parameters: "name" and "location"
	name := r.URL.Query().Get("name")
	location := r.URL.Query().Get("location")

	// Example of user information
	user := UserInfo{
		Name:     name,
		Age:      25, // You could extend this by getting data from a database
		Location: location,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// API route to track how many messages have been sent (simple counter)
func getMessageCount(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	counter++
	mu.Unlock()

	// Return the current count
	w.Header().Set("Content-Type", "application/json")
	resp := Response{
		Message: fmt.Sprintf("This message has been sent %d times", counter),
	}
	json.NewEncoder(w).Encode(resp)
}

func main() {
	// API route for GET request (for the initial "Hello from Go API!" message)
	http.HandleFunc("/api/message", getMessage)

	// API route for POST request (send a custom message)
	http.HandleFunc("/api/custom-message", sendCustomMessage)

	// API route for GET request (user info)
	http.HandleFunc("/api/user-info", getUserInfo)

	// API route to track the message count
	http.HandleFunc("/api/message-count", getMessageCount)

	// Serve static files for the front-end
	http.Handle("/", http.FileServer(http.Dir("../frontend/public")))

	// Start server on port 8080
	port := ":8080"
	fmt.Printf("Server is running at http://localhost%s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
