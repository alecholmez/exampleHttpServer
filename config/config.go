package config

import (
	"log"

	mgo "gopkg.in/mgo.v2"

	"github.com/BurntSushi/toml"
)

// Config ...
type Config struct {
	Database Database `toml:"database"`
	Server   Server   `toml:"server"`
}

// Database ..
type Database struct {
	Port int    `toml:"port"`
	Host string `toml:"host"`
}

// Server ...
type Server struct {
	Port  int  `toml:"port"`
	Oauth bool `toml:"oauth"`
}

// NewConfig ...
func NewConfig(dir string) Config {
	var conf Config

	// Parse the config toml file
	if _, err := toml.DecodeFile(dir, &conf); err != nil {
		log.Fatal(err)
	}

	// If the database config is nil then populate the fields
	// using appropriate localhost values
	if conf.Database.Host == "" && conf.Database.Port == 0 {
		conf.Database.Host = "localhost"
		conf.Database.Port = 27017
	}

	// If the server config is nil then populate the port to the default
	if conf.Server.Port == 0 {
		conf.Server.Port = 6060
	}

	return conf
}

// NewMongoSession ...
// Creates a mongo session with a mongo instance
// to pass around to the application as middleware
func NewMongoSession(conn string) *mgo.Session {
	sess, err := mgo.Dial(conn)
	if err != nil {
		panic(err)
	}

	return sess
}
