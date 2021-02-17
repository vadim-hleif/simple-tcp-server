package notifications

type UsersStorage interface {
	// SaveUser add provided userID to all friends for a quick search in the future
	SaveUser(userID int, friends []int)

	// GetAllFriends returns IDs of all friends, who have provided userID in friends[]int
	GetAllFriends(userID int) ([]int, error)
}
