package services

import (
	"../domains"
	"../utils"
)

type userService struct {}

var (
	UserService userService
)

func (u *userService) GetUser(userId int64) (*domains.User, *utils.ApplicationError) {
	return domains.UserDao.GetUser(userId)
}
