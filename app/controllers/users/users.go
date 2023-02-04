package users

import "github.com/gin-gonic/gin"

func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "List users...",
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
