package service

import (
	"api/app/models"
	"api/app/repository"
	"net/url"
)

type ListMusclesService struct {
	MusclesRepository *repository.MusclesRepository
	Params            url.Values
}

func (s *ListMusclesService) Main() *[]models.Muscle {
	return s.MusclesRepository.FindAll(s.Params)
}
