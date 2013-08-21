package main

import (
	"github.com/cbmi/etlog/encoding"
	"github.com/cojac/assert"
	"labix.org/v2/mgo"
	"testing"
)

type Message struct {
	Extra map[string]interface{} `json:"-" bson:",inline"`
}

func TestInsertDoc(t *testing.T) {
	d := Message{}

	mongoHost := "0.0.0.0:27017"

	s, _ := mgo.Dial(mongoHost)
	defer s.Close()

	// Insert the data into the collection
	c := s.DB("etlog").C("logs")
    c.RemoveId(1)

	encoding.UnmarshalJSON([]byte(`
        {
            "_id": 1,
            "timestamp": "2013-08-13T05:43:03.32344",
            "action": "update",
            "script": {
                "uri": "https://github.com/cbmi/project/blob/master/parse-users.py",
                "version": "a32f87cb"
            },
            "source": {
                "type": "delimited",
                "delimiter": ",",
                "uri": "148.29.12.100/path/to/users.csv",
                "name": "users.csv",
                "line": 5,
                "column": 4
            },
            "target": {
                "type": "relational",
                "uri": "148.29.12.101:5236",
                "database": "socialapp",
                "table": "users",
                "row": { "id": 38 },
                "column": "email"
            }
        }
    `), &d)

	insertDoc(&d)

	// Retrieve message
	var r []Message
	c.FindId(1).All(&r)

	assert.Equal(t, 1, len(r))

	_, ok := r[0].Extra["source"]
	assert.True(t, ok)
}
