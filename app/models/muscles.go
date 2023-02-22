package models

import (
	"time"
)

type Muscle struct {
	Id        int             `json:"id"        gorm:"primary_key -all"`
	Name      string          `json:"name"`
	Image     string          `json:"image"`
	Portions  []MusclePortion `json:"portions" gorm:"foreignKey:MuscleId"`
	CreatedAt time.Time       `json:"createdAt" gorm:"column:createdAt;autoCreateTime"`
	UpdatedAt time.Time       `json:"updatedAt" gorm:"column:updatedAt;autoUpdateTime"`
}

type MusclePortion struct {
	Id        int        `json:"id"        gorm:"primary_key -all"`
	Name      string     `json:"name"`
	Image     string     `json:"image"`
	MuscleId  int        `json:"muscleGroup_id" gorm:"column:muscleGroup_id;"`
	Exercises []Exercise `json:"exercises" gorm:"many2many:exercise_musclePortion;ForeignKey:Id;joinForeignKey:muscle_id;"`
	CreatedAt time.Time  `json:"createdAt" gorm:"column:createdAt;autoCreateTime"`
	UpdatedAt time.Time  `json:"updatedAt" gorm:"column:updatedAt;autoUpdateTime"`
}

type ExerciseMuscle struct {
}

type Tabler interface {
	TableName() string
}

func (MusclePortion) TableName() string {
	return "musclePortion"
}
