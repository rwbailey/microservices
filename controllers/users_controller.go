package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/rwbailey/microservices/utils"

	"github.com/rwbailey/microservices/services"
)

func GetUser(res http.ResponseWriter, req *http.Request) {
	// Get the values needed from the request
	userIdParam := req.URL.Query().Get("user_id")
	userId, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad request",
		}
		jsonValue, _ := json.Marshal(apiErr)

		res.WriteHeader(apiErr.StatusCode)
		res.Write([]byte(jsonValue))
		return
	}

	// Pass values to service layer
	user, apiErr := services.UsersService.GetUser(userId)
	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		res.WriteHeader(apiErr.StatusCode)
		res.Write([]byte(jsonValue))
		// handle error and return to client
		return
	}

	// return user to client
	jsonValue, _ := json.Marshal(user)
	res.Write(jsonValue)
}
