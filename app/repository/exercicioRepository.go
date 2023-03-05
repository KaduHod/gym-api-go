package repository

import (
	apiErrors "api/app/helpers/errors"
	"api/app/helpers/integers"
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

func (r *ExercicioRepository) FindById(id int) *models.Exercise {
	var exercise models.Exercise
	r.Db.Where("id = ?", integers.ToString(id)).Find(&exercise)
	return &exercise
}

func (r *ExercicioRepository) FindByIdWithMuscles(id int, muscle bool, muscleRole bool) *models.Exercise {
	var exercise models.Exercise

	idString := integers.ToString(id)

	query := r.Db.Model(exercise)
	if muscleRole {
		query.Preload("ExerciseMusclePortion", func(db *gorm.DB) *gorm.DB {
			return db.Preload("MusclePortion")
		})
	}
	if muscle {
		query.Preload("MusclePortions")
	}
	query.Where("id = ?", idString).First(&exercise)
	return &exercise
}

func (r *ExercicioRepository) FindByMuscleId(muscleId int) *[]models.Exercise {
	var exercises *[]models.Exercise
	query := r.Db
	query.Preload("MusclePortions")
	query.Find(&exercises)
	return exercises
}

func (r *ExercicioRepository) FindAll(params url.Values) *[]models.Exercise {
	var exercicios []models.Exercise
	query := r.Db

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
