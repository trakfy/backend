package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/trakfy/backend/db"
	"github.com/trakfy/backend/models"
	"github.com/trakfy/backend/utils"
)

type CreateApiRequest struct {
	Name string `json:"name" binding:"required"`
}

func CreateApi(c *gin.Context) {
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

	var body CreateApiRequest
	err = c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}

	api := &models.Api{}
	err = db.DB.Where("user_id = ? AND name = ?", user.ID, body.Name).First(api).Error
	if err == nil {
		c.JSON(400, gin.H{"error": "Api already exists"})
		return
	}

	api = &models.Api{
		ID:     utils.GenerateUUID(),
		UserID: user.ID,
		Name:   body.Name,
	}
	err = db.DB.Create(api).Error
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(200, gin.H{
		"id":      api.ID,
		"user_id": api.UserID,
		"name":    api.Name,
	})

}
