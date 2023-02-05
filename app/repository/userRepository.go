package repository

import (
	"api/app/helpers/errors"
	"api/app/models"
	"fmt"

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
	fmt.Println(users, &users)
	result := r.Db.Find(&users)
	errors.CheckPanic(result.Error)
	return &users
}
