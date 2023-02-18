package muscles

import (
	"api/app/config"
	"api/app/repository"
	service "api/app/services/muscles"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	db := config.DatabaseConnection()
	MusclesGroupsRepository := repository.NewMuscelRepository(db)
	listMusclesGroupsService := service.ListMusclesGroupsService{
		MusclesGroupsRepository: &MusclesGroupsRepository,
		Params:                  c.Request.URL.Query(),
	}

	c.JSON(200, gin.H{
		"muscles": listMusclesGroupsService.Main(),
	})
}

func ListPortions(c *gin.Context) {
	db := config.DatabaseConnection()
	MusclesGroupsRepository := repository.NewMuscelRepository(db)
	listMusclesGroupsService := service.ListMusclesGroupsService{
		MusclesGroupsRepository: &MusclesGroupsRepository,
		Params:                  c.Request.URL.Query(),
	}

	c.JSON(200, gin.H{
		"muscles": listMusclesGroupsService.Main(),
	})
}
