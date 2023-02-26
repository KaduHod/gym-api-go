package service

import (
	"api/app/models"
	"api/app/repository"
	"net/url"
)

type ListMusclesGroupsService struct {
	MusclesGroupsRepository *repository.MusclesGroupsRepository
	Params                  url.Values
	Portions                bool
}

func (s *ListMusclesGroupsService) Main() *[]models.Muscle {
	if s.Portions {
		return s.MusclesGroupsRepository.FindAllWithPortions(s.Params)
	}
	return s.MusclesGroupsRepository.FindAll(s.Params)
}
