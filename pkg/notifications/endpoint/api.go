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
	UserLoggedIn func(UserLoginRequest) UserLoginResponse
	UserLogOut   func(int) UserLogoutResponse
}

func MakeEndpoints(storage service.UsersStorage) UsersNotificationsApi {
	return UsersNotificationsApi{
		UserLoggedIn: func(request UserLoginRequest) UserLoginResponse {
			storage.SaveUser(request.UserId, request.FriendsIds)

			friends, err := storage.GetAllFriends(request.UserId)
			// just ignore when no one friend are online
			if err != nil {
				return UserLoginResponse{}
			}

			return UserLoginResponse{
				OnlineFriendsIds: friends,
			}
		},
		UserLogOut: func(userId int) UserLogoutResponse {
			friends, err := storage.GetAllFriends(userId)

			// just ignore when no one friend are online
			if err != nil {
				return UserLogoutResponse{}
			}

			return UserLogoutResponse{
				OnlineFriendsIds: friends,
			}
		},
	}
}
