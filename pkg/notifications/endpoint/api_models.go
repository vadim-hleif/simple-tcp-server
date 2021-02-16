package endpoint

type UserLoginRequest struct {
	UserId     int
	FriendsIds []int
}

type StatusNotification struct {
	UserId int
	Online bool
}
