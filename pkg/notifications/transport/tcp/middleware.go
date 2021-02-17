package tcp

import (
	"context"

	kit "github.com/go-kit/kit/endpoint"

	"simple-tcp-server/pkg/notifications/endpoint"
)

func tcpNotificationsMiddleWare(server *tcpServer) kit.Middleware {
	return func(next kit.Endpoint) kit.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			response, err = next(ctx, request)

			res := response.(endpoint.UserStatusChangedResponse)
			server.sendNotifications(res.UserID, res.OnlineFriendsIDs, res.IsOnline)

			return response, err
		}
	}
}
