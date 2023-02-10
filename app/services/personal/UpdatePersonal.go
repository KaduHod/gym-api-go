package service

import (
	"api/app/models"
	"api/app/repository"
)

type UpdatePersonalService struct {
	PersonalRepository *repository.PersonalRepository
	Personal           *models.Personal
}

func (s *UpdatePersonalService) Main() error {
	var queryPersonal models.Personal
	error := s.PersonalRepository.First(s.Personal.Id, &queryPersonal)
	if error != nil {
		return error
	}

	s.PersonalRepository.Update(s.Personal)
	return nil
}
