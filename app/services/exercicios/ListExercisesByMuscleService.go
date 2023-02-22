package service

import (
	"api/app/models"
	"api/app/repository"
)

type ListExercisesByMuscleService struct {
	ExerciseRepository *repository.ExercicioRepository
	MuscleId           int
}

func (s *ListExercisesByMuscleService) Main() *[]models.Exercise {
	return s.ExerciseRepository.FindByMuscleId(s.MuscleId)
}
