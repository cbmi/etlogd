package main

import (
	"github.com/cbmi/etlog/encoding"
	"io"
	"io/ioutil"
	"labix.org/v2/mgo/bson"
	"log"
)

func handleFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func handleError(err error) bool {
	if err != nil {
		log.Printf("error: %s\n", err)
		return true
	}
	return false
}

func handleMessage(cfg *Config, r io.Reader) {
	var d bson.M

	b, err := ioutil.ReadAll(r)

	if !handleError(err) {
		err = encoding.UnmarshalJSON(b, &d)

		if !handleError(err) {
			C(cfg).Insert(d)
		}
	}
}
