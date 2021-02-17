package service

import (
	"fmt"
	"sync"

	"simple-tcp-server/pkg/notifications"
)

func NewInMemoryUsersStorage() notifications.UsersStorage {
	return &inMemoryUsersStorage{
		storage: map[int]map[int]bool{},
	}
}

type inMemoryUsersStorage struct {
	// user_id and his friends (set)
	storage map[int]map[int]bool
	mu      sync.Mutex
}

func (s *inMemoryUsersStorage) SaveUser(userID int, friends []int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, friendID := range friends {
		friendsIDs, ok := s.storage[friendID]

		if ok {
			friendsIDs[userID] = true
		} else {
			s.storage[friendID] = map[int]bool{userID: true}
		}
	}
}

func (s *inMemoryUsersStorage) GetAllFriends(userID int) ([]int, error) {
	s.mu.Lock()
	friends, ok := s.storage[userID]
	s.mu.Unlock()

	if ok {
		return keysAsSlice(friends), nil
	}

	return nil, fmt.Errorf("no friends are found")
}

func keysAsSlice(data map[int]bool) []int {
	keys := make([]int, len(data))

	i := 0
	for key := range data {
		keys[i] = key
		i++
	}

	return keys
}
