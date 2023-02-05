package personal

import (
	"api/app/config"
	"api/app/helpers/errors"
	"api/app/models"
	"api/app/repository"
	service "api/app/services/personal"
	"encoding/json"
	"io/ioutil"

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
	db := config.DatabaseConnection()
	personalRepository := repository.NewPersonalRepository(db)
	permissionRepository := repository.NewPermissionRepository(db)
	jsonBody, err := ioutil.ReadAll(c.Request.Body)
	personal := models.Personal{}
	json.Unmarshal([]byte(jsonBody), &personal)
	errors.Check(err)
	createPersonalService := service.CreatePersonalService{
		PersonalRepository:   &personalRepository,
		PermissionRepository: &permissionRepository,
		Personal:             &personal,
	}

	erro := createPersonalService.Main()
	if erro != nil {
		c.JSON(400, gin.H{
			"error": erro.Error(),
		})
		return
	}

	c.JSON(201, gin.H{
		"message":  "created",
		"personal": personal,
	})
}
