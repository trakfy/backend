package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"github.com/trakfy/backend/handlers"
)

func SetupRoutes(router *gin.Engine) {
	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	// Crons
	cron := cron.New()
	cron.AddFunc("@daily", handlers.RenewalApiKeyCron)
	cron.Start()

	// Routes
	userGroup := router.Group("/user")
	userGroup.GET("/", handlers.UserInfo)

	apiGroup := router.Group("/api")
	apiGroup.POST("/", handlers.CreateApi)

	apiPlanGroup := router.Group("/api_plan")
	apiPlanGroup.POST("/", handlers.CreateApiPlan)

	apiSubscriptionGroup := router.Group("/api_subscription")
	apiSubscriptionGroup.POST("/", handlers.CreateApiSubscription)

	apiKeyGroup := router.Group("/api_key")
	apiKeyGroup.POST("/", handlers.CreateApiKey)
	apiKeyGroup.PATCH("/renewal", handlers.RenewalApiKeyRoute)
}
