package services

import "github.com/rwbailey/microservices/domain"

func GetUser(userId int64) (*domain.User, error) {
	return domain.GetUser(userId)
}
