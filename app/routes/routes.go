package routes

import (
	"api/app/controllers/alunos"
	"api/app/controllers/exercises"
	"api/app/controllers/muscles"
	"api/app/controllers/personal"
	"api/app/controllers/users"
	"api/app/middlewares/validate"
	"api/app/models"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {

	groupAlunos := router.Group("/alunos")
	groupAlunos.GET("/", alunos.All)
	groupAlunos.POST("/", validate.CreateUserMiddleware[models.Aluno], alunos.Create)
	groupAlunos.PUT("/", validate.UpdateUserMiddleware[models.Aluno], alunos.Update)

	groupPersonal := router.Group("/personais")
	groupPersonal.GET("/", personal.All)
	groupPersonal.POST("/", validate.CreateUserMiddleware[models.Personal], personal.Create)
	groupPersonal.PUT("/", validate.UpdateUserMiddleware[models.Personal], personal.Update)

	groupExercises := router.Group("/exercicios")
	groupExercises.GET("/", exercises.All)

	groupUsers := router.Group("/users")
	groupUsers.GET("/", users.Index)
	groupUsers.POST("/", validate.CreateUserMiddleware[models.User], users.Create)
	groupUsers.PUT("/", users.Update)
	groupUsers.DELETE("/", users.Delete)

	groupMuscles := router.Group("muscles")
	groupMuscles.GET("/", muscles.List)

}
