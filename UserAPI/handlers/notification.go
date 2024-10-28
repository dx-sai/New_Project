package handlers

import (
	"UserAPI/database"
	"UserAPI/kafka"
	"UserAPI/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SubscribePayload struct {
	UserID               string   `json:"user_id"`
	Topics               []string `json:"topics"`
	NotificationChannels struct {
		Email             string `json:"email"`
		SMS               string `json:"sms"`
		PushNotifications bool   `json:"push_notifications"`
	} `json:"notification_channels"`
}

type UnsubscribePayload struct {
	UserID string   `json:"user_id"`
	Topics []string `json:"topics"`
}

type NotificationMessage struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type SendNotificationPayload struct {
	Topic   string                 `json:"topic"`
	Event   map[string]interface{} `json:"event"`
	Message NotificationMessage    `json:"message"`
}

func Subscribe(c *gin.Context) {
	var payload SubscribePayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Storingeach and every topic subscription in the database
	for _, topic := range payload.Topics {
		subscription := models.Subscription{
			UserID:            payload.UserID,
			Topic:             topic,
			Email:             payload.NotificationChannels.Email,
			SMS:               payload.NotificationChannels.SMS,
			PushNotifications: payload.NotificationChannels.PushNotifications,
		}
		database.DB.Create(&subscription)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Subscribed successfully to topics"})
}

func Unsubscribe(c *gin.Context) {
	var payload UnsubscribePayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, topic := range payload.Topics {
		database.DB.Where("user_id = ? AND topic = ?", payload.UserID, topic).Delete(&models.Subscription{})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Unsubscribed successfully from topics"})
}

func FetchSubscriptions(c *gin.Context) {
	userID := c.Param("user_id")
	var subscriptions []models.Subscription
	database.DB.Where("user_id = ?", userID).Find(&subscriptions)

	c.JSON(http.StatusOK, gin.H{"subscriptions": subscriptions})
}

func SendNotification(c *gin.Context) {
	var payload SendNotificationPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Fetch all subscribers for the topic
	var subscriptions []models.Subscription
	result := database.DB.Where("topic = ?", payload.Topic).Find(&subscriptions)

	// Check how many subscriptions are found
	fmt.Printf("Found %d subscriptions for topic %s\n", len(subscriptions), payload.Topic)

	if result.Error != nil {
		println("Error fetching subscriptions:", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch subscriptions"})
		return
	}

	for _, subscription := range subscriptions {
		// Print the entire subscription data
		fmt.Printf("Subscription data: %+v\n", subscription)

		if subscription.Email != "" {
			println("Calling sendEmailNotification for email:", subscription.Email)
			sendEmailNotification(subscription.Email, payload.Message.Title, payload.Message.Body)
		} else {
			println("No email found for this subscription")
		}

		if subscription.SMS != "" {
			sendSMSNotification(subscription.SMS, payload.Message.Body)
		}

		if subscription.PushNotifications {
			sendPushNotification(subscription.UserID, payload.Message.Body)
		}
	}

	kafka.SendMessage(payload.Topic, fmt.Sprintf("Notification: %s", payload.Message.Body))

	c.JSON(http.StatusOK, gin.H{"message": "Notification sent to subscribers"})
}

// sending notifications to different channels

func sendEmailNotification(email, title, body string) {

	println("Sending email to:", email)
	println("Title:", title)
	println("Body:", body)
}

func sendSMSNotification(sms, message string) {
	println("Sending SMS to:", sms, "with message:", message)
}

func sendPushNotification(userID, message string) {
	println("Sending in-app notification to user:", userID, "with message:", message)
}
