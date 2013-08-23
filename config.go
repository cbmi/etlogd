package main

import (
	"code.google.com/p/gcfg"
	"flag"
	"fmt"
	"labix.org/v2/mgo"
	"strings"
)

type general struct {
	Verbose bool
}

type server struct {
	Type string
	Host string
}

type mongo struct {
	Host       string
	Database   string
	Collection string
	session    *mgo.Session
}

type Config struct {
	General general
	Server  server
	Mongo   mongo
}

var (
	defaultServerPorts = map[string]int{
		"http": 4000,
		"tcp":  4001,
		"udp":  4002,
	}

	verboseOutput = flag.Bool("verbose", false, "verbose output")
	configPath    = flag.String("config", "", "path to ini-style config file")

	serverType = flag.String("type", "", "server type [http (default), tcp, udp]")
	serverHost = flag.String("host", "", "server host")

	mongoHost       = flag.String("mongo", "", "mongo host")
	mongoDatabase   = flag.String("database", "", "mongo database name")
	mongoCollection = flag.String("collection", "", "mongo collection name")
)

// Performs validation and sets additional values
func validate(cfg *Config) {
	s, err := mgo.Dial(cfg.Mongo.Host)
	handleFatal(err)
	cfg.Mongo.session = s
}

// Initializes new config with the default values
func newConfig() *Config {
	return &Config{
		general{
			Verbose: false,
		},
		server{
			Type: "http",
			Host: "0.0.0.0",
		},
		mongo{
			Host:       "0.0.0.0:27017",
			Database:   "etlog",
			Collection: "logs",
		},
	}
}

func NewConfig() *Config {
	cfg := newConfig()
	validate(cfg)
	return cfg
}

// Initializes new config and merges parsed command-line arguments
func NewConfigWithArgs() *Config {
	cfg := newConfig()
	flag.Parse()

	// Update config, with the supplied config file
	if *configPath != "" {
		err := gcfg.ReadFileInto(cfg, *configPath)
		handleFatal(err)
	}

	// Override default and config with explicit flags
	if *verboseOutput {
		cfg.General.Verbose = *verboseOutput
	}

	if *serverType != "" {
		cfg.Server.Type = *serverType
	}

	if *serverHost != "" {
		cfg.Server.Host = *serverHost
	}

	if *mongoHost != "" {
		cfg.Mongo.Host = *mongoHost
	}

	if *mongoDatabase != "" {
		cfg.Mongo.Database = *mongoDatabase
	}

	if *mongoCollection != "" {
		cfg.Mongo.Collection = *mongoCollection
	}

	// Set port if not defined in the host string
	if !strings.Contains(cfg.Server.Host, ":") {
		serverPort := defaultServerPorts[cfg.Server.Type]
		cfg.Server.Host = fmt.Sprintf("%s:%d", cfg.Server.Host, serverPort)
	}

	validate(cfg)
	return cfg
}
