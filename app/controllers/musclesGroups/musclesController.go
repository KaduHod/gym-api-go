package musclesGroups

import (
	"api/app/config"
	"api/app/repository"
	service "api/app/services/muscles"
	"strconv"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	db := config.DatabaseConnection()
	portionsBool := false
	portionsParam := c.Query("portions")
	if portionsParam == "true" || portionsParam == "1" {
		portionsBool = true
	}
	MusclesGroupsRepository := repository.NewMuscelRepository(db)
	listMusclesGroupsService := service.ListMusclesGroupsService{
		MusclesGroupsRepository: &MusclesGroupsRepository,
		Params:                  c.Request.URL.Query(),
		Portions:                portionsBool,
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

func Exercises(c *gin.Context) {
	db := config.DatabaseConnection()
	muscleGroupName := c.Param("name")
	MusclesGroupsRepository := repository.NewMuscelRepository(db)

	service := service.ListExercisesByMuscleGroupService{
		MuscleGroupRepository: &MusclesGroupsRepository,
		MuscleGroupName:       muscleGroupName,
	}

	c.JSON(200, gin.H{
		"muscles": service.Main(),
	})
}
