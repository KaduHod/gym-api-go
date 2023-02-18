package workout

import "github.com/gin-gonic/gin"

func List(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "working",
	})
}

func Create(c *gin.Context) {
	c.JSON(201, gin.H{
		"message": "workout created!",
	})
}
