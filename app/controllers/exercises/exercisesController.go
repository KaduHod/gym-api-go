package exercises

import (
	"api/app/config"
	"api/app/repository"
	service "api/app/services/exercicios"

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
