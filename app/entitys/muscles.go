package entitys

import (
	"time"
)

type MuscleE struct {
	Id        int              `json:"id"        gorm:"primaryKey;autoIncrement;"`
	Name      string           `json:"name"`
	Image     string           `json:"image"`
	Portions  []MusclePortionE `json:"portions" gorm:"foreignKey:MuscleId;references:Id"`
	CreatedAt time.Time        `json:"createdAt" gorm:"column:createdAt;autoCreateTime"`
	UpdatedAt time.Time        `json:"updatedAt" gorm:"column:updatedAt;autoUpdateTime"`
}

type MusclePortionE struct {
	Id        int        `json:"id"        gorm:"primaryKey;autoIncrement;"`
	Name      string     `json:"name"`
	Image     string     `json:"image"`
	MuscleId  int        `json:"muscle_group_id" gorm:"column:muscleGroup_id;"`
	Muscle    MuscleE    `json:"group;"`
	Exercises []Exercise `json:"exercises" gorm:"many2many:exercise_muscle_portions;foreignKey:Id;joinForeignKey:MusclePortionId;References:Id"`
	CreatedAt time.Time  `json:"createdAt" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time  `json:"updatedAt" gorm:"column:updated_at;autoUpdateTime"`
}
