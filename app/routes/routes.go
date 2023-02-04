package routes

import (
	"gymapp/app/controllers/users"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	usersGroup(router)
}

func usersGroup(router *gin.Engine) {
	group := router.Group("/users")
	{
		group.GET("/", users.Index)
		group.POST("/", users.Create)
		group.PUT("/", users.Update)
		group.DELETE("/", users.Delete)
	}
}
