package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// User ..
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

// ServeDocs ...
func ServeDocs(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

// ListUsers ...
func ListUsers(w http.ResponseWriter, r *http.Request) {
	sess := GetMongo(r.Context())

	col := sess.DB("users").C("users")

	var result struct {
		Users []User `json:"users"`
	}

	err := col.Find(nil).All(&result.Users)
	if err != nil {
		log.Println(err)
	}

	bytes, err := json.MarshalIndent(result.Users, "", "    ")
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(200)
	w.Header().Add("Content-Type", "application/json")
	w.Write(bytes)
}

// CreateUser ...
func CreateUser(w http.ResponseWriter, r *http.Request) {

}

// UpdateUser ...
func UpdateUser(w http.ResponseWriter, r *http.Request) {

}

// DeleteUser ...
func DeleteUser(w http.ResponseWriter, r *http.Request) {

}
