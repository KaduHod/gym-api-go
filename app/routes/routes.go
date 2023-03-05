package routes

import (
	"api/app/controllers/alunos"
	"api/app/controllers/exercises"
	"api/app/controllers/muscles"
	"api/app/controllers/musclesGroups"
	"api/app/controllers/personal"
	"api/app/controllers/users"
	"api/app/controllers/workout"
	"api/app/middlewares/validate"
	"api/app/models"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {

	newFunction(router)

}

func newFunction(router *gin.Engine) {

	router.GET("/TEST", alunos.Test)
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
	groupExercises.GET("/:exerciseId", exercises.GetExercise)
	groupExercises.GET("/muscle/:muscleId", exercises.ListByMuscle)

	groupUsers := router.Group("/users")
	groupUsers.GET("/", users.Index)
	groupUsers.POST("/", validate.CreateUserMiddleware[models.User], users.Create)
	groupUsers.PUT("/", users.Update)
	groupUsers.DELETE("/", users.Delete)

	groupMusclesGroups := router.Group("/muscle-groups")
	groupMusclesGroups.GET("/", musclesGroups.List)
	groupMusclesGroups.GET("/:id/portion", musclesGroups.ListPortions)
	groupMusclesGroups.GET("/exercises/:name", musclesGroups.Exercises)

	groupMusclePortions := router.Group("/muscles")
	groupMusclePortions.GET("/", muscles.List)
	groupMusclePortions.GET("/:id/exercises", muscles.Exercises)

	groupWorkout := router.Group("/workouts")
	groupWorkout.GET("/", workout.List)
	groupWorkout.POST("/", workout.Create)
}
