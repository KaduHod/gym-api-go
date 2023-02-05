package repository

import (
	"api/app/helpers/errors"
	"api/app/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

// User repository constructor
func NewUserRepository(Db *gorm.DB) UserRepository {
	return UserRepository{Db: Db}
}

func (r UserRepository) Save(user models.User) {

}

func (r *UserRepository) FindAll() *[]models.User {
	var users []models.User
	result := r.Db.Find(&users)
	errors.CheckPanic(result.Error)
	return &users
}
