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

func Test_makeUserLoginEndpoint(t *testing.T) {
	mock := &userStorageMock{}
	userLoginEndpoint := makeUserLoginEndpoint(mock)

	mock.mockedFriends = []int{2, 4}
	response, _ := userLoginEndpoint(nil, UserLoginRequest{
		UserId:     1,
		FriendsIds: []int{2, 3, 4},
	})
	res := response.(UserStatusChangedResponse)

	assert.ElementsMatch(t, mock.mockedFriends, res.OnlineFriendsIds)
	assert.Equal(t, mock.lastSavedUserId, 1)
	assert.ElementsMatch(t, mock.lastSavedFriendsIds, []int{2, 3, 4})

}
