package router

import (
	"adamnasrudin03/challenge-wallet/app/controller"

	"github.com/gin-gonic/gin"
)

func UserRouter(e *gin.Engine, userController controller.UserController) {
	userRoutes := e.Group("/api/v1/auth")
	{
		userRoutes.POST("/register", userController.Register)
		userRoutes.POST("/login", userController.Login)
	}
}
