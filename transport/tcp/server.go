package tcp

import (
	"bufio"
	"encoding/json"
	"log"
	"net"
	"simple-tcp-server/transport/tcp/messages"
	"sync"
)

type MessageHandler interface {
	OnMessage(messages.Payload)
}

var connectionsByUserId = new(sync.Map)

func SendMessageToUser(userId int, messageSupplier func() []byte) {
	c, ok := connectionsByUserId.Load(userId)

	if ok {
		c.(net.Conn).Write(messageSupplier())
	}
}

func StartServer(port string, handler MessageHandler) {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	defer listen.Close()
	log.Println("server started at port:", port)

	for {
		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}
		log.Println("connection accepted:", conn.RemoteAddr())

		go handle(conn, handler)
	}
}

func handle(conn net.Conn, handler MessageHandler) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		bytes := scanner.Bytes()

		var payload messages.Payload
		err := json.Unmarshal(bytes, &payload)
		if err != nil {
			log.Println(err)
			break
		}

		// save connection
		connectionsByUserId.Store(payload.UserId, conn)
		handler.OnMessage(payload)
	}

	log.Println(conn.RemoteAddr(), "will be closed")
	_ = conn.Close()
}
