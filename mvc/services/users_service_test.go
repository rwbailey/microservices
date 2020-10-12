package services

import (
	"log"
	"net/http"
	"testing"

	"github.com/rwbailey/microservices/mvc/domain"
	"github.com/rwbailey/microservices/mvc/utils"

	"github.com/stretchr/testify/assert"
)

type userDaoMock struct{}

var (
	getUserFunction func(userId int64) (*domain.User, *utils.ApplicationError)
)

func init() {
	domain.UserDao = &userDaoMock{}
}

func (*userDaoMock) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	log.Println("testing with mock data")
	return getUserFunction(123)
}

func TestGetUserNotFoundInDatebase(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			Message:    "user 0 was not found",
			StatusCode: http.StatusNotFound,
		}
	}

	user, err := UsersService.GetUser(0)

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "user 0 was not found", err.Message)
}

func TestGetUserNoError(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return &domain.User{
				Id: 123,
			},
			nil
	}

	user, err := UsersService.GetUser(123)

	assert.NotNil(t, user)
	assert.Nil(t, err)
	assert.EqualValues(t, 123, user.Id)
}
