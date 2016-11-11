package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	mgo "gopkg.in/mgo.v2"
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

type key int

const (
	mongoSessKey key = iota
)

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

// WithMongo ...
func WithMongo(sess *mgo.Session) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sessCopy := sess.Copy()
			defer sess.Copy()

			ctx := context.WithValue(r.Context(), mongoSessKey, sessCopy)
			r = r.WithContext(ctx)
			h.ServeHTTP(w, r)
		})
	}
}

// GetMongo ...
func GetMongo(ctx context.Context) *mgo.Session {
	sess, ok := ctx.Value(mongoSessKey).(*mgo.Session)
	if !ok {
		panic("No session in context")
	}

	return sess
}
