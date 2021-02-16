package service

import (
	"fmt"
	"simple-tcp-server/pkg/notifications"
	"sync"
)

func NewInMemoryUsersStorage() notifications.UsersStorage {
	return &inMemoryUsersStorage{}
}

type inMemoryUsersStorage struct {
	storage sync.Map
}

func (s *inMemoryUsersStorage) SaveUser(userId int, friends []int) {
	for _, friendId := range friends {
		friendsIds, ok := s.storage.Load(friendId)

		if ok {
			s.storage.Store(friendId, append(friendsIds.([]int), userId))
		} else {
			s.storage.Store(friendId, []int{userId})
		}
	}
}

func (s *inMemoryUsersStorage) GetAllFriends(userId int) ([]int, error) {
	friends, ok := s.storage.Load(userId)

	if ok {
		return friends.([]int), nil
	}

	return nil, fmt.Errorf("no friends are found")
}
