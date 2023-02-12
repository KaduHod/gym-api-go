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

	if s.Personal.Cellphone == "" {
		s.Personal.Cellphone = queryPersonal.Cellphone
	}
	if s.Personal.Email == "" {
		s.Personal.Email = queryPersonal.Email
	}
	if s.Personal.Name == "" {
		s.Personal.Name = queryPersonal.Name
	}
	if s.Personal.Nickname == "" {
		s.Personal.Nickname = queryPersonal.Nickname
	}
	if s.Personal.Password == "" {
		s.Personal.Password = queryPersonal.Password
	}

	s.PersonalRepository.Update(s.Personal)
	return nil

}
