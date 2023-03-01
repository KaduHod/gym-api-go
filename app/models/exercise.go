package models

import "time"

type Exercise struct {
	Id                    int                     `json:"id,omitempty"        gorm:"primaryKey;autoIncrement;"`
	Name                  string                  `json:"name,omitempty"`
	Force                 string                  `json:"force,omitempty"`
	Execution             string                  `json:"execution,omitempty"`
	Mechanic              string                  `json:"mechanic,omitempty"`
	MusclePortions        []MusclePortion         `json:"muscles,omitempty" gorm:"many2many:ExerciseMusclePortion;"`
	ExerciseMusclePortion []ExerciseMusclePortion `json:"musclesAndRoles,omitempty"`
	Link                  string                  `json:"link,omitempty"`
	CreatedAt             time.Time               `json:"createdAt,omitempty" gorm:"column:createdAt;autoCreateTime"`
	UpdatedAt             time.Time               `json:"updatedAt,omitempty" gorm:"column:updatedAt;autoUpdateTime"`
}

type ExerciseMusclePortion struct {
	Id              int           `json:"id,omitempty"        gorm:"primaryKey;autoIncrement;"`
	ExerciseId      int           `json:"exercise_id,omitempty"`
	Exercise        Exercise      `json:"exercise,omitempty"`
	MusclePortionId int           `json:"-,omitempty" gorm:"column:muscle_id"`
	MusclePortion   MusclePortion `json:"muscle_portion,omitempty,muscle"`
	Role            string        `json:"role,omitempty,role"`
	CreatedAt       time.Time     `json:"created_at,omitempty" gorm:"column:createdAt;autoCreateTime"`
	UpdatedAt       time.Time     `json:"updated_at,omitempty" gorm:"column:updatedAt;autoUpdateTime"`
}

func (Exercise) TableName() string {
	return "exercicios"
}

func (ExerciseMusclePortion) TableName() string {
	return "exercise_muscle_portions"
}

func (e Exercise) GetId() int {
	return e.Id
}

func (e Exercise) GetName() string {
	return e.Name
}

func (e Exercise) GetForce() string {
	return e.Force
}

func (e Exercise) GetExecution() string {
	return e.Execution
}

func (e Exercise) GetMechanic() string {
	return e.Mechanic
}

func (e Exercise) GetLink() string {
	return e.Link
}

func (e Exercise) GetCreatedAt() time.Time {
	return e.CreatedAt
}
func (e Exercise) GetUpdatedAt() time.Time {
	return e.UpdatedAt
}
