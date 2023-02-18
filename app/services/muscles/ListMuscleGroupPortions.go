package service

import (
	"api/app/models"
	"api/app/repository"
	"errors"
	"net/url"
	"strconv"
	"strings"
)

type ListMusclesGroupPortionService struct {
	MuscleGroupRepository   *repository.MusclesGroupsRepository
	MusclePortionRepository *repository.MusclesPortionRepository
	Params                  url.Values
}

func (s *ListMusclesGroupPortionService) Main() {

}

func (s *ListMusclesGroupPortionService) getMusclesGroup() (*models.Muscle, error) {
	var muscleGroup models.Muscle
	groupId, ok := s.Params["MuscleGroup_id"]
	if !ok {
		return nil, errors.New("Not find muscleGRoup_id in params")
	}

	groupIdInt, ok := strconv.Atoi(strings.Join(groupId, " "))
	s.MuscleGroupRepository.First(groupId)
	return &muscleGroup, nil
}

func (s *ListMusclesGroupPortionService) getMUscleGroupPortions() (*[]models.MusclePortion, error) {
	var musclePortions []models.MusclePortion
	return &musclePortions, nil
}
