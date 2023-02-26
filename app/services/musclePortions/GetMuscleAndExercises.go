package service

import (
	"api/app/models"
	"api/app/repository"
)

type GetPortionWithExercisesService struct {
	PortionsRepository *repository.MusclesPortionRepository
	PortionId          int
}

func (s *GetPortionWithExercisesService) Main() *models.MusclePortion {
	return s.PortionsRepository.FindWithExercises(s.PortionId)
}
