package services

import (
	"api/app/models"
	"api/app/repository"
)

type GetAllUsersService struct {
	UserRepository *repository.UserRepository
}

// Get all users in database
func (s *GetAllUsersService) GetAllUsers() []models.User {
	return *s.UserRepository.FindAll()
}

// Verify data received from request and create a user if there is not some error
func (s *GetAllUsersService) CreateUser() {

}
