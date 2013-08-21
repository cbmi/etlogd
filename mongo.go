package main

import (
	"labix.org/v2/mgo"
)

// Takes an array of bytes containing the JSON data and inserts it
// into mongo
func insertDoc(d interface{}) {
	mongoHost := "0.0.0.0:27017"

	// Connect to mongo, if there is an issue here, we are in trouble
	s, err := mgo.Dial(mongoHost)
	defer s.Close()
	handleError(err)

	// Insert the data into the collection
	c := s.DB("etlog").C("logs")
	err = c.Insert(d)
	handleError(err)
}
