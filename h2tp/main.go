package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alecholmez/http-server/config"
	"github.com/spf13/viper"
)

func main() {
	config.NewConfig(".", "config")
	p := viper.GetInt("server.port")
	port := fmt.Sprintf(":%d", p)
	if port == ":0" {
		port = ":6060"
	}

	NewRouter(routes)

	log.Printf("Server is locally listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
