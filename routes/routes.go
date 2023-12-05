package routes

import (
	"github.com/gin-gonic/gin"
	userHandler "github.com/trakfy/backend/handlers/users"
)

func SetupRoutes(router *gin.Engine) {
	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	userGroup := router.Group("/user")
	userGroup.GET("/info", userHandler.UserInfo)
}
