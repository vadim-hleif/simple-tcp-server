package main

import (
	"flag"
	"fmt"
	
	"simple-tcp-server/endpoints"
	"simple-tcp-server/service"
	"simple-tcp-server/transport/tcp"
)

func main() {
	storage := service.NewInMemoryUsersStorage()

	server := tcp.NewTcpServer(endpoints.MakeEndpoints(storage))

	var port int
	flag.IntVar(&port, "port", 8080, "server port")
	flag.Parse()

	server.StartServer(fmt.Sprintf(":%d", port))
}
