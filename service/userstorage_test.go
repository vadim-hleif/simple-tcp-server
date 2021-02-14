package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSaveUser_should_not_lose_data_about_previous_users(t *testing.T) {
	SaveUser(1, []int{2, 3, 4})
	SaveUser(2, []int{1, 3, 4})
	SaveUser(3, []int{2, 4})

	friends, err := GetAllFriends(4)

	assert.Nil(t, err)
	assert.ElementsMatch(t, []int{1, 2, 3}, friends)

	friends, err = GetAllFriends(2)

	assert.Nil(t, err)
	assert.ElementsMatch(t, []int{1, 3}, friends)

	friends, err = GetAllFriends(3)
	assert.Nil(t, err)
	assert.ElementsMatch(t, []int{1, 2}, friends)
}
