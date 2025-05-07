package handlers

import (
	"net/http"

	"github.com/sharipovr/pflDockerBilling/models"
	"github.com/sharipovr/pflDockerBilling/storage"

	"github.com/gin-gonic/gin"
)

// Получение списка пользователей
func GetUsers(c *gin.Context) {
	var users []models.User
	storage.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

// Создание пользователя
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := storage.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// Получение пользователя по ID
func GetUserByID(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if err := storage.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetUserByEmail(c *gin.Context) {
	email := c.Param("email")
	var user models.User
	if err := storage.DB.Where("email = ?", email).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}
