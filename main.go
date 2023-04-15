package main

import (
	"fmt"
	"net/http"

	"adamnasrudin03/challenge-wallet/app"
	"adamnasrudin03/challenge-wallet/app/configs"
	"adamnasrudin03/challenge-wallet/app/controller"
	routers "adamnasrudin03/challenge-wallet/app/router"
	"adamnasrudin03/challenge-wallet/pkg/database"
	"adamnasrudin03/challenge-wallet/pkg/helpers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = database.SetupMySQLConnection()

	repo     = app.WiringRepository(db)
	services = app.WiringService(repo)

	userController controller.UserController = controller.NewUserController(services)
)

func main() {
	defer database.CloseMySQLConnection(db)
	config := configs.GetInstance()

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.Default())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, helpers.APIResponse("welcome its server", http.StatusOK, nil))
	})

	// Route here
	routers.UserRouter(router, userController)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, helpers.APIResponse("page not found", http.StatusNotFound, nil))
	})

	listen := fmt.Sprintf(":%v", config.Appconfig.Port)
	router.Run(listen)
}
