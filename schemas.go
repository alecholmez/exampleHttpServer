package main

// User ...
type User struct {
	Name  string `json:"user"`
	Email string `json:"email"`
	ID    string `json:"_id"`
}

// Users ...
type Users []User
