package main

import (
	"WikiLinkParser/client"
	"WikiLinkParser/server"
	"WikiLinkParser/worker"
	"flag"
	"fmt"
)

const (
	SERVER_MODE = "server"
	WORKER_MODE = "worker"
	CLIENT_MODE = "client"
)

var (
	mode = flag.String("mode", SERVER_MODE, "server or worker mode")
)

func main() {
	flag.Parse()
	switch *mode {
	case SERVER_MODE:
		server.Run()
	case WORKER_MODE:
		worker.Run()
	case CLIENT_MODE:
		client.Run()
	default:
		fmt.Println("Unsupported mode", *mode)
	}
}
