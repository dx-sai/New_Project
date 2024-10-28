package models

import (
	"github.com/jinzhu/gorm"
)

type Subscription struct {
	gorm.Model
	UserID            string `json:"user_id"`
	Topic             string `json:"topic"`
	Email             string `json:"email"`
	SMS               string `json:"sms"`
	PushNotifications bool   `json:"push_notifications"`
}
