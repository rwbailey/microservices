package domain

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rwbailey/microservices/utils"
)

var (
	// mock database data
	users = map[int64]*User{
		123: {Id: 123, FirstName: "Richard", LastName: "Bailey", Email: "richard@example.com"},
	}
)

type userDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct{}

var (
	UserDao userDaoInterface
)

func init() {
	UserDao = &userDao{}
}

func (*userDao) GetUser(userId int64) (*User, *utils.ApplicationError) {
	log.Println("accessing database")
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not found",
	}
}
