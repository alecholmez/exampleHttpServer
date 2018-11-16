package main

import "net/http"

// Route ...
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is a wrapper for holding many routes
type Routes []Route

// CreateRoutes removes the need for a global routes array
func CreateRoutes() Routes {
	return []Route{
		Route{"Docs", "GET", "/", ServeDocs},
		Route{"List All Users", "GET", "/users", ListUsers},
		Route{"Create User", "POST", "/users", CreateUser},
		Route{"Delete User", "DELETE", "/users/{id}", DeleteUser},
		Route{"Update User", "PUT", "/users/{id}", UpdateUser},
		Route{"Get User", "GET", "/users/{id}", GetUser},
	}
}
