package repository

import (
	apiErrors "api/app/helpers/errors"
	databaseErrors "api/app/helpers/errors/database"
	"api/app/models"
	"errors"
	"fmt"
	"net/url"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type PersonalRepository struct {
	Db *gorm.DB
}

func NewPersonalRepository(Db *gorm.DB) PersonalRepository {
	return PersonalRepository{
		Db: Db,
	}
}
func (r *PersonalRepository) Create(personal *models.Personal) error {
	result := r.Db.Table("users").Create(&personal)
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

func (r *PersonalRepository) First(id int, personal *models.Personal) error {
	result := r.Db.Table("users").Select([]string{"users.*"}).Joins("JOIN users_permissions on users.id = users_permissions.user_id").Where("users_permissions.permission_id =?", PermissionTypes["Personal"]).Where("users.id = ?", id).Find(&personal)
	apiErrors.Check(result.Error)
	if personal.Id == 0 {
		return errors.New("Personal not found!")
	}
	return nil
}

func (r *PersonalRepository) Update(personalParams *models.Personal) {
	fmt.Println("aqui", personalParams)
	r.Db.Table("users").Updates(personalParams)
}

func (r *PersonalRepository) FindAll(params url.Values) []models.Personal {
	var personais []models.Personal
	query := r.Db.Table("users").Select([]string{"users.*"}).Joins("JOIN users_permissions on users.id = users_permissions.user_id").Where("users_permissions.permission_id =?", 2)

	if params != nil {
		for key, value := range params {
			whereClause := "users." + key + " = ?"
			query.Where(whereClause, value)
		}
	}

	result := query.Find(&personais)
	apiErrors.CheckPanic(result.Error)
	return personais
}
