package app

import (
	"github.com/rwbailey/microservices/controllers"
)

func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
}
