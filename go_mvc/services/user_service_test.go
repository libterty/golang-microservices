package services

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"

	"../domains"
	"../utils"
)

var (
	userDaoMock usersDaoMock

	getUserFunction func(userId int64) (*domains.User, *utils.ApplicationError)
)

func init() {
	domains.UserDao = &usersDaoMock{}
}

type usersDaoMock struct {}

func (m *usersDaoMock) GetUser(userId int64) (*domains.User, *utils.ApplicationError) {
	return getUserFunction(userId )
}

func TestUserNotFound(t *testing.T) {
	getUserFunction = func(userId int64) (*domains.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			StatusCode: http.StatusNotFound,
			Message: "User was not found: 0",
			Code: "not_found",
		}
	}
	user, err := UserService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, "User was not found: 0", err.Message)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)
}

func TestUserFound(t *testing.T)  {
	getUserFunction = func(userId int64) (*domains.User, *utils.ApplicationError) {
		return &domains.User{
			Id: 1,
		}, nil
	}
	user, err := UserService.GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, int64(1), user.Id)

}