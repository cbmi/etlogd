package main

import (
	"github.com/bruth/assert"
	"labix.org/v2/mgo/bson"
	"testing"
)

var jsonString = `
{
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
`

func TestInsertDoc(t *testing.T) {
	cfg := NewConfigWithArgs()
	defer cfg.Mongo.session.Close()

	var d, r bson.M

	// Drop the test database by default
	db := DB(cfg)
	db.DropDatabase()

	col := C(cfg)

	// Unmarshal JSON string into document
	unmarshalJSON([]byte(jsonString), &d)

	// Insert then fetch one
	col.Insert(&d)
	col.Find(nil).One(&r)

	// Ensure the _id key has been created
	v, ok := r["_id"]
	assert.True(t, ok)

	// This looks odd due to the type assertion. This converts it into
	// the interface{} value of "source" to a bson.M (another map) so "uri"
	// can be accessed
	v, _ = r["source"].(bson.M)["uri"]
	assert.Equal(t, v, "148.29.12.100/path/to/users.csv")
}
