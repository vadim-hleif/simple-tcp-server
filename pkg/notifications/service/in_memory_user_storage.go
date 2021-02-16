package service

import (
	"fmt"
	"sync"

	"simple-tcp-server/pkg/notifications"
)

func NewInMemoryUsersStorage() notifications.UsersStorage {
	return &inMemoryUsersStorage{
		storage: make(map[int][]int),
	}
}

type inMemoryUsersStorage struct {
	storage map[int][]int
	mu      sync.Mutex
}

func (s *inMemoryUsersStorage) SaveUser(userId int, friends []int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, friendId := range friends {
		friendsIds, ok := s.storage[friendId]

		if ok {
			s.storage[friendId] = append(friendsIds, userId)
		} else {
			s.storage[friendId] = []int{userId}
		}
	}
}

func (s *inMemoryUsersStorage) GetAllFriends(userId int) ([]int, error) {
	s.mu.Lock()
	friends, ok := s.storage[userId]
	s.mu.Unlock()

	if ok {
		return friends, nil
	}

	return nil, fmt.Errorf("no friends are found")
}
