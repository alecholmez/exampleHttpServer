package config

import (
	"log"

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

	if _, err := toml.DecodeFile(dir, &conf); err != nil {
		log.Fatal(err)
	}

	return conf
}
