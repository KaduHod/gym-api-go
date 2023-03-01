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
func (r *MusclesPortionRepository) FindWithExercises(portionId int) *models.MusclePortion {
	var muscle models.MusclePortion
	muscle.Id = portionId
	r.Db.Preload("Muscle").Preload("Exercises").Find(&muscle)
	return &muscle
}

func (r *MusclesPortionRepository) All() *[]models.MusclePortion {
	var muscles []models.MusclePortion
	r.Db.Find(&muscles)
	return &muscles
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
				query.Joins("JOIN muscleGroup on musclePortion.muscleGroup_id = muscleGroup.id")
			} else {
				query.Where(key, value)
			}
		}
	}

	query.Find(&muscles)
	return &muscles
}

func (r *MusclesPortionRepository) FindByGroupId(groupId int) *[]models.MusclePortion {
	var muscles []models.MusclePortion
	query := r.Db.Table("musclePortion").Select("musclePortion.*").Where("muscleGroup_id = ?", groupId)

	query.Find(&muscles)
	return &muscles
}
