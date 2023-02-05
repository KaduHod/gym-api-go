package personal

import (
	"api/app/config"
	"api/app/repository"
	service "api/app/services/personal"

	"github.com/gin-gonic/gin"
)

func All(c *gin.Context) {
	db := config.DatabaseConnection()
	personalRepository := repository.NewPersonalRepository(db)
	listPersonaisService := service.ListPersonalService{
		PersonalRepository: &personalRepository,
		Params:             c.Request.URL.Query(),
	}

	c.JSON(200, gin.H{
		"personais": listPersonaisService.Main(),
	})
}

func Create(c *gin.Context) {

}
