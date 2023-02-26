package exercises

import (
	"api/app/config"
	errorsHelper "api/app/helpers/errors"
	stringHelper "api/app/helpers/string"
	"api/app/repository"
	service "api/app/services/exercicios"
	"strconv"

	"github.com/gin-gonic/gin"
)

func All(c *gin.Context) {
	Db := config.DatabaseConnection()
	exerciseRepository := repository.NewExercicioRepository(Db)
	listExercisesService := service.ListExerciciosService{
		ExercicioRepository: &exerciseRepository,
		ExercicioParams:     c.Request.URL.Query(),
	}

	c.JSON(200, gin.H{
		"message": listExercisesService.Main(),
	})
}

func GetExercise(c *gin.Context) {
	muscleBool := false
	muscleRoleBool := false
	muscle := c.Query("muscle")
	muscleRole := c.Query("muscleRole")
	exerciseId := c.Param("exerciseId")
	exerciseIdInt, err := stringHelper.ToInteger(exerciseId)

	if err != nil {
		errorsHelper.Check(err)
	}

	if muscle == "true" || muscle == "1" {
		muscleBool = true
	}

	if muscleRole == "true" || muscle == "1" {
		muscleRoleBool = true
	}

	dB := config.DatabaseConnection()
	exerciseRepository := repository.NewExercicioRepository(dB)

	service := service.GetExerciseById{
		ExerciseRepository: &exerciseRepository,
		ExerciseId:         exerciseIdInt,
		Muscles:            muscleBool,
		MusclesRole:        muscleRoleBool,
	}

	c.JSON(200, gin.H{
		"exercise": service.Main(),
	})
}

func ListByMuscle(c *gin.Context) {
	Db := config.DatabaseConnection()
	muscleIdstr := c.Param("muscleId")
	muscleId, err := strconv.Atoi(muscleIdstr)
	if err != nil {
		c.String(500, "Server error")
	}
	exerciseRepository := repository.NewExercicioRepository(Db)
	listExercisesByMuscleService := service.ListExercisesByMuscleService{
		ExerciseRepository: &exerciseRepository,
		MuscleId:           muscleId,
	}

	exercises := listExercisesByMuscleService.Main()
	c.JSON(200, gin.H{
		"exercises": exercises,
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
