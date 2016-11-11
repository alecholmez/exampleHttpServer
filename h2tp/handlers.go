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
	Name      string        `json:"name" bson:"name"`
	Email     string        `json:"email" bson:"email"`
	Age       int           `json:"age" bson:"age"`
	ID        bson.ObjectId `json:"_id" bson:"_id"`
	CreatedAt time.Time     `json:"createdAt" bson:"createdAt"`
}

// ServeDocs ...
func ServeDocs(w http.ResponseWriter, r *http.Request) {
	conf := GetConf(r.Context())

	// Serve the file using the path in the configuration object
	http.ServeFile(w, r, conf.Docs.URL)
}

// ListUsers ...
func ListUsers(w http.ResponseWriter, r *http.Request) {
	sess := GetMongo(r.Context())
	defer sess.Close()
	col := sess.DB("users").C("users")

	var result struct {
		Users []bson.M `json:"users"`
	}

	err := col.Find(nil).All(&result.Users)
	if err != nil {
		log.Println(err)
	}

	// Write the create object to the response
	// so the user can see what they created
	code := config.WriteResponse(w, result.Users)
	w.WriteHeader(code)
}

// GetUser ...
func GetUser(w http.ResponseWriter, r *http.Request) {
	sess := GetMongo(r.Context())
	defer sess.Close()
	col := sess.DB("users").C("users")

	vars := mux.Vars(r)
	id := vars["id"]

	var result struct {
		User bson.M `json:"user"`
	}

	err := col.FindId(bson.ObjectIdHex(id)).All(&result.User)
	if err != nil {
		log.Println(err)
	}

	// Write the create object to the response
	// so the user can see what they created
	code := config.WriteResponse(w, result.User)
	w.WriteHeader(code)
}

// CreateUser ...
func CreateUser(w http.ResponseWriter, r *http.Request) {
	sess := GetMongo(r.Context())
	defer sess.Close()
	col := sess.DB("users").C("users")

	var user User
	// Read in request
	if ok := config.ReadRequest(w, r, &user); !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Populate necessary data
	user.ID = bson.NewObjectId()
	user.CreatedAt = time.Now()

	// Insert the object into mongo
	err := col.Insert(user)
	if err != nil {
		log.Println(err)
	}

	// Write the create object to the response
	// so the user can see what they created
	code := config.WriteResponse(w, user)
	w.WriteHeader(code)
}

// UpdateUser ...
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	sess := GetMongo(r.Context())
	defer sess.Close()
	col := sess.DB("users").C("users")

	vars := mux.Vars(r)
	id := vars["id"]

	var user User
	if ok := config.ReadRequest(w, r, &user); !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := col.UpdateId(bson.ObjectIdHex(id), user)
	if err != nil {
		log.Println(err)
	}

	code := config.WriteResponse(w, user)
	w.WriteHeader(code)
}

// DeleteUser ...
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	sess := GetMongo(r.Context())
	defer sess.Close()
	col := sess.DB("users").C("users")

	vars := mux.Vars(r)
	id := vars["id"]

	err := col.Remove(bson.M{"_id": id})
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(http.StatusOK)
}
