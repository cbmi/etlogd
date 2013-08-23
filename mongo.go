package main

import (
	"labix.org/v2/mgo"
)

// Convenience function to get a database given a config
func DB(cfg *Config) *mgo.Database {
	return cfg.Mongo.session.DB(cfg.Mongo.Database)
}

// Convenience function to get a collection given a config
func C(cfg *Config) *mgo.Collection {
	return cfg.Mongo.session.DB(cfg.Mongo.Database).C(cfg.Mongo.Collection)
}
