package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"simple-tcp-server/pkg/notifications/service"
)

type Endpoints struct {
	UserLoginEndpoint  endpoint.Endpoint
	UserLogoutEndpoint endpoint.Endpoint
}

func MakeEndpoints(storage service.UsersStorage) Endpoints {
	return Endpoints{
		UserLoginEndpoint:  makeUserLoginEndpoint(storage),
		UserLogoutEndpoint: makeUserLogoutEndpoint(storage),
	}
}

func makeUserLoginEndpoint(storage service.UsersStorage) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UserLoginRequest)
		storage.SaveUser(req.UserId, req.FriendsIds)

		friends, err := storage.GetAllFriends(req.UserId)
		// just ignore when no one friend are online
		if err != nil {
			return UserLoginResponse{}, err
		}

		return UserLoginResponse{OnlineFriendsIds: friends}, nil
	}
}

func makeUserLogoutEndpoint(storage service.UsersStorage) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		userId := request.(int)
		friends, err := storage.GetAllFriends(userId)

		// just ignore when no one friend are online
		if err != nil {
			return UserLogoutResponse{}, err
		}

		return UserLogoutResponse{OnlineFriendsIds: friends}, nil
	}
}
