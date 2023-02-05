package services

import (
	"api/app/models"
	"api/app/repository"
)

type GetAllUsersService struct {
	UserRepository *repository.UserRepository
}

func (s *GetAllUsersService) GetAllUsers() []models.User {
	return *s.UserRepository.FindAll()
}
