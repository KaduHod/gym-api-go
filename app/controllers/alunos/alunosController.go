package alunos

import (
	"api/app/config"
	"api/app/repository"

	"github.com/gin-gonic/gin"
)

func All(c *gin.Context) {
	db := config.DatabaseConnection()
	alunoRepository := repository.NewAlunosRepository(db)
	query := c.Request.URL.Query()
	c.JSON(200, gin.H{
		"alunos": alunoRepository.FindAll(query),
	})
}

func Create(c *gin.Context) {
	c.JSON(201, gin.H{
		"message": "Creating user...",
	})
}

func Update(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Updating user...",
	})
}

func Delete(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Deleting user...",
	})
}
