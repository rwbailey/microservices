package services

import (
	"github.com/rwbailey/microservices/domain"
	"github.com/rwbailey/microservices/utils"
)

type usersService struct{}

var (
	UsersService usersService
)

func (*usersService) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
