package endpoint

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type userStorageMock struct {
	lastSavedUserId     int
	lastSavedFriendsIds []int

	mockedFriends []int
}

func (u *userStorageMock) SaveUser(userId int, friends []int) {
	u.lastSavedUserId = userId
	u.lastSavedFriendsIds = friends
}

func (u *userStorageMock) GetAllFriends(userId int) ([]int, error) {
	return u.mockedFriends, nil
}

func Test_UserLoginEndpoint(t *testing.T) {
	mock := &userStorageMock{}
	userLoginEndpoint := makeUserLoginEndpoint(mock)

	mock.mockedFriends = []int{2, 4}
	request := UserLoginRequest{
		UserId:     1,
		FriendsIds: []int{2, 3, 4},
	}

	response, _ := userLoginEndpoint(nil, request)
	res := response.(UserStatusChangedResponse)

	assert.ElementsMatch(t, res.OnlineFriendsIds, mock.mockedFriends)
	assert.Equal(t, res.UserId, request.UserId)
	assert.True(t, res.IsOnline)

	assert.ElementsMatch(t, mock.lastSavedFriendsIds, request.FriendsIds)
	assert.Equal(t, mock.lastSavedUserId, request.UserId)
}
