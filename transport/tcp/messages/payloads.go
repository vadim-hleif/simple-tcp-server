package messages

type Payload struct {
	UserId  int `json:"user_id"`
	Friends []int
}

type UserOnlineNotification struct {
	UserId int  `json:"friend_id"`
	Online bool `json:"online"`
}
