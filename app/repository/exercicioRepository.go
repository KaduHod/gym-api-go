package repository

import (
	apiErrors "api/app/helpers/errors"
	"api/app/models"
	"net/url"
	"strings"

	"gorm.io/gorm"
)

type ExercicioRepository struct {
	Db *gorm.DB
}

func NewExercicioRepository(Db *gorm.DB) ExercicioRepository {
	return ExercicioRepository{Db}
}

func (r *ExercicioRepository) FindAll(params url.Values) *[]models.Exercise {
	var exercicios []models.Exercise
	query := r.Db.Table("exercicios").Select([]string{"exercicios.*"})

	if params != nil {
		for key, value := range params {
			if key == "name" {
				for _, value := range value {
					whereClause := "exercicios." + key + " like ?"
					s := []string{"%", value, "%"}
					query.Where(whereClause, strings.Join(s, ""))
				}
			} else {
				whereClause := "exercicios." + key + " = ?"
				query.Where(whereClause, value)
			}
		}
	}

	result := query.Find(&exercicios)
	apiErrors.CheckPanic(result.Error)
	return &exercicios
}

func (r *ExercicioRepository) FindExercisesByMusclePortion(musclePortionId int) {

}

func (r *ExercicioRepository) FindoExercisesByMuscleGroup(muscleGroupId int) {

}
