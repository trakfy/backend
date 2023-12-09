package handlers

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/trakfy/backend/db"
	"github.com/trakfy/backend/models"
	"github.com/trakfy/backend/utils"
)

type CreateApiKeyRequest struct {
	ApiSubscriptionID uuid.UUID `json:"api_subscription_id" binding:"required"`
}

func CreateApiKey(c *gin.Context) {
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

	var body CreateApiKeyRequest
	err = c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}

	apiSubscription := &models.ApiSubscription{}
	err = db.DB.Where("id = ? AND user_id = ?", body.ApiSubscriptionID, user.ID).First(apiSubscription).Error
	if err != nil {
		c.JSON(401, gin.H{"error": "Api Subscription not found"})
		return
	}

	apiKey := &models.ApiKey{}
	err = db.DB.Where("user_id = ? AND api_subscription_id = ? and valid = true", user.ID, body.ApiSubscriptionID).First(apiKey).Error
	if err == nil {
		apiKey.Valid = false
		err = db.DB.Save(apiKey).Error
		if err != nil {
			c.JSON(500, gin.H{"error": "Internal Server Error"})
			return
		}
	}

	apiKey = &models.ApiKey{
		ID:                utils.GenerateUUID(),
		ApiSubscriptionID: body.ApiSubscriptionID,
		UserID:            user.ID,
		Key:               utils.GenerateSHA256(),
		Valid:             true,
		QuotaUsed:         apiKey.QuotaUsed,
		RenewalDate:       time.Now().AddDate(0, 1, 0),
	}
	err = db.DB.Create(apiKey).Error
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(200, apiKey)
}

func ValidateKey(c *gin.Context) {
	input_key := c.GetHeader("trakfy-api-key")

	api_key := &models.ApiKey{}
	err := db.DB.Where("key = ?", input_key).First(api_key).Error
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid API Key"})
		return
	}

	if !api_key.Valid {
		c.JSON(401, gin.H{"error": "Invalid API Key"})
		return
	}

	user := &models.User{}
	err = db.DB.Where("id = ?", api_key.UserID).First(user).Error
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid API Key"})
		return
	}

	apiSubscription := &models.ApiSubscription{}
	err = db.DB.Where("id = ?", api_key.ApiSubscriptionID).First(apiSubscription).Error
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid API Key"})
		return
	}

	apiPlan := &models.ApiPlan{}
	err = db.DB.Where("id = ?", apiSubscription.ApiPlanID).First(apiPlan).Error
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid API Key"})
		return
	}

	if api_key.QuotaUsed >= int(apiPlan.RequestLimit) {
		c.JSON(401, gin.H{"error": "Quota exceeded"})
		return
	}

	api_key.QuotaUsed += 1
	err = db.DB.Save(api_key).Error
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(200, gin.H{"message": "ok"})
}

func RenewalApiKeyRoute(c *gin.Context) {
	RenewalApiKeyCron()
	c.JSON(200, gin.H{"message": "ok"})
}

func RenewalApiKeyCron() {
	fmt.Println("RenewalApiKey")

	apiKeys := []models.ApiKey{}
	err := db.DB.Where("renewal_date < ? AND valid = true", time.Now()).Find(&apiKeys).Error
	if err != nil {
		return
	}

	for _, apiKey := range apiKeys {
		apiKey.QuotaUsed = 0
		apiKey.RenewalDate = time.Now().AddDate(0, 1, 0)
		err = db.DB.Save(apiKey).Error
		if err != nil {
			return
		}
	}
}
