package users

import (
	"api/app/config"
	"api/app/repository"
	"api/app/services"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	db := config.DatabaseConnection()
	userRepo := repository.NewUserRepository(db)

	getAllUsersService := services.GetAllUsersService{
		UserRepository: &userRepo,
	}

	c.JSON(200, gin.H{
		"users": getAllUsersService.GetAllUsers(),
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
