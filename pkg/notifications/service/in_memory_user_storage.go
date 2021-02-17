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

func (s *inMemoryUsersStorage) SaveUser(userID int, friends []int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, friendID := range friends {
		friendsIDs, ok := s.storage[friendID]

		if ok {
			s.storage[friendID] = append(friendsIDs, userID)
		} else {
			s.storage[friendID] = []int{userID}
		}
	}
}

func (s *inMemoryUsersStorage) GetAllFriends(userID int) ([]int, error) {
	s.mu.Lock()
	friends, ok := s.storage[userID]
	s.mu.Unlock()

	if ok {
		return friends, nil
	}

	return nil, fmt.Errorf("no friends are found")
}
