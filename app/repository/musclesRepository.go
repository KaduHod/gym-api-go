package repository

import (
	"api/app/models"
	"strings"

	"gorm.io/gorm"
)

type MusclesRepository struct {
	Db *gorm.DB
}

func NewMuscelRepository(Db *gorm.DB) MusclesRepository {
	return MusclesRepository{
		Db: Db,
	}
}

func (r *MusclesRepository) FindAll(params map[string][]string) *[]models.Muscle {
	var muscles []models.Muscle
	query := r.Db.Table("muscles")

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
