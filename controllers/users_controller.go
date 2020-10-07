package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/rwbailey/microservices/services"
)

func GetUser(res http.ResponseWriter, req *http.Request) {
	// Get the values needed from the request
	userIdParam := req.URL.Query().Get("user_id")
	userId, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		// Return bad request to the client
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("user_id must be a number"))
		return
	}

	// Pass values to service layer
	user, err := services.GetUser(userId)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte(err.Error()))
		// handle error and return to client
		return
	}

	// return user to client
	jsonValue, _ := json.Marshal(user)
	res.Write(jsonValue)
}
