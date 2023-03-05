package personal

import (
	"api/app/config"
	"api/app/models"
	"api/app/repository"
	service "api/app/services/personal"
	"fmt"
	"reflect"

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

func Update(c *gin.Context) {
	personal := c.MustGet("UserParams").(models.Personal)

	db := config.DatabaseConnection()
	personalRepository := repository.NewPersonalRepository(db)
	updateService := service.UpdatePersonalService{
		PersonalRepository: &personalRepository,
		Personal:           &personal,
	}

	erro := updateService.Main()

	if erro != nil && erro.Error() == "Personal not found!" {
		c.JSON(400, gin.H{
			"error": erro.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message":  "updated",
		"personal": personal,
	})

}

func Create(c *gin.Context) {
	personal := c.MustGet("UserParams").(models.Personal)
	fmt.Println(personal, reflect.TypeOf(personal))
	db := config.DatabaseConnection()
	personalRepository := repository.NewPersonalRepository(db)
	permissionRepository := repository.NewPermissionRepository(db)
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
