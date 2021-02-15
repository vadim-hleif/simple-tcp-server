package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveUser_should_not_lose_data_about_previous_users(t *testing.T) {
	service := NewInMemoryUsersStorage()
	service.SaveUser(1, []int{2, 3, 4})
	service.SaveUser(2, []int{1, 3, 4})
	service.SaveUser(3, []int{2, 4})

	friends, err := service.GetAllFriends(4)

	assert.Nil(t, err)
	assert.ElementsMatch(t, []int{1, 2, 3}, friends)

	friends, err = service.GetAllFriends(2)

	assert.Nil(t, err)
	assert.ElementsMatch(t, []int{1, 3}, friends)

	friends, err = service.GetAllFriends(3)
	assert.Nil(t, err)
	assert.ElementsMatch(t, []int{1, 2}, friends)
}
