package service

var storage = make(map[int][]int)

// SaveUser add provided userId to all friends for a quick search in the future
func SaveUser(userId int, friends []int) {
	for _, friendId := range friends {
		friendsIds, ok := storage[friendId]

		if ok {
			storage[friendId] = append(friendsIds, userId)
		} else {
			storage[friendId] = []int{userId}
		}
	}
}

func FindAllFriends(userId int) ([]int, bool) {
	friends, ok := storage[userId]

	return friends, ok
}
