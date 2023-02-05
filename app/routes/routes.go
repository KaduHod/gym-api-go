package routes

import (
	"api/app/controllers/alunos"
	"api/app/controllers/personal"
	"api/app/controllers/users"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	usersGroup(router)
	alunosGroups(router)
	personalGroup(router)
	router.GET("/", func(c *gin.Context) {
		fmt.Println("oi")
	})
}

func alunosGroups(router *gin.Engine) {
	group := router.Group("/alunos")
	{
		group.GET("/", alunos.All)
		group.POST("/", alunos.Create)
	}
}

func personalGroup(router *gin.Engine) {
	group := router.Group("/personal")
	{
		group.GET("/", personal.All)
		group.POST("/", personal.Create)
	}
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
