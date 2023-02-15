package service

import (
	"api/app/models"
	"api/app/repository"
	"net/url"
)

type ListExerciciosService struct {
	ExercicioRepository *repository.ExercicioRepository
	ExercicioParams     url.Values
}

func (s *ListExerciciosService) Main() *[]models.Exercise {
	return s.ExercicioRepository.FindAll(s.ExercicioParams)
}
