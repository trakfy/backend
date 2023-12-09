package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/trakfy/backend/db"
	"github.com/trakfy/backend/models"
	"github.com/trakfy/backend/utils"
)

type CreateApiSubscriptionRequest struct {
	ApiPlanID string `json:"api_plan_id" binding:"required"`
}

func CreateApiSubscription(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	token = token[len("Bearer "):]
	claims, err := utils.DecodeToken(token)
	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	email := claims.Email
	user := &models.User{}
	err = db.DB.Where("email = ?", email).First(user).Error
	if err != nil {
		c.JSON(401, gin.H{"error": "User not found"})
		return
	}

	var body CreateApiSubscriptionRequest
	err = c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	apiPlan := &models.ApiPlan{}
	err = db.DB.Where("id = ?", body.ApiPlanID).First(apiPlan).Error
	if err != nil {
		c.JSON(401, gin.H{"error": "Api Plan not found"})
		return
	}

	apiSubscription := &models.ApiSubscription{}
	err = db.DB.Where("user_id = ? AND api_plan_id = ?", user.ID, body.ApiPlanID).First(apiSubscription).Error
	if err == nil {
		c.JSON(401, gin.H{"error": "Api Subscription already exists"})
		return
	}

	apiSubscription = &models.ApiSubscription{
		ID:        utils.GenerateUUID(),
		UserID:    user.ID,
		ApiPlanID: apiPlan.ID,
	}
	err = db.DB.Create(apiSubscription).Error
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(200, apiSubscription)
}
