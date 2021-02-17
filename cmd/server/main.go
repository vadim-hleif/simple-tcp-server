package main

import (
	"flag"
	"fmt"

	"simple-tcp-server/pkg/notifications/endpoint"
	"simple-tcp-server/pkg/notifications/service"
	"simple-tcp-server/pkg/notifications/transport/tcp"
)

func main() {
	storage := service.NewInMemoryUsersStorage()

	server := tcp.NewTcpServer(endpoint.MakeEndpoints(storage))

	var port int
	flag.IntVar(&port, "port", 8080, "server port")
	flag.Parse()

	server.StartServer(fmt.Sprintf(":%d", port))
}
