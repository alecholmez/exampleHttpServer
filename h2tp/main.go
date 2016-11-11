package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alecholmez/http-server/config"
)

func main() {
	c := config.NewConfig("config.toml")

	port := fmt.Sprintf(":%d", c.Server.Port)
	if port == ":0" {
		port = ":6060"
	}

	stack := NewStack(RTS, c)

	log.Printf("Server is locally listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, stack))
}
