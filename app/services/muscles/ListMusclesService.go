package service

import (
	"api/app/models"
	"api/app/repository"
	"net/url"
)

type ListMusclesGroupsService struct {
	MusclesGroupsRepository *repository.MusclesGroupsRepository
	Params                  url.Values
}

func (s *ListMusclesGroupsService) Main() *[]models.Muscle {
	return s.MusclesGroupsRepository.FindAll(s.Params)
}
