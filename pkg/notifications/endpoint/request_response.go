package endpoint

type UserLoginRequest struct {
	UserID     int
	FriendsIDs []int
}

type UserStatusChangedResponse struct {
	UserID           int
	IsOnline         bool
	OnlineFriendsIDs []int
}
