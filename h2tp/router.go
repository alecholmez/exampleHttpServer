package main

import (
	"fmt"
	"net/http"
	"time"
)

// Adapter ...
type Adapter func(http.Handler) http.Handler

// Adapt ...
func Adapt(h http.Handler, adapters ...Adapter) http.Handler {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}

// Log ...
func Log(route Route) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			const layout = "Jan 2, 2006 at 3:04:05 PM (MST)"

			time := time.Now().Format(layout)
			fmt.Printf("%s: %s  -  %s\n", time, route.Method, route.Pattern)

			h.ServeHTTP(w, r)
		})
	}
}

// NewRouter ...
func NewRouter(r Routes) {

	for _, route := range r {
		var handler http.Handler
		handler = route.HandlerFunc

		http.Handle(route.Pattern, Adapt(http.Handler(handler), Log(route)))
	}
}
