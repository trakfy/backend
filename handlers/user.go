package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/trakfy/backend/db"
	"github.com/trakfy/backend/models"
	"github.com/trakfy/backend/utils"
)

func UserInfo(c *gin.Context) {
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
		userEntity := &models.User{
			ID:    utils.GenerateUUID(),
			Email: email,
		}

		err := db.DB.Create(userEntity).Error
		if err != nil {
			c.JSON(500, gin.H{"error": "Internal Server Error"})
			return
		}

		user = userEntity
	}

	c.JSON(200, gin.H{
		"id":    user.ID,
		"email": user.Email,
	})
}
