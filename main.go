package main

import (
	"log"
	"net/http"

	"github.com/alecholmez/testHttpServer/config"
	"github.com/julienschmidt/httprouter"
	"github.com/spf13/viper"
)

func main() {
	// Get configuration file
	err := config.NewConfig(".", "config")
	if err != nil {
		panic(err)
	}
	port := viper.GetString("server.port")

	router := httprouter.New()

	router.GET("/", Index)
	router.GET("/users", ListUsers)
	router.GET("/users/:id", GetUser)
	router.POST("/users", CreateUser)

	log.Fatal(http.ListenAndServe(port, router))
}
