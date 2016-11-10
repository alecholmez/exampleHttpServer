package main

import (
	"context"
	"net/http"

	"github.com/deciphernow/commons/middleware"
	mgo "gopkg.in/mgo.v2"
)

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

type key int

const mongoSessKey key = iota

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

// WithMongo ...
func WithMongo(sess *mgo.Session) middleware.Middleware {
	return middleware.MiddlewareFunc(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqSess := sess.Copy()
			defer reqSess.Close()

			ctx := r.Context()
			ctx = context.WithValue(r.Context(), mongoSessKey, reqSess)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	})
}
