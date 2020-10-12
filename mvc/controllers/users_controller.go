package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/rwbailey/microservices/mvc/services"
	"github.com/rwbailey/microservices/mvc/utils"
)

func GetUser(c *gin.Context) {
	// Get the values needed from the request
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad request",
		}

		utils.RespondError(c, apiErr)
		return
	}

	// Pass values to service layer
	user, apiErr := services.UsersService.GetUser(userId)
	if apiErr != nil {
		utils.RespondError(c, apiErr)
		return
	}

	// return user to client
	utils.Respond(c, http.StatusOK, user)
}
