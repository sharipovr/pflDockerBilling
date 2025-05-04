package handlers

import (
	"errors"
	"net/http"
	"os"

	"github.com/sharipovr/pflDockerBilling/models"
	"github.com/sharipovr/pflDockerBilling/storage"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/customer"
	"github.com/stripe/stripe-go/v78/paymentmethod"
	"github.com/stripe/stripe-go/v78/subscription"
)

type SubscriptionRequest struct {
	Email   string `json:"email"`
	PriceID string `json:"price_id"`
}

func CreateStripeSubscription(c *gin.Context) {
	var req SubscriptionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	// 1. Проверяем наличие пользователя в локальной базе данных по email
	var user models.User
	if err := storage.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error: " + err.Error()})
		return
	}

	// 2. Создаём Customer в Stripe
	cust, err := customer.New(&stripe.CustomerParams{
		Email: stripe.String(req.Email),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Stripe customer creation failed: " + err.Error()})
		return
	}

	// 3. Подключаем тестовый Payment Method к Customer
	pm, err := paymentmethod.Attach(
		"pm_card_visa",
		&stripe.PaymentMethodAttachParams{
			Customer: stripe.String(cust.ID),
		},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Payment method attachment failed: " + err.Error()})
		return
	}

	// 4. Устанавливаем Payment Method дефолтным
	_, err = customer.Update(cust.ID, &stripe.CustomerParams{
		InvoiceSettings: &stripe.CustomerInvoiceSettingsParams{
			DefaultPaymentMethod: stripe.String(pm.ID),
		},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Set default payment method failed: " + err.Error()})
		return
	}

	// 5. Создаём подписку в Stripe
	stripeSub, err := subscription.New(&stripe.SubscriptionParams{
		Customer: stripe.String(cust.ID),
		Items: []*stripe.SubscriptionItemsParams{
			{Price: stripe.String(req.PriceID)},
		},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Stripe subscription creation failed: " + err.Error()})
		return
	}

	// 6. Создаём локальную подписку с корректным user_id
	localSub := models.Subscription{
		UserID:      user.ID, // Важно! Используем ID найденного пользователя
		Email:       req.Email,
		Plan:        req.PriceID,
		Status:      string(stripeSub.Status),
		StripeSubID: stripeSub.ID,
	}

	if err := storage.DB.Create(&localSub).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Local subscription creation failed: " + err.Error()})
		return
	}

	// 7. Возвращаем успешный ответ
	c.JSON(http.StatusCreated, gin.H{
		"subscription_id": stripeSub.ID,
		"status":          stripeSub.Status,
		"local_id":        localSub.ID,
	})
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
