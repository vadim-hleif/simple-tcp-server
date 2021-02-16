package endpoint

type UserLoginRequest struct {
	UserId     int
	FriendsIds []int
}

type UserStatusChangedResponse struct {
	UserId           int
	IsOnline         bool
	OnlineFriendsIds []int
}
