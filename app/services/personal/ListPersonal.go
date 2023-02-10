package service

import (
	"api/app/models"
	"api/app/repository"
	"net/url"
)

type ListPersonalService struct {
	PersonalRepository *repository.PersonalRepository
	Params             url.Values
}

func (s *ListPersonalService) Main() []models.Personal {
	return s.PersonalRepository.FindAll(s.Params)
}
