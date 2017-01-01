package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alecholmez/http-server/config"
	"github.com/gorilla/mux"

	"gopkg.in/mgo.v2/bson"
)

// User ..
type User struct {
	Name      string    `json:"name" bson:"name"`
	Email     string    `json:"email" bson:"email"`
	Age       int       `json:"age" bson:"age"`
	ID        string    `json:"_id" bson:"_id"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
}

// ServeDocs ...
func ServeDocs(w http.ResponseWriter, r *http.Request) {
	// Extract the config object from the context in the request
	conf := GetConf(r.Context())

	// Serve the file using the path in the configuration object
	http.ServeFile(w, r, conf.Docs.URL)
}

// ListUsers ...
func ListUsers(w http.ResponseWriter, r *http.Request) {
	// Setup metrics
	ctx := r.Context()
	mon := GetScope(ctx)
	defer mon.Task()(&ctx)(nil)

	// Extract the mongo session from the context in the request
	sess := GetMongo(r.Context())
	col := sess.DB("users").C("users")

	// create a wrapper struct to hold the results
	var result struct {
		Users []bson.M `json:"users"`
	}

	// Retrieve all the user objects in the mongo collection
	err := col.Find(nil).All(&result.Users)
	if err != nil {
		log.Println(err)
	}

	// Write the json to the response
	_ = config.WriteResponse(w, result.Users)
}

// GetUser ...
func GetUser(w http.ResponseWriter, r *http.Request) {
	// Setup metrics
	ctx := r.Context()
	mon := GetScope(ctx)
	defer mon.Task()(&ctx)(nil)

	// Extract the mongo session from the context in the request
	sess := GetMongo(r.Context())
	col := sess.DB("users").C("users")

	// Get the "id" variable from the URL
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		log.Println("ID was empty")
		return
	}

	// Create a wrapper to hold the user object
	// that is to be querried from the mongo collection
	var result struct {
		User bson.M `json:"user"`
	}

	// Find and retrieve the user object in the mongo collection
	err := col.Find(bson.M{"_id": id}).One(&result.User)
	if err != nil {
		log.Println(err)
	}

	// Write the create object to the response
	// so the user can see what they created
	_ = config.WriteResponse(w, result.User)
}

// CreateUser ...
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Setup metrics
	ctx := r.Context()
	mon := GetScope(ctx)
	defer mon.Task()(&ctx)(nil)

	// Extract the mongo session from the context in the request
	sess := GetMongo(r.Context())
	col := sess.DB("users").C("users")

	var user User
	// Read in request
	if ok := config.ReadRequest(w, r, &user); !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Populate necessary data
	user.ID = config.GenID()
	user.CreatedAt = time.Now()

	// Insert the object into mongo
	err := col.Insert(user)
	if err != nil {
		log.Println(err)
	}

	// Write the create object to the response
	// so the user can see what they created
	_ = config.WriteResponse(w, user)
}

// UpdateUser ...
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Setup metrics
	ctx := r.Context()
	mon := GetScope(ctx)
	defer mon.Task()(&ctx)(nil)

	// Extract the mongo session from the context in the request
	sess := GetMongo(r.Context())
	col := sess.DB("users").C("users")

	// Get the "id" variable from the URL
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		log.Println("ID was empty")
		return
	}

	// Get the updated data from the incoming request
	var user User
	if ok := config.ReadRequest(w, r, &user); !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Update the user object with the proper id
	// and the new data
	selected := bson.M{"_id": id}
	updated := bson.M{"$set": user}
	err := col.Update(selected, updated)
	if err != nil {
		log.Println(err)
	}

	// Write the json to the response
	_ = config.WriteResponse(w, user)
}

// DeleteUser ...
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Setup metrics
	ctx := r.Context()
	mon := GetScope(ctx)
	defer mon.Task()(&ctx)(nil)

	// Extract the mongo session from the context in the request
	sess := GetMongo(r.Context())
	col := sess.DB("users").C("users")

	// Get the "id" variable from the URL
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		log.Println("ID was empty")
		return
	}

	// Remove the user object in the mongo collection
	// with the appropriate id
	err := col.Remove(bson.M{"_id": id})
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(http.StatusOK)
}
