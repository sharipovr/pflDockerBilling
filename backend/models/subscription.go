package models

import "gorm.io/gorm"

type Subscription struct {
	gorm.Model
	UserID uint   `json:"user_id"`
	Plan   string `json:"plan"`
	Status string `json:"status"` // active, inactive, cancelled
}
