package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alecholmez/http-server/config"
)

func main() {
	// Load in the config file and define the environment
	c := config.NewConfig("../config.toml")

	Setup()
	Start("0.0.0.0", 9000)

	go func() {
		// Allow time for zipkin to start up but don't block the process
		log.Println("sleeping for 10 seconds...")
		time.Sleep(10 * time.Second)

		log.Println("Starting zipkin registration")
		if err := RegisterZipkin("zipkin", 9411); err != nil {
			log.Fatal(err)
		}
		log.Println("Finished zipkin registration")
	}()

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
