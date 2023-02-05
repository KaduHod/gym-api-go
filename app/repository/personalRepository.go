package repository

import (
	apiErrors "api/app/helpers/errors"
	"api/app/models"
	"fmt"
	"net/url"

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
func (r *PersonalRepository) Create() {

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
	fmt.Println(personais)
	apiErrors.CheckPanic(result.Error)
	return personais
}
