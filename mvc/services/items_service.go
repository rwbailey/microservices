package services

import (
	"net/http"

	"github.com/rwbailey/microservices/mvc/domain"
	"github.com/rwbailey/microservices/mvc/utils"
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
