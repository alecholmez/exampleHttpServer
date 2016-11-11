package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alecholmez/testHttpServer/config"
	"github.com/spf13/viper"
)

func main() {
	config.NewConfig(".", "config")
	p := viper.GetInt("server.port")
	port := fmt.Sprintf(":%d", p)
	if port == "" {
		port = ":6060"
	}

	NewRouter(routes)

	log.Printf("Server is locally listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
