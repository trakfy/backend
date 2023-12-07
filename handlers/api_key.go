package handlers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/trakfy/backend/db"
	"github.com/trakfy/backend/models"
	"github.com/trakfy/backend/utils"
)

type CreateApiKeyRequest struct {
	ApiID uuid.UUID `json:"api_id" binding:"required"`
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

	apiKey := &models.ApiKey{}
	err = db.DB.Where("user_id = ? AND api_id = ? and valid = true", user.ID, body.ApiID).First(apiKey).Error
	if err == nil {
		apiKey.Valid = false
		err = db.DB.Save(apiKey).Error
		if err != nil {
			c.JSON(500, gin.H{"error": "Internal Server Error"})
			return
		}
	}

	apiKey = &models.ApiKey{
		ID:          utils.GenerateUUID(),
		ApiID:       body.ApiID,
		UserID:      user.ID,
		Key:         utils.GenerateSHA256(),
		Valid:       true,
		QuotaUsed:   apiKey.QuotaUsed,
		RenewalDate: time.Now().AddDate(0, 1, 0),
	}
	err = db.DB.Create(apiKey).Error
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(200, gin.H{
		"id":           apiKey.ID,
		"api_id":       apiKey.ApiID,
		"user_id":      apiKey.UserID,
		"key":          apiKey.Key,
		"valid":        apiKey.Valid,
		"quota_used":   apiKey.QuotaUsed,
		"renewal_date": apiKey.RenewalDate,
	})
}
