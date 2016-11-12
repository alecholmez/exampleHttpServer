package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/alecholmez/http-server/config"

	mgo "gopkg.in/mgo.v2"
)

// Adapter ...
// Adapter is a wrapper handler
type Adapter func(http.Handler) http.Handler

// Adapt ...
// Wraps each handler with the middleware provided
func Adapt(h http.Handler, adapters ...Adapter) http.Handler {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}

type key int

const (
	mongoSessKey key = iota
	confKey      key = iota
)

// Log ...
// A home-made logger to log the route calls to standard output
func Log(route Route) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			h.ServeHTTP(w, r)

			log.Printf(
				"%s\t%s\t%s\t%s",
				route.Method,
				route.Pattern,
				route.Name,
				time.Since(start),
			)
		})
	}
}

// WithMongo ...
// Copies the mongo session passed to it
// and sticks in the context in the Request.
// This allows for every handler to access the mongo session
func WithMongo(sess *mgo.Session) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// pull a mongo connection from the pool
			sessCopy := sess.Copy()
			defer sessCopy.Close()

			ctx := r.Context()
			ctx = context.WithValue(r.Context(), mongoSessKey, sessCopy)
			r = r.WithContext(ctx)
			h.ServeHTTP(w, r)
		})
	}
}

// WithConf ...
// Copies the configuration object around
// to all middleware for ease of access to global variables
func WithConf(c config.Config) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// Copy the config object to the request context
			ctx := r.Context()
			ctx = context.WithValue(r.Context(), confKey, c)
			r = r.WithContext(ctx)
			h.ServeHTTP(w, r)
		})
	}
}

// GetMongo ...
// Helper function for pulling the mongo session out of the request context
func GetMongo(ctx context.Context) *mgo.Session {
	sess, ok := ctx.Value(mongoSessKey).(*mgo.Session)
	if !ok {
		panic("No session in context")
	}

	return sess
}

// GetConf ...
// Retrieves a config object from the context in the request
func GetConf(ctx context.Context) config.Config {
	conf, ok := ctx.Value(confKey).(config.Config)
	if !ok {
		panic("No configuration in context")
	}

	return conf
}
