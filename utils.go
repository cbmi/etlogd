package main

import (
	"github.com/cbmi/etlog/encoding"
	"io"
	"io/ioutil"
	"log"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func handleMessage(r io.Reader) {
	var t map[string]interface{}
	b, _ := ioutil.ReadAll(r)
	err := encoding.UnmarshalJSON(b, &t)
	handleError(err)
	insertDoc(t)
}
