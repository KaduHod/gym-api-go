package repository

import "api/app/models"

type UserRepositoryI interface {
	Save(user models.User)
	Update(user models.User)
	Delete(user models.User)
	FindById(userId int64)
	FindAll() []models.User
}
