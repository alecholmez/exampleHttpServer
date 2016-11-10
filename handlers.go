package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Index ...
func Index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(w, "Hello World!")
}

// ListUsers ...
func ListUsers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var user User
	user.Name = "Alec"
	user.Email = "alecholmez@gmail.com"
	user.ID = GenID(user.Email)

	bytes, error := json.MarshalIndent(user, "", "    ")
	if error != nil {
		log.Fatal(error)
	}

	// Write the user json object to the browser response
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	w.Write(bytes)
}

// CreateUser ...
func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
	}

	if user.Email == "" {
		log.Println("Can't create an ID")
	} else {
		user.ID = GenID(user.Email)
	}

	bytes, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	w.Write(bytes)
}

// GetUser ...
func GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userID := p.ByName("id")
	var users Users
	var user User

	for _, u := range users {
		if u.ID == userID {
			fmt.Println("Found a match")
			user = u
		}
	}

	bytes, err := json.MarshalIndent(user, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	w.Write(bytes)
}
