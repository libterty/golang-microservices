package domains

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := UserDao.GetUser(0)
	assert.Nil(t, user, "We are not expect a user with id")
	assert.NotNil(t, err)
	assert.EqualValues(t, "User was not found: 0", err.Message)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)
}

func TestGetUser(t *testing.T) {
	user, err := UserDao.GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 1, user.Id)
	assert.EqualValues(t, "lib", user.FirstName)
	assert.EqualValues(t, "11", user.LastName)
	assert.EqualValues(t, "lib@test.com", user.Email)
}
