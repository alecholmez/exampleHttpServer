package main

import "net/http"

// Middleware ...
// A Middleware Wraps an existing http.Handler.
type Middleware interface {
	Wrap(http.Handler) http.Handler
}

// MiddlewareFunc ...
// The MiddlewareFunc type is an adapter to allow the use of ordinary functions
// as HTTP middleware. If f is a function with the appropriate signature,
// MiddlewareFunc(f) is a Middleware that calls f.
type MiddlewareFunc func(next http.Handler) http.Handler

// Wrap ...
// Wrap calls f(next).
func (m MiddlewareFunc) Wrap(next http.Handler) http.Handler {
	return m(next)
}

// Chain ...
/*
Compose a set of middleware into a single middleware.
Given the following:
	Chain(f, g, h)
f would be run first, followed by g and h.
*/
func Chain(ms ...Middleware) Middleware {
	return MiddlewareFunc(func(next http.Handler) http.Handler {
		for i := len(ms) - 1; i >= 0; i-- {
			next = ms[i].Wrap(next)
		}
		return next
	})
}
