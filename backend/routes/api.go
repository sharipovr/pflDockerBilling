package routes

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/sharipovr/pflDockerBilling/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	// Настройка CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{os.Getenv("FRONTEND_URL")}, // адрес твоего фронтенда
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type"},
	}))

	api := router.Group("/api")
	{
		api.GET("/users", handlers.GetUsers)
		api.POST("/users", handlers.CreateUser)

		api.POST("/subscriptions", handlers.CreateSubscription)

		api.GET("/users/:id", handlers.GetUserByID)                // теперь принимаем id
		api.GET("/subscriptions", handlers.GetSubscriptionsByUser) // теперь с фильтрацией по юзер id
		api.POST("/subscriptions/stripe", handlers.CreateStripeSubscription)

		api.POST("/webhooks/stripe", handlers.HandleStripeWebhook)
		api.GET("/users/email/:email", handlers.GetUserByEmail)
	}

	return router
}
