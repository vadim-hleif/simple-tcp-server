package tcp

import (
	"bufio"
	"encoding/json"
	"log"
	"net"
	"sync"

	"simple-tcp-server/pkg/notifications/endpoint"
	"simple-tcp-server/pkg/notifications/transport/tcp/messages"
)

type Server interface {
	StartServer(port string)
}

type tcpServer struct {
	notificationsApi    endpoint.UsersNotificationsApi
	connectionsByUserId sync.Map
}

func NewTcpServer(notificationsApi endpoint.UsersNotificationsApi) Server {
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

		response := server.notificationsApi.UserLoggedIn(endpoint.UserLoginRequest{
			UserId:     payload.UserId,
			FriendsIds: payload.Friends,
		})
		server.sendNotifications(payload.UserId, response.OnlineFriendsIds, true)

		// save connection
		server.connectionsByUserId.Store(payload.UserId, conn)
	}

	log.Println(conn.RemoteAddr(), "will be closed")
	response := server.notificationsApi.UserLogOut(payload.UserId)
	server.sendNotifications(payload.UserId, response.OnlineFriendsIds, false)
	server.connectionsByUserId.Delete(payload.UserId)

	_ = conn.Close()
}

// send notification to each friend about a new status of user
// uses internal state to detect connection by user_id
func (server *tcpServer) sendNotifications(userId int, onlineFriendsIds []int, isUserOnline bool) {
	for _, friendId := range onlineFriendsIds {
		connection, ok := server.connectionsByUserId.Load(friendId)

		if ok {
			bytes, _ := json.Marshal(messages.UserStatusNotification{
				UserId: userId,
				Online: isUserOnline,
			})

			connection.(net.Conn).Write(bytes)
		}
	}
}
