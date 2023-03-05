package alunos

import (
	"api/app/config"
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

func Test(c *gin.Context) {
	db := config.DatabaseConnection()
	alunoRepo := repository.NewAlunosRepository(db)

	c.JSON(200, gin.H{
		"data": alunoRepo.FindFirstBy(map[string]interface{}{
			"email": "TesteGolangService@mail.com",
		}),
	})
}

func Create(c *gin.Context) {
	aluno := c.MustGet("UserParams").(models.Aluno)
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
		"aluno": aluno,
	})
}

func Update(c *gin.Context) {
	alunoParams := c.MustGet("UserParams").(models.Aluno)
	db := config.DatabaseConnection()
	alunoRepository := repository.NewAlunosRepository(db)
	updateService := alunoServices.UpdateAlunoService{
		AlunoRepository: &alunoRepository,
		AlunoParams:     &alunoParams,
	}

	err := updateService.Main()

	if err != nil && err.Error() == "Aluno not found!" {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "alunoParams",
	})
}

func Delete(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Deleting user...",
	})
}
