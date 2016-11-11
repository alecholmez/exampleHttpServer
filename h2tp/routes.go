package main

import (
	"net/http"

	"github.com/alecholmez/http-server/server"
)

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

var routes = Routes{
	Route{
		"Docs",
		"GET",
		"/",
		server.ServeDocs,
	},
}
