package main

import (
	"fmt"
	"net"
	"net/http"
	"sync"

	mgo "gopkg.in/mgo.v2"

	"github.com/alecholmez/testHttpServer/config"
	"github.com/deciphernow/commons/middleware"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"github.com/rs/xaccess"
	"github.com/rs/xlog"
	"github.com/spf13/viper"
)

func main() {
	// Get configuration file
	err := config.NewConfig(".", "config")
	if err != nil {
		panic(err)
	}
	host := viper.GetString("server.host")
	port := viper.GetInt("server.port")
	mongoConnString := viper.GetString("database.mongoConnString")

	logConf := xlog.Config{
		Level: xlog.LevelInfo,
		Fields: xlog.F{
			"role": "http-server",
			"host": host,
		},
		Output: xlog.NewOutputChannel(xlog.NewConsoleOutput()),
	}
	logger := xlog.New(logConf)

	mongoSess, err := mgo.Dial(mongoConnString)
	if err != nil {
		panic(err)
	}
	defer mongoSess.Close()

	stack := middleware.Chain(
		middleware.MiddlewareFunc(xlog.NewHandler(logConf)),
		middleware.MiddlewareFunc(xaccess.NewHandler()),
		middleware.MiddlewareFunc(cors.Default().Handler),
		WithMongo(mongoSess),
	)

	router := httprouter.New()

	router.GET("/", Index)
	router.GET("/users", ListUsers)
	router.GET("/users/:id", GetUser)
	router.POST("/users", CreateUser)

	var wg sync.WaitGroup

	wg.Add(1)
	addr := fmt.Sprintf("%s:%d", host, port)
	logger.Infof("team-service listening on %s (HTTP)", addr)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Fatal(err)
	}

	go func() {
		http.Serve(ln, stack.Wrap(router))
		wg.Done()
	}()
}
