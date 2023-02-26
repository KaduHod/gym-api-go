package service

import (
	"api/app/models"
	"api/app/repository"
)

type ListExercisesByMuscleGroupService struct {
	MuscleGroupRepository *repository.MusclesGroupsRepository
	MuscleGroupName       string
}

func (s *ListExercisesByMuscleGroupService) Main() *models.Muscle {
	return s.MuscleGroupRepository.FindByNameWithPortionAndExercises(s.MuscleGroupName)
}
