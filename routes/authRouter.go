package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/tuasegun/golang-jwt-project/controllers"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	// Your authentication routes setup here
	incomingRoutes.POST("users/signup", controller.Signup())
	incomingRoutes.POST("users/login", controller.Login())
}
