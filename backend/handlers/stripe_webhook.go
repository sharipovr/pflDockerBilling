package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/webhook"
)

func HandleStripeWebhook(c *gin.Context) {
	const MaxBodyBytes = int64(65536)
	requestBody, err := io.ReadAll(io.LimitReader(c.Request.Body, MaxBodyBytes))
	if err != nil {
		c.AbortWithStatus(http.StatusServiceUnavailable)
		return
	}

	webhookSecret := os.Getenv("STRIPE_WEBHOOK_SECRET")
	sigHeader := c.GetHeader("Stripe-Signature")

	event, err := webhook.ConstructEvent(requestBody, sigHeader, webhookSecret)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	switch event.Type {
	case "invoice.payment_succeeded":
		var invoice stripe.Invoice
		if err := json.Unmarshal(event.Data.Raw, &invoice); err == nil {
			fmt.Printf("✅ Invoice %s успешно оплачена.\n", invoice.ID)
			// Тут можно обновить статус в базе данных
		}
	case "invoice.payment_failed":
		var invoice stripe.Invoice
		if err := json.Unmarshal(event.Data.Raw, &invoice); err == nil {
			fmt.Printf("❌ Ошибка оплаты инвойса %s.\n", invoice.ID)
			// Тут можно отправить уведомление клиенту
		}
	default:
		fmt.Printf("Необработанный тип события: %s\n", event.Type)
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
