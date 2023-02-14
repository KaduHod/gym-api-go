package routes

import (
	"api/app/controllers/alunos"
	"api/app/controllers/personal"
	"api/app/controllers/users"
	"api/app/middlewares/validate"
	"api/app/models"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {

	groupAlunos := router.Group("/aluno")
	groupAlunos.GET("/", alunos.All)
	groupAlunos.POST("/", validate.CreateUserMiddleware[models.Aluno], alunos.Create)
	groupAlunos.PUT("/", alunos.Update)

	groupPersonal := router.Group("/personal")
	groupPersonal.GET("/", personal.All)
	groupPersonal.POST("/", validate.CreateUserMiddleware[models.Personal], personal.Create)
	groupPersonal.PUT("/", personal.Update)

	groupUsers := router.Group("/user")
	groupUsers.GET("/", users.Index)
	groupUsers.POST("/", validate.CreateUserMiddleware[models.User], users.Create)
	groupUsers.PUT("/", users.Update)
	groupUsers.DELETE("/", users.Delete)

}
