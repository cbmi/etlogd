package main

import (
	"fmt"
	"net"
)

func ServeUdp(host string, port int) {
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", host, port))
	handleError(err)

	conn, err := net.ListenUDP("udp", addr)
	handleError(err)

	// Close on exit
	defer conn.Close()

	fmt.Printf("Listening on udp://%v...", addr)

	for {
		// Handle connection in a goroutine
		go handleMessage(conn)
	}
}
