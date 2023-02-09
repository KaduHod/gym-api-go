package alunos

import (
	"api/app/config"
	"api/app/helpers/requests"
	"api/app/models"
	"api/app/repository"
	alunoServices "api/app/services/alunos"

	"github.com/gin-gonic/gin"
)

func All(c *gin.Context) {
	db := config.DatabaseConnection()
	alunoRepo := repository.NewAlunosRepository(db)

	listAlunosService := alunoServices.ListAlunosService{
		AlunoRepository: &alunoRepo,
		Params:          c.Request.URL.Query(),
	}

	c.JSON(200, gin.H{
		"alunos": listAlunosService.Main(),
	})
}

func Create(c *gin.Context) {
	aluno := models.User{}
	requests.GetBodyJson(c.Request.Body, &aluno)

	db := config.DatabaseConnection()
	alunoRepository := repository.NewAlunosRepository(db)
	permissionRepository := repository.NewPermissionRepository(db)

	createAlunoService := alunoServices.CreateAlunoService{
		AlunoRepository:      &alunoRepository,
		PermissionRepository: &permissionRepository,
		Aluno:                &aluno,
	}
	erro := createAlunoService.Main()

	if erro != nil {
		c.JSON(400, gin.H{
			"error": erro.Error(),
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "Created",
		"aluno":   aluno,
	})
}

func Update(c *gin.Context) {
	var alunoUpdateParams models.Aluno
	requests.GetBodyJson(c.Request.Body, &alunoUpdateParams)
	db := config.DatabaseConnection()

	alunoRepository := repository.NewAlunosRepository(db)

	updateService := alunoServices.UpdateAlunoService{
		AlunoRepository: &alunoRepository,
		AlunoParams:     &alunoUpdateParams,
	}

	alunoUpdated := updateService.Main()

	c.JSON(200, gin.H{
		"message": alunoUpdated,
	})
}

func Delete(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Deleting user...",
	})
}
