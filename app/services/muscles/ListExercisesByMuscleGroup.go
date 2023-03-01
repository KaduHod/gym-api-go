package service

import (
	"api/app/models"
	"api/app/repository"
	"errors"
)

type ListExercisesByMuscleGroupService struct {
	MuscleGroupRepository *repository.MusclesGroupsRepository
	MuscleGroupName       string
}

func (s *ListExercisesByMuscleGroupService) Main() (*models.Muscle, error) {
	var muscle *models.Muscle
	muscle = s.MuscleGroupRepository.FindByNameWithPortionAndExercises(s.MuscleGroupName)

	if muscle.Id == 0 {
		return muscle, errors.New("Muscle not found")
	}
	return muscle, nil
}
