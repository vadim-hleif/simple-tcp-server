package service

import (
	"fmt"
	"sync"
)

type UsersStorage interface {
	// SaveUser add provided userId to all friends for a quick search in the future
	SaveUser(userId int, friends []int)

	// GetAllFriends returns ids of all friends, who have provided userId in friends[]int
	GetAllFriends(userId int) ([]int, error)
}

func NewInMemoryUsersStorage() UsersStorage {
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
