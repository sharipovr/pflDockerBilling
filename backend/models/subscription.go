package models

import "gorm.io/gorm"

type Subscription struct {
	gorm.Model
	UserID      uint   `json:"user_id"`
	Email       string `json:"email"`
	Plan        string `json:"plan"`
	Status      string `json:"status"`        // active, inactive, cancelled
	StripeSubID string `json:"stripe_sub_id"` // ID подписки в Stripe
}
