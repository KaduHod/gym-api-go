package service

import (
	"api/app/models"
	"api/app/repository"
	"errors"
)

type ListMusclesGroupPortionService struct {
	MuscleGroupRepository   *repository.MusclesGroupsRepository
	MusclePortionRepository *repository.MusclesPortionRepository
	MuscleGroupId           int
}

func (s *ListMusclesGroupPortionService) Main() (*models.Muscle, error) {

	muscle, err := s.getMusclesGroup()

	if err != nil || muscle.Id == 0 {
		return muscle, errors.New("Muscle group not find")
	}

	return muscle, nil
}

func (s *ListMusclesGroupPortionService) getMusclesGroup() (*models.Muscle, error) {
	var muscleGroup *models.Muscle
	muscleGroup = s.MuscleGroupRepository.FindFirstWithPortions(s.MuscleGroupId)
	return muscleGroup, nil
}

func (s *ListMusclesGroupPortionService) getMuscleGroupPortions() (*[]models.MusclePortion, error) {
	var musclePortions *[]models.MusclePortion
	musclePortions = s.MusclePortionRepository.FindByGroupId(s.MuscleGroupId)
	return musclePortions, nil
}
