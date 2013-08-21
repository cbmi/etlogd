package main

import (
	"fmt"
	"net"
)

func tcpHandle(l net.Listener) {
	// Wait for a connection
	conn, _ := l.Accept()
	// Close on exit
	defer conn.Close()

	fmt.Println("Connected to client...")

	// Handle connection in a goroutine
	handleMessage(conn)
}

func ServeTcp(host string, port int) {
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	handleError(err)

	// Close on exit
	defer l.Close()

	fmt.Printf("Listening on tcp://%v...\n", l.Addr())

	for {
		go tcpHandle(l)
	}
}
