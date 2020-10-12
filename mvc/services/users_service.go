package services

import (
	"github.com/rwbailey/microservices/mvc/domain"
	"github.com/rwbailey/microservices/mvc/utils"
)

type usersService struct{}

var (
	UsersService usersService
)

func (*usersService) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	user, err := domain.UserDao.GetUser(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}
