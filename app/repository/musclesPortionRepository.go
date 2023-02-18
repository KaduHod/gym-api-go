package repository

import (
	"api/app/models"
	"strings"

	"gorm.io/gorm"
)

type MusclesPortionRepository struct {
	Db *gorm.DB
}

func NewMusclePortionRepository(Db *gorm.DB) MusclesPortionRepository {
	return MusclesPortionRepository{
		Db: Db,
	}
}

func (r *MusclesPortionRepository) FindAll(params map[string][]string) *[]models.MusclePortion {
	var muscles []models.MusclePortion
	query := r.Db.Table("musclePortion").Select("musclePortion.*")

	if params != nil {
		for key, value := range params {
			if key == "name" {
				for _, value := range value {
					whereClause := key + " like ?"
					s := []string{"%", value, "%"}
					query.Where(whereClause, strings.Join(s, ""))
				}
			} else if key == "muscleGroup_id" {
				query.Joins("JOIN muscleGroup on musclePortion.muscleGRoup_id = muscleGroup.id")
			} else {
				query.Where(key, value)
			}
		}
	}

	query.Find(&muscles)
	return &muscles
}
