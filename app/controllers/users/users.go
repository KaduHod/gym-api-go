package users

import (
	"api/app/config"
	"api/app/repository"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	db := config.DatabaseConnection()
	userRepository := repository.NewUserRepository(db)
	users := userRepository.FindAll()
	c.JSON(200, gin.H{
		"users": users,
	})
}

func Create(c *gin.Context) {
	c.JSON(201, gin.H{
		"message": "Creating user...",
	})
}

func Update(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Updating user...",
	})
}

func Delete(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Deleting user...",
	})
}
