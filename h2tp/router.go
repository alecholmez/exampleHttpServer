package main

import (
	"fmt"
	"net/http"

	"github.com/alecholmez/http-server/config"
	"github.com/gorilla/mux"
)

// NewStack ...
// A stack to wrap each http handler with the middleware provided
func NewStack(routes Routes, c config.Config) *mux.Router {

	// Create the mongo session to pass to the middleware
	mongoConnString := fmt.Sprintf("%s:%d", c.Database.Host, c.Database.Port)
	sess := config.NewMongoSession(mongoConnString)

	// Create the router
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc

		// A router is created with each handler defined in the routes array
		// which will be initiated with a logger and a mongo session
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(
				Adapt(http.Handler(handler),
					Log(route),
					WithMongo(sess),
					WithConf(c),
					WithMetrics(),
				),
			)
	}

	return router
}
