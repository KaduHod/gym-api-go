package muscles

import (
	"api/app/config"
	"api/app/repository"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	db := config.DatabaseConnection()
	musclesRepository := repository.NewMusclePortionRepository(db)
	c.JSON(200, gin.H{
		"muscles": musclesRepository.All(),
	})
}
