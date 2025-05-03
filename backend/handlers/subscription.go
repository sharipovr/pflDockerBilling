package handlers

import (
	"net/http"

	"github.com/sharipovr/pflDockerBilling/models"
	"github.com/sharipovr/pflDockerBilling/storage"

	"github.com/gin-gonic/gin"
)

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

// Получение подписок пользователя по user_id
func GetSubscriptionsByUser(c *gin.Context) {
	var subscriptions []models.Subscription
	userId := c.Query("user_id")

	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id parameter is required"})
		return
	}

	if err := storage.DB.Where("user_id = ?", userId).Find(&subscriptions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, subscriptions)
}
