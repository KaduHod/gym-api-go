package repository

import (
	apiErrors "api/app/helpers/errors"
	databaseErrors "api/app/helpers/errors/database"
	integersHelper "api/app/helpers/integers"
	"api/app/models"
	"errors"
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
	r.Db.Table("users").Model(&personalParams).Omit("id").Where("users.id = ?", personalParams.Id).Updates(personalParams)
}

func (r *PersonalRepository) FindAll(params url.Values) []models.Personal {
	var personais []models.Personal
	query := r.Db.Select([]string{"users.*"}).Joins("JOIN users_permissions on users.id = users_permissions.user_id").Where("users_permissions.permission_id =?", PermissionTypes["Personal"])

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

func (r *PersonalRepository) FindFirstBy(params map[string]interface{}) *models.Personal {
	var personais models.Personal
	query := r.Db.Select([]string{"users.*"}).Joins("JOIN users_permissions on users.id = users_permissions.user_id").Where("users_permissions.permission_id =?", PermissionTypes["Personal"])

	if params != nil {
		query.Where(params)
	}

	result := query.Find(&personais)
	apiErrors.CheckPanic(result.Error)
	return &personais
}

func (r *PersonalRepository) DeletePersonal(personal *models.Personal) {
	var deletePermissionResult any
	r.Db.Raw("delete from users_permissions where user_id = " + integersHelper.ToString(personal.Id)).Scan(deletePermissionResult)
	var deletePersonalResult any
	r.Db.Delete(personal).Scan(deletePersonalResult)
}
