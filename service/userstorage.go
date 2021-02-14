package service

import (
	"fmt"
	"sync"
)

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

// GetAllFriends returns ids of all friends, who have provided userId in friends[]int
func GetAllFriends(userId int) ([]int, error) {
	friends, ok := storage.Load(userId)

	if ok {
		return friends.([]int), nil
	}

	return nil, fmt.Errorf("no friends are found")
}
