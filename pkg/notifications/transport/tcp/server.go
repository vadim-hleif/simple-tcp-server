package tcp

import (
	"bufio"
	"context"
	"encoding/json"
	"log"
	"net"
	"sync"

	kit "github.com/go-kit/kit/endpoint"

	"simple-tcp-server/pkg/notifications/endpoint"
	"simple-tcp-server/pkg/notifications/transport/tcp/messages"
)

type Server interface {
	StartServer(port string)
}

type tcpServer struct {
	endpoints           endpoint.Endpoints
	connectionsByUserId sync.Map
}

func NewTcpServer(endpoints endpoint.Endpoints) Server {
	var server = &tcpServer{}
	middleware := tcpNotificationsMiddleWare(server)

	server.endpoints = endpoint.Endpoints{
		UserLoginEndpoint:  middleware(endpoints.UserLoginEndpoint),
		UserLogoutEndpoint: middleware(endpoints.UserLogoutEndpoint),
	}

	return server
}

func tcpNotificationsMiddleWare(server *tcpServer) kit.Middleware {
	return func(next kit.Endpoint) kit.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			response, err = next(ctx, request)

			res := response.(endpoint.UserStatusChangedResponse)
			server.sendNotifications(res.UserId, res.OnlineFriendsIds, res.IsOnline)

			return response, err
		}
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

		server.endpoints.UserLoginEndpoint(nil, endpoint.UserLoginRequest{
			UserId:     payload.UserId,
			FriendsIds: payload.Friends,
		})
		// save user's connection
		server.connectionsByUserId.Store(payload.UserId, conn)
	}

	log.Println(conn.RemoteAddr(), "will be closed")

	server.endpoints.UserLogoutEndpoint(nil, payload.UserId)

	// remove user's connection
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
