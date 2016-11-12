package config

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	mgo "gopkg.in/mgo.v2"

	"github.com/BurntSushi/toml"
)

// Config ...
type Config struct {
	Database Database `toml:"database"`
	Server   Server   `toml:"server"`
	Docs     Docs     `toml:"docs"`
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

// Docs ...
type Docs struct {
	URL string `toml:"url"`
}

// NewConfig ...
// Parses the config file using a 3rd party toml parser
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

	if conf.Docs.URL == "" {
		conf.Docs.URL = "docs/index.html"
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

// ReadRequest ...
// Reads the incoming request and decodes it into the given interface
func ReadRequest(w http.ResponseWriter, r *http.Request, value interface{}) (ok bool) {
	err := json.NewDecoder(r.Body).Decode(value)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return false
	}

	return true
}

// WriteResponse ...
// Writes the given interface to the response in json form
func WriteResponse(w http.ResponseWriter, value interface{}) int {
	bytes, err := json.MarshalIndent(value, "", "    ")
	if err != nil {
		return http.StatusInternalServerError
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
	return http.StatusOK
}

// GenID ...
// Creates a random hash using the current system time in nano seconds as the seed
func GenID() string {
	// Use the current timestamp as a seed for the random number genereator
	source := rand.NewSource(time.Now().UnixNano())
	// Create a string as the ID
	s := strconv.Itoa(rand.New(source).Int())

	//SHA1 hash generation
	hasher := sha1.New()
	hasher.Write([]byte(s))
	hash := hex.EncodeToString(hasher.Sum(nil))

	return hash
}
