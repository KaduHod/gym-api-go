package muscles

import (
	"api/app/config"
	stringHelpers "api/app/helpers/string"
	"api/app/repository"
	services "api/app/services/musclePortions"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	db := config.DatabaseConnection()
	musclesRepository := repository.NewMusclePortionRepository(db)
	c.JSON(200, gin.H{
		"muscles": musclesRepository.All(),
	})
}

func Exercises(c *gin.Context) {
	db := config.DatabaseConnection()
	portionId, err := stringHelpers.ToInteger(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Param portionId error",
		})
	}
	musclesRepository := repository.NewMusclePortionRepository(db)
	service := services.GetPortionWithExercisesService{
		PortionsRepository: &musclesRepository,
		PortionId:          portionId,
	}

	c.JSON(200, gin.H{
		"muscle": service.Main(),
	})
}
