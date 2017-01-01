package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alecholmez/http-server/config"
)

func main() {
	// Load in the config file and define the environment
	c := config.NewConfig("../config.toml")

	Setup()
	Start("0.0.0.0", 9000)
	RegisterZipkin("zipkin", 9411)

	// Define the port
	port := fmt.Sprintf(":%d", c.Server.Port)
	if port == ":0" {
		port = ":6060"
	}

	// Create the middleware stack
	stack := NewStack(RTS, c)

	// Listen and serve using the middleware stack
	log.Printf("Server is locally listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, stack))
}
