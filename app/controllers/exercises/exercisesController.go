package exercises

import (
	"api/app/config"
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
