package tcp

import (
	"bufio"
	"fmt"
	"net"
)

type MessageHandler func([]byte)

func StartServer(port string, handler MessageHandler) {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}

		go handle(conn, handler)
	}
}

func handle(conn net.Conn, handler MessageHandler) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		bytes := scanner.Bytes()
		handler(bytes)
	}

	fmt.Println(conn.RemoteAddr(), "will be closed")
	conn.Close()
}
