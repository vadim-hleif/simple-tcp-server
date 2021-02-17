package messages

type Payload struct {
	UserID     int   `json:"user_id"`
	FriendsIDs []int `json:"friends"`
}

type UserStatusNotification struct {
	UserID int  `json:"friend_id"`
	Online bool `json:"online"`
}
