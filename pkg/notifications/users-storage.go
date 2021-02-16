package notifications

type UsersStorage interface {
	// SaveUser add provided userId to all friends for a quick search in the future
	SaveUser(userId int, friends []int)

	// GetAllFriends returns ids of all friends, who have provided userId in friends[]int
	GetAllFriends(userId int) ([]int, error)
}
