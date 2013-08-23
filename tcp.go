package main

import (
	"log"
	"net"
)

func ServeTcp(cfg *Config) {
	defer cfg.Mongo.session.Close()

	l, err := net.Listen("tcp", cfg.Server.Host)
	handleFatal(err)
	defer l.Close()

	log.Printf("Listening on tcp://%v...\n", l.Addr())

	for {
		conn, err := l.Accept()

		if !handleError(err) {
			log.Println("Connected to client")
			go handleMessage(cfg, conn)
		}
	}
}
