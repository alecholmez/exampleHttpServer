package main

import (
	"fmt"
	"net/http"

	"github.com/alecholmez/http-server/config"
)

// NewStack ...
func NewStack(r Routes, c config.Config) {

	mongoConnString := fmt.Sprintf("%s:%d", c.Database.Host, c.Database.Port)
	sess := config.NewMongoSession(mongoConnString)

	for _, route := range r {
		var handler http.Handler
		handler = route.HandlerFunc

		http.Handle(route.Pattern, Adapt(http.Handler(handler), Log(route), WithMongo(sess)))
	}
}
