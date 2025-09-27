package controllers

import (
	"github.com/36kone/api-go/database"
	"github.com/36kone/api-go/models"
	"github.com/gin-gonic/gin"
)

func ReadAllUsers(c *gin.Context) {
	var users []models.User
	result := database.DB.Find(&users)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.JSON(200, users)
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	database.DB.Create(&user)
	c.JSON(200, user)
}
