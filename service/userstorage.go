package service

import "sync"

var storage = new(sync.Map)

// SaveUser add provided userId to all friends for a quick search in the future
func SaveUser(userId int, friends []int) {
	for _, friendId := range friends {
		friendsIds, ok := storage.Load(friendId)

		if ok {
			storage.Store(friendId, append(friendsIds.([]int), userId))
		} else {
			storage.Store(friendId, []int{userId})
		}
	}
}

func FindAllFriends(userId int) ([]int, bool) {
	friends, ok := storage.Load(userId)

	return friends.([]int), ok
}
