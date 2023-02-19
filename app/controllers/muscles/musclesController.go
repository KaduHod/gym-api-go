package muscles

import (
	"api/app/config"
	"api/app/repository"
	service "api/app/services/muscles"
	"strconv"

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
	muscleGroupId := c.Param("id")
	groupIdInt, erro := strconv.Atoi(muscleGroupId)
	if erro != nil {
		c.JSON(400, gin.H{
			"muscles": erro,
		})
		c.Abort()
		return
	}

	MusclesGroupsRepository := repository.NewMuscelRepository(db)
	MusclesPortionREository := repository.NewMusclePortionRepository(db)
	listMusclesPortionService := service.ListMusclesGroupPortionService{
		MuscleGroupRepository:   &MusclesGroupsRepository,
		MusclePortionRepository: &MusclesPortionREository,
		MuscleGroupId:           groupIdInt,
	}

	muscles, erro := listMusclesPortionService.Main()

	if erro != nil {
		c.String(400, "Muscle group not find")
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"muscles": muscles,
	})
}
