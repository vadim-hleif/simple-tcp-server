package main

import (
	"encoding/json"
	"fmt"
	"simple-tcp-server/service"
	"simple-tcp-server/transport/tcp"
)

func main() {
	tcp.StartServer(":8080", func(bytes []byte) {
		var payload tcp.Payload

		err := json.Unmarshal(bytes, &payload)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(payload)
			service.SaveUser(payload.UserId, payload.Friends)
			friends, ok := service.FindAllFriends(payload.UserId)

			if ok {
				fmt.Println(friends)
				//todo notify all of them, need to persist theirs connections
			}
		}
	})
}
