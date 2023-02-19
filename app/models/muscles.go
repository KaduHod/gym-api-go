package models

import "time"

type Muscle struct {
	Id        int              `json:"id"        gorm:"primary_key -all"`
	Name      string           `json:"name"`
	Image     string           `json:"image"`
	Portions  *[]MusclePortion `json:"portions" gorm:"foreignKey:MuscleGroupId"`
	CreatedAt time.Time        `json:"createdAt" gorm:"column:createdAt;autoCreateTime"`
	UpdatedAt time.Time        `json:"updatedAt" gorm:"column:updatedAt;autoUpdateTime"`
}

type MusclePortion struct {
	Id            int       `json:"id"        gorm:"primary_key -all"`
	Name          string    `json:"name"`
	Image         string    `json:"image"`
	MuscleGroupId int       `json:"muscleGroup_id" gorm:"column:muscleGroup_id"`
	CreatedAt     time.Time `json:"createdAt" gorm:"column:createdAt;autoCreateTime"`
	UpdatedAt     time.Time `json:"updatedAt" gorm:"column:updatedAt;autoUpdateTime"`
}
