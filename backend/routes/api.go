package routes

import (
	"github.com/sharipovr/pflDockerBilling/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/users", handlers.GetUsers)
		api.POST("/users", handlers.CreateUser)

		api.GET("/subscriptions", handlers.GetSubscriptions)
		api.POST("/subscriptions", handlers.CreateSubscription)
	}

	return router
}
