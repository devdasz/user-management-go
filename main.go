package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// User model
type User struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Dob         string `json:"dob"`
	Address     string `json:"address"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
}

// Mock database
var users []User

// getUsers function
func getUsers(w http.ResponseWriter, r *http.Request) {
	// set json content type
	w.Header().Set("Content-Type", "application/json")
	//Encoded the whole list and send
	json.NewEncoder(w).Encode(users)
}

// getUserById function
func getUserById(w http.ResponseWriter, r *http.Request) {
	// set json content type
	w.Header().Set("Content-Type", "application/json")
	// params
	params := mux.Vars(r)
	for _, item := range users {
		if item.ID == params["id"] {
			// Encoded only the match request
			json.NewEncoder(w).Encode(item)
			return
		}

	}
}

// createUser function
func createUser(w http.ResponseWriter, r *http.Request) {
	// set json content type
	w.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.ID = strconv.Itoa(rand.Intn(1000))
	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}

// updateUser function
func updateUser(w http.ResponseWriter, r *http.Request) {
	// set json content type
	w.Header().Set("Content-Type", "application/json")
	// params
	params := mux.Vars(r)
	//loop over the users, range
	//delete the user with the i.d that you've sent
	//add a new user - the user we send in the body of postman/client
	for index, item := range users {
		if item.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			var user User
			_ = json.NewDecoder(r.Body).Decode(&user)
			user.ID = params["id"]
			users = append(users, user)
			json.NewEncoder(w).Encode(user)
			return
		}
	}

}

// deleteUser function
func deleteUser(w http.ResponseWriter, r *http.Request) {
	// set json content type
	w.Header().Set("Content-Type", "application/json")
	// params
	params := mux.Vars(r)
	for index, item := range users {
		if item.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(users)
}

func main() {
	users = append(users,
		User{
			ID:          "23",
			Name:        "test",
			Dob:         "020689",
			Address:     "India",
			Description: "test descrption",
		})
	users = append(users,
		User{
			ID:          "57",
			Name:        "test",
			Dob:         "020689",
			Address:     "India",
			Description: "test descrption",
		})

	r := mux.NewRouter()
	// r.HandleFunc("/movies", getMovies).Methods("GET")

	r.HandleFunc("/v1/users", getUsers).Methods("GET")
	r.HandleFunc("/v1/users/{id}", getUserById).Methods("GET")
	r.HandleFunc("/v1/users", createUser).Methods("POST")
	r.HandleFunc("/v1/users", updateUser).Methods("PUT")
	r.HandleFunc("/v1/users", deleteUser).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))

}
