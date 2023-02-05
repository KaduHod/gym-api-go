package repository

import (
	"api/app/helpers/errors"
	"api/app/models"
	"net/url"

	"gorm.io/gorm"
)

type AlunosRepository struct {
	Db *gorm.DB
}

func NewAlunosRepository(Db *gorm.DB) AlunosRepository {
	return AlunosRepository{Db: Db}
}

func (r *AlunosRepository) FindAll(params url.Values) *[]models.User {
	var alunos []models.User
	query := r.Db.Table("users").Select([]string{"users.*"}).Joins("JOIN users_permissions on users.id = users_permissions.user_id").Where("users_permissions.permission_id =?", 1)

	for key, value := range params {
		whereClause := "users." + key + " = ?"
		query.Where(whereClause, value)
	}

	result := query.Find(&alunos)
	errors.CheckPanic(result.Error)
	return &alunos
}
