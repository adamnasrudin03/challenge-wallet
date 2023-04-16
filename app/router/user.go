package router

import (
	"adamnasrudin03/challenge-wallet/app/controller"
	"adamnasrudin03/challenge-wallet/app/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRouter(e *gin.Engine, userController controller.UserController) {
	authRoutes := e.Group("/api/v1/auth")
	{
		authRoutes.POST("/register", userController.Register)
		authRoutes.POST("/login", userController.Login)
	}

	userRoutes := e.Group("/api/v1/users")
	{
		userRoutes.Use(middlewares.Authentication())
		userRoutes.GET("/my-wallet", middlewares.CheckAuthorization(), userController.MyWallet)
	}
}
