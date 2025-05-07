package handlers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sharipovr/pflDockerBilling/handlers"
	"github.com/sharipovr/pflDockerBilling/models"
	"github.com/sharipovr/pflDockerBilling/storage"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUserHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Создаем тестовую in-memory базу
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	// Мигрируем модель
	db.AutoMigrate(&models.User{})
	storage.DB = db // Временная подмена storage.DB на тестовую БД

	router := gin.Default()
	router.POST("/users", handlers.CreateUser)

	w := httptest.NewRecorder()
	reqBody := bytes.NewBufferString(`{"name":"Test User","email":"test@example.com"}`)
	req, _ := http.NewRequest("POST", "/users", reqBody)
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	// Дополнительно проверим, что запись появилась в базе данных
	var createdUser models.User
	err = db.First(&createdUser, "email = ?", "test@example.com").Error
	assert.NoError(t, err)
	assert.Equal(t, "Test User", createdUser.Name)
}
