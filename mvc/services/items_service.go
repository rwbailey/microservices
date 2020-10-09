package services

import (
	"net/http"

	"github.com/rwbailey/microservices/domain"
	"github.com/rwbailey/microservices/utils"
)

type itemsService struct{}

var (
	ItemsService itemsService
)

func (*itemsService) GetItem(itemId int64) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "implement me",
		StatusCode: http.StatusInternalServerError,
		Code:       "internal server error",
	}
}
