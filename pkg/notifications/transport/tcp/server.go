package tcp

import (
	"bufio"
	"encoding/json"
	"log"
	"net"
	"sync"

	"simple-tcp-server/pkg/notifications/endpoint"
)

type Server interface {
	StartServer(port string)
}

type tcpServer struct {
	endpoints           endpoint.Endpoints
	connectionsByUserID sync.Map
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
	var payload Payload

	for scanner.Scan() {
		bytes := scanner.Bytes()

		err := json.Unmarshal(bytes, &payload)
		if err != nil {
			log.Println(err)
			break
		}

		server.endpoints.UserLoginEndpoint(nil, endpoint.UserLoginRequest{
			UserID:     payload.UserID,
			FriendsIDs: payload.FriendsIDs,
		})
		// save user's connection
		server.connectionsByUserID.Store(payload.UserID, conn)
	}

	log.Println(conn.RemoteAddr(), "will be closed")

	server.endpoints.UserLogoutEndpoint(nil, payload.UserID)

	// remove user's connection
	server.connectionsByUserID.Delete(payload.UserID)
	_ = conn.Close()
}

// send notification to each friend about a new status of user
// uses internal state to detect connection by user_id
func (server *tcpServer) sendNotifications(userID int, onlineFriendsIDs []int, isUserOnline bool) {
	bytes, _ := json.Marshal(UserStatusNotification{
		UserID: userID,
		Online: isUserOnline,
	})

	for _, friendID := range onlineFriendsIDs {
		connection, ok := server.connectionsByUserID.Load(friendID)

		if ok {
			connection.(net.Conn).Write(bytes)
		}
	}
}
