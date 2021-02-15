package tcp

import (
	"bufio"
	"encoding/json"
	"log"
	"net"
	"sync"

	"simple-tcp-server/endpoints"
	"simple-tcp-server/transport/tcp/messages"
)

type Server interface {
	StartServer(port string)
}

type tcpServer struct {
	notificationsApi    endpoints.UsersNotificationsApi
	connectionsByUserId sync.Map
}

func NewTcpServer(notificationsApi endpoints.UsersNotificationsApi) Server {
	return &tcpServer{
		notificationsApi: notificationsApi,
	}
}

func (server *tcpServer) StartServer(port string) {
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

		go server.handle(conn)
	}
}

func (server *tcpServer) handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	var payload messages.Payload

	for scanner.Scan() {
		bytes := scanner.Bytes()

		err := json.Unmarshal(bytes, &payload)
		if err != nil {
			log.Println(err)
			break
		}

		server.notificationsApi.UserLoggedIn(endpoints.UserLoginRequest{
			UserId:     payload.UserId,
			FriendsIds: payload.Friends,
		}, server.sendNotifications)

		// save connection
		server.connectionsByUserId.Store(payload.UserId, conn)
	}

	log.Println(conn.RemoteAddr(), "will be closed")
	server.notificationsApi.UserLogOut(payload.UserId, server.sendNotifications)
	_ = conn.Close()
}

// send notifications by user id
// uses internal state to detect connection by user_id
func (server *tcpServer) sendNotifications(messagesByFried map[int]endpoints.StatusNotification) {
	for id, notification := range messagesByFried {
		connection, ok := server.connectionsByUserId.Load(id)

		if ok {
			bytes, _ := json.Marshal(messages.UserStatusNotification{
				UserId: notification.UserId,
				Online: notification.Online,
			})

			connection.(net.Conn).Write(bytes)
		}
	}
}
