package handlers

import (
	"net/http"

	"github.com/sharipovr/pflDockerBilling/models"
	"github.com/sharipovr/pflDockerBilling/storage"

	"github.com/gin-gonic/gin"
)

// Список подписок пользователя
func GetSubscriptions(c *gin.Context) {
	var subscriptions []models.Subscription
	storage.DB.Find(&subscriptions)
	c.JSON(http.StatusOK, subscriptions)
}

// Создание подписки
func CreateSubscription(c *gin.Context) {
	var sub models.Subscription
	if err := c.ShouldBindJSON(&sub); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	storage.DB.Create(&sub)
	c.JSON(http.StatusCreated, sub)
}
