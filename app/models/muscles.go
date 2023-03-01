package models

import (
	"time"
)

type Muscle struct {
	Id        int              `json:"id,omitempty"        gorm:"primaryKey;autoIncrement;"`
	Name      string           `json:"name,omitempty"`
	Image     string           `json:"image,omitempty"`
	Portions  *[]MusclePortion `json:"portions,omitempty" gorm:"foreignKey:MuscleId;references:Id;"`
	CreatedAt time.Time        `json:"createdAt,omitempty" gorm:"column:createdAt;autoCreateTime"`
	UpdatedAt time.Time        `json:"updatedAt,omitempty" gorm:"column:updatedAt;autoUpdateTime"`
}

type MusclePortion struct {
	Id        int        `json:"id,omitempty"        gorm:"primaryKey;autoIncrement;"`
	Name      string     `json:"name,omitempty"`
	Image     string     `json:"image,omitempty"`
	MuscleId  int        `json:"muscle_group_id,omitempty" gorm:"column:muscleGroup_id;"`
	Muscle    *Muscle    `json:"group,omitempty"`
	Exercises []Exercise `json:"exercises,omitempty" gorm:"many2many:exercise_muscle_portions;foreignKey:Id;joinForeignKey:MusclePortionId;References:Id"`
	CreatedAt time.Time  `json:"createdAt,omitempty" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty" gorm:"column:updated_at;autoUpdateTime"`
}

type ExerciseMuscle struct {
}

type Tabler interface {
	TableName() string
}

func (MusclePortion) TableName() string {
	return "muscle_portion"
}

func (Muscle) TableName() string {
	return "muscleGroup"
}
