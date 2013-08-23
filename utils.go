package main

import (
	"encoding/json"
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

// This decodes it as json into an intermediate map, encodes it into BSON,
// the decodes it back into the store. This take advantages of the `inline`
// tag to support arbirtary fields. This can be used as an alternate to
// `json.Unmarshal`
// See: http://godoc.org/labix.org/v2/mgo/bson#Marshal
func unmarshalJSON(b []byte, v interface{}) error {
	var j map[string]interface{}
	json.Unmarshal(b, &j)
	b, _ = bson.Marshal(&j)
	return bson.Unmarshal(b, v)
}

func handleMessage(cfg *Config, r io.Reader) {
	var d bson.M

	b, err := ioutil.ReadAll(r)

	if !handleError(err) {
		err = unmarshalJSON(b, &d)

		if !handleError(err) {
			C(cfg).Insert(d)
		}
	}
}
