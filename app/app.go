package app

import (
	"net/http"

	"github.com/rwbailey/microservices/controllers"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

// StartApp startes the application
func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
