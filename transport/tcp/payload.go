package tcp

type Payload struct {
	UserId  int `json:"user_id"`
	Friends []int
}
