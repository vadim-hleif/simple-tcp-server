package endpoint

import (
	"simple-tcp-server/pkg/notifications/service"
)

// callbacks accept map with users ids and messages for them
// e.g.
//1 -> {"online": false}
//2 -> {"online": false}
//3 -> any text
type UsersNotificationsApi struct {
	UserLoggedIn func(UserLoginRequest, func(map[int]StatusNotification))
	UserLogOut   func(int, func(map[int]StatusNotification))
}

func MakeEndpoints(storage service.UsersStorage) UsersNotificationsApi {
	return UsersNotificationsApi{
		UserLoggedIn: func(request UserLoginRequest, sendNotifications func(map[int]StatusNotification)) {
			storage.SaveUser(request.UserId, request.FriendsIds)

			friends, err := storage.GetAllFriends(request.UserId)
			// just ignore when no one friend are online
			if err != nil {
				return
			}

			notifications := toStatusNotifications(request.UserId, true, friends)
			sendNotifications(notifications)
		},
		UserLogOut: func(userId int, sendNotifications func(map[int]StatusNotification)) {
			friends, err := storage.GetAllFriends(userId)

			// just ignore when no one friend are online
			if err != nil {
				return
			}

			notifications := toStatusNotifications(userId, false, friends)
			sendNotifications(notifications)
		},
	}
}

func toStatusNotifications(userId int, isOnline bool, friends []int) map[int]StatusNotification {
	notifications := make(map[int]StatusNotification)

	for _, id := range friends {
		notifications[id] = StatusNotification{
			UserId: userId,
			Online: isOnline,
		}
	}

	return notifications
}
