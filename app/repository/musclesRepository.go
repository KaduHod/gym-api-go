package repository

import (
	"api/app/models"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type MusclesGroupsRepository struct {
	Db *gorm.DB
}

func NewMuscelRepository(Db *gorm.DB) MusclesGroupsRepository {
	return MusclesGroupsRepository{
		Db: Db,
	}
}

func (r *MusclesGroupsRepository) FindAll(params map[string][]string) *[]models.Muscle {
	var muscles []models.Muscle
	query := r.Db.Table("muscleGroup")

	if params != nil {
		for key, value := range params {
			if key == "name" {
				for _, value := range value {
					whereClause := key + " like ?"
					s := []string{"%", value, "%"}
					query.Where(whereClause, strings.Join(s, ""))
				}
			} else {
				query.Where(key, value)
			}
		}
	}

	query.Find(&muscles)
	return &muscles
}

func (r *MusclesGroupsRepository) First(groupId int) *models.Muscle {
	var muscle *models.Muscle
	query := r.Db.Table("muscleGroup").Where("id = ?", groupId)
	query.Find(&muscle)
	return muscle
}

func (r *MusclesGroupsRepository) FindFirstWithPortions(groupId int) *models.Muscle {
	var muscle models.Muscle
	fmt.Println(groupId)
	muscle.Id = groupId
	query := r.Db.Preload("Portions").Table("muscleGroup")
	query.Find(&muscle)
	return &muscle
}
