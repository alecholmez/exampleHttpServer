package main

import "net/http"

// Route ...
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes ...
// A wrapper for holding many routes
type Routes []Route

// RTS ...
var routes = Routes{
	Route{"Docs", "GET", "/", ServeDocs},
	Route{"Users", "GET", "/users", ListUsers},
}
