package service

import (
	"api/app/models"
	"api/app/repository"
	"errors"
)

type GetExerciseById struct {
	ExerciseRepository *repository.ExercicioRepository
	ExerciseId         int
	Muscles            bool
	MusclesRole        bool
}

func (s *GetExerciseById) Main() (*models.Exercise, error) {
	var exercise *models.Exercise

	if s.MusclesRole {
		exercise = s.getExerciseWithMuscleRoles()
	} else if s.Muscles {
		exercise = s.getExerciseWithMuscles()
	} else {
		exercise = s.getExercise()
	}

	if exercise.Id == 0 {
		return exercise, errors.New("Exercise not find")
	}

	return exercise, nil
}

func (s *GetExerciseById) getExercise() *models.Exercise {
	return s.ExerciseRepository.FindById(s.ExerciseId)
}

func (s *GetExerciseById) getExerciseWithMuscleRoles() *models.Exercise {
	return s.ExerciseRepository.FindByIdWithMuscles(s.ExerciseId, false, s.MusclesRole)
}
func (s *GetExerciseById) getExerciseWithMuscles() *models.Exercise {
	return s.ExerciseRepository.FindByIdWithMuscles(s.ExerciseId, s.Muscles, false)
}
