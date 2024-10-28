package routes

import (
	"UserAPI/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// API Endpoints
	r.POST("/subscribe", handlers.Subscribe)
	r.POST("/unsubscribe", handlers.Unsubscribe)
	r.GET("/subscriptions/:user_id", handlers.FetchSubscriptions)
	r.POST("/notifications/send", handlers.SendNotification)

	return r
}
