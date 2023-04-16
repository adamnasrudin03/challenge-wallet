package router

import (
	"adamnasrudin03/challenge-wallet/app/controller"
	"adamnasrudin03/challenge-wallet/app/middlewares"

	"github.com/gin-gonic/gin"
)

func TransactionRouter(e *gin.Engine, txHandler controller.TransactionController) {

	txRoutes := e.Group("/api/v1/transactions")
	{
		txRoutes.Use(middlewares.Authentication())
		txRoutes.POST("/", middlewares.CheckAuthorization(), txHandler.Create)
	}
	topUpRoutes := e.Group("/api/v1/top-up")
	{
		topUpRoutes.Use(middlewares.Authentication())
		topUpRoutes.POST("/", middlewares.CheckAuthorization(), txHandler.TopUp)
	}
}
