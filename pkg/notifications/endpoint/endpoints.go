package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"simple-tcp-server/pkg/notifications"
)

type Endpoints struct {
	UserLoginEndpoint  endpoint.Endpoint
	UserLogoutEndpoint endpoint.Endpoint
}

func MakeEndpoints(storage notifications.UsersStorage) Endpoints {
	return Endpoints{
		UserLoginEndpoint:  makeUserLoginEndpoint(storage),
		UserLogoutEndpoint: makeUserLogoutEndpoint(storage),
	}
}

func makeUserLoginEndpoint(storage notifications.UsersStorage) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UserLoginRequest)
		storage.SaveUser(req.UserID, req.FriendsIDs)

		friends, err := storage.GetAllFriends(req.UserID)
		// just ignore when no one friend are online
		if err != nil {
			return UserStatusChangedResponse{UserID: req.UserID, IsOnline: true}, err
		}

		return UserStatusChangedResponse{UserID: req.UserID, IsOnline: true, OnlineFriendsIDs: friends}, nil
	}
}

func makeUserLogoutEndpoint(storage notifications.UsersStorage) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		userID := request.(int)
		friends, err := storage.GetAllFriends(userID)

		// just ignore when no one friend are online
		if err != nil {
			return UserStatusChangedResponse{UserID: userID, IsOnline: false}, err
		}

		return UserStatusChangedResponse{UserID: userID, IsOnline: false, OnlineFriendsIDs: friends}, nil
	}
}
