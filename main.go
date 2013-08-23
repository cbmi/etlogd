package main

import (
	"flag"
	"os"
)

func main() {
	cfg := NewConfigWithArgs()

	switch cfg.Server.Type {
	case "udp":
		ServeUdp(cfg)
	case "tcp":
		ServeTcp(cfg)
	case "http":
		ServeHttp(cfg)
	default:
		flag.Usage()
		os.Exit(1)
	}
}
