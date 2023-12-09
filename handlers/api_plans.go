package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/trakfy/backend/db"
	"github.com/trakfy/backend/models"
	"github.com/trakfy/backend/utils"
)

type CreateApiPlanRequest struct {
	ApiID        string `json:"api_id" binding:"required"`
	Name         string `json:"name" binding:"required"`
	ValueCents   int64  `json:"value_cents"`
	RequestLimit int64  `json:"request_limit" binding:"required"`
}

func CreateApiPlan(c *gin.Context) {
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

	var body CreateApiPlanRequest
	err = c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	api := &models.Api{}
	err = db.DB.Where("user_id = ? AND id = ?", user.ID, body.ApiID).First(api).Error
	if err != nil {
		c.JSON(401, gin.H{"error": "Api not found"})
		return
	}

	apiPlan := &models.ApiPlan{
		ID:           utils.GenerateUUID(),
		ApiID:        api.ID,
		Name:         body.Name,
		ValueCents:   body.ValueCents,
		RequestLimit: body.RequestLimit,
	}
	err = db.DB.Create(apiPlan).Error
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(200, apiPlan)

}
