package endpoint

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type userStorageMock struct {
	lastSavedUserID     int
	lastSavedFriendsIDs []int

	mockedFriends []int
}

func (u *userStorageMock) SaveUser(userID int, friends []int) {
	u.lastSavedUserID = userID
	u.lastSavedFriendsIDs = friends
}

func (u *userStorageMock) GetAllFriends(userID int) ([]int, error) {
	return u.mockedFriends, nil
}

func Test_UserLoginEndpoint(t *testing.T) {
	mock := &userStorageMock{}
	userLoginEndpoint := makeUserLoginEndpoint(mock)

	mock.mockedFriends = []int{2, 4}
	request := UserLoginRequest{
		UserID:     1,
		FriendsIDs: []int{2, 3, 4},
	}

	response, _ := userLoginEndpoint(nil, request)
	res := response.(UserStatusChangedResponse)

	assert.ElementsMatch(t, res.OnlineFriendsIDs, mock.mockedFriends)
	assert.Equal(t, res.UserID, request.UserID)
	assert.True(t, res.IsOnline)

	assert.ElementsMatch(t, mock.lastSavedFriendsIDs, request.FriendsIDs)
	assert.Equal(t, mock.lastSavedUserID, request.UserID)
}

func Test_UserLogoutEndpoint(t *testing.T) {
	mock := &userStorageMock{}
	userLogoutEndpoint := makeUserLogoutEndpoint(mock)

	mock.mockedFriends = []int{2, 4}
	request := 1

	response, _ := userLogoutEndpoint(nil, request)
	res := response.(UserStatusChangedResponse)

	assert.ElementsMatch(t, res.OnlineFriendsIDs, mock.mockedFriends)
	assert.Equal(t, res.UserID, request)
	assert.False(t, res.IsOnline)
}
