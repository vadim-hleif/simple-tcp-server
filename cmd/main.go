package main

import (
	"encoding/json"
	"simple-tcp-server/service"
	"simple-tcp-server/transport/tcp"
	"simple-tcp-server/transport/tcp/messages"
)

type FriendsNotificationsHandler struct{}

func main() {
	tcp.StartServer(":8080", FriendsNotificationsHandler{})
}

func (t FriendsNotificationsHandler) OnMessage(payload messages.Payload) {
	service.SaveUser(payload.UserId, payload.Friends)

	friends, err := service.GetAllFriends(payload.UserId)

	// just ignore when no friend are provided
	if err != nil {
		return
	}

	for _, friendId := range friends {
		tcp.SendMessageToUser(friendId, func() []byte {
			message := messages.UserOnlineNotification{
				UserId: payload.UserId,
				Online: true,
			}
			bytes, _ := json.Marshal(message)
			return bytes
		})
	}
}
