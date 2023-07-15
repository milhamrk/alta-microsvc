package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	// Retrieve user ID from request parameters
	vars := mux.Vars(r)
	userID := vars["id"]

	// Fetch user from the database or any other data source
	// Replace this with your actual logic
	user := fetchUserFromDatabase(userID)

	// Return user information as JSON response
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"id":"%s", "name":"%s", "email":"%s"}`, user.ID, user.Name, user.Email)
}

func fetchUserFromDatabase(userID string) *User {
	// Simulated function to fetch user from the database
	// Replace this with your actual logic
	return &User{
		ID:    userID,
		Name:  "John Doe",
		Email: "john@example.com",
	}
}

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	router := mux.NewRouter()

	// Define the route for getting a user by ID
	router.HandleFunc("/users/{id}", GetUser).Methods("GET")

	// Start the server on port 8080
	log.Fatal(http.ListenAndServe(":8080", router))
}
