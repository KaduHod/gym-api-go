package service

import (
	"api/app/repository"
	"net/url"
)

type CreateWorkoutService struct {
	WorkoutRepository *repository.WorkoutRepository,
	ExerciseRepository *repository.ExercicioRepository,
	WorkoutParams *models.Workout,
	ExercisesParams *models.Exercises
}
