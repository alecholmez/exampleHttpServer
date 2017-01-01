package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alecholmez/http-server/config"
	"github.com/fvbock/endless"

	mgo "gopkg.in/mgo.v2"
	zipkin "gopkg.in/spacemonkeygo/monkit-zipkin.v2"
	monkit "gopkg.in/spacemonkeygo/monkit.v2"
	"gopkg.in/spacemonkeygo/monkit.v2/environment"
	"gopkg.in/spacemonkeygo/monkit.v2/present"
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
	mongoSessKey  key = iota
	confKey       key = iota
	monkitTaskKey key = iota
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

// Create a environment scope for all the handlers to share
var mon = monkit.Package()

// WithMetrics ...
// Adds metrics functionality to each individual http handler in a golang service
func WithMetrics() Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// Copy the context out of the request
			ctx := r.Context()

			// Add the monkit scope into the request context
			newCtx := context.WithValue(ctx, monkitTaskKey, mon)
			r = r.WithContext(newCtx)
			h.ServeHTTP(w, r)
		})
	}
}

// GetScope ...
func GetScope(ctx context.Context) *monkit.Scope {
	scope, ok := ctx.Value(monkitTaskKey).(*monkit.Scope)
	if !ok {
		panic("no monkit scope in context")
	}
	return scope
}

// Setup registers the service environment into the  monkit default registry ...
func Setup() {
	environment.Register(monkit.Default)
}

// Start takes the host and port for the metrics server to listen on ...
func Start(host string, port int) {
	addr := fmt.Sprintf("%s:%d", host, port)
	fmt.Printf("Service listening at: %s\n", addr)

	// Start the metrics server on a go routine (non-blocking async)
	// Endless allows for 0-downtime updates/delpoyments
	go func() {
		err := endless.ListenAndServe(addr, present.HTTP(monkit.Default))
		if err != nil {
			panic("Couldn't start metrics server")
		}
	}()
}

// RegisterZipkin ...
func RegisterZipkin(host string, port int) error {
	// Check host and port
	if host == "" || port == 0 {
		return errors.New("Host and Port must have a valid value")
	}

	// Define a zipkin collector using the given host and port
	addr := fmt.Sprintf("%s:%d", host, port)
	collector, err := zipkin.NewScribeCollector(addr)
	if err != nil {
		return err
	}

	// Register the zipkin collector with the zipkin instance and monkit registry
	zipkin.RegisterZipkin(monkit.Default, collector, zipkin.Options{
		Fraction: 1,
	})

	return nil
}

// WithZipkin adds the zipkin trace headers into all the functions/handlers with a monkit task ...
func WithZipkin() Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

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
