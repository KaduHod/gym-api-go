package muscles

import (
	"api/app/config"
	"api/app/repository"
	service "api/app/services/muscles"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	db := config.DatabaseConnection()
	musclesRepository := repository.NewMuscelRepository(db)
	listMusclesService := service.ListMusclesService{
		MusclesRepository: &musclesRepository,
		Params:            c.Request.URL.Query(),
	}

	c.JSON(200, gin.H{
		"muscles": listMusclesService.Main(),
	})
}
