package main

import (
	"log"
	"net"
)

func ServeUdp(cfg *Config) {
	defer cfg.Mongo.session.Close()

	addr, err := net.ResolveUDPAddr("udp", cfg.Server.Host)
	handleFatal(err)

	conn, err := net.ListenUDP("udp", addr)
	handleFatal(err)
	defer conn.Close()

	log.Printf("Listening on udp://%v...\n", addr)

	for {
		go handleMessage(cfg, conn)
	}
}
