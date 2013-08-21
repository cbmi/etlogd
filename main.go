package main

import (
	"flag"
	"fmt"
	"os"
)

var defaultServerPorts = map[string]int{
	"http": 4000,
	"tcp":  4001,
	"udp":  4002,
}

func main() {
	serverTypeFlag := flag.String("type", "http", "server type: http, tcp, or udp")
	serverHostFlag := flag.String("host", "", "server host")
	serverPortFlag := flag.Int("port", 0, "server port")

	flag.Parse()

	//
	serverType := *serverTypeFlag
	serverHost := *serverHostFlag
	serverPort := *serverPortFlag

	if serverPort == 0 {
		serverPort = defaultServerPorts[serverType]
	}

	switch serverType {
	case "udp":
		ServeUdp(serverHost, serverPort)
	case "tcp":
		ServeTcp(serverHost, serverPort)
	case "http":
		ServeHttp(serverHost, serverPort)
	default:
		fmt.Println("server type must be http, tcp, or udp")
		os.Exit(1)
	}
}
