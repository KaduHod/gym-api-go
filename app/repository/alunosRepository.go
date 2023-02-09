package repository

import (
	apiErrors "api/app/helpers/errors"
	databaseErrors "api/app/helpers/errors/database"
	"api/app/models"
	"errors"
	"net/url"

	"github.com/go-sql-driver/mysql"
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

	if params != nil {
		for key, value := range params {
			whereClause := "users." + key + " = ?"
			query.Where(whereClause, value)
		}
	}

	result := query.Find(&alunos)
	apiErrors.CheckPanic(result.Error)
	return &alunos
}

func (r *AlunosRepository) First(id int, aluno *models.Aluno) {
	result := r.Db.Table("users").Select([]string{"users.*"}).Joins("JOIN users_permissions on users.id = users_permissions.user_id").Where("users_permissions.permission_id =?", PermissionTypes["Aluno"]).Where("users.id = ?", id).Find(&aluno)
	apiErrors.Check(result.Error)
}

func (r *AlunosRepository) Update(alunoParams *models.Aluno) {
	r.Db.Table("users").Model(alunoParams).Save(&alunoParams)
}

func (r *AlunosRepository) Create(aluno *models.User) error {
	result := r.Db.Create(&aluno)
	if result.Error != nil {
		err := result.Error.(*mysql.MySQLError)
		hasError, message := databaseErrors.CheckError(err)

		if hasError {
			return errors.New(message)
		}
	}
	apiErrors.Check(result.Error)
	return nil
}
