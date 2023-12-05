package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/trakfy/backend/handlers"
)

func SetupRoutes(router *gin.Engine) {
	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	userGroup := router.Group("/user")
	userGroup.GET("/", handlers.UserInfo)
}
