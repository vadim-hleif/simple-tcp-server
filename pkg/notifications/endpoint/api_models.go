package endpoint

type UserLoginRequest struct {
	UserId     int
	FriendsIds []int
}

type UserLoginResponse struct {
	OnlineFriendsIds []int
}

type UserLogoutResponse struct {
	OnlineFriendsIds []int
}
