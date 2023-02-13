package routes

import (
	"api/app/controllers/alunos"
	"api/app/controllers/personal"
	"api/app/controllers/users"
	"api/app/middlewares"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {

	groupAlunos := router.Group("/aluno")
	groupAlunos.GET("/", alunos.All)
	groupAlunos.POST("/", middlewares.CreateUserMiddleware, alunos.Create)
	groupAlunos.PUT("/", alunos.Update)

	groupPersonal := router.Group("/personal")
	groupPersonal.GET("/", personal.All)
	groupPersonal.POST("/", middlewares.CreateUserMiddleware, personal.Create)
	groupPersonal.PUT("/", personal.Update)

	groupUsers := router.Group("/user")
	groupUsers.GET("/", users.Index)
	groupUsers.POST("/", users.Create)
	groupUsers.PUT("/", users.Update)
	groupUsers.DELETE("/", users.Delete)

}
