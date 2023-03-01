package models

import "time"

type Exercise struct {
	Id                    int                     `json:"omitempty,id"        gorm:"primaryKey;autoIncrement;"`
	Name                  string                  `json:"omitempty,name"`
	Force                 string                  `json:"omitempty,force"`
	Execution             string                  `json:"omitempty,execution"`
	Mechanic              string                  `json:"omitempty,mechanic"`
	MusclePortions        []MusclePortion         `json:"omitempty,muscles" gorm:"many2many:exercise_muscle_portions;foreignKey:Id;joinForeignKey:ExerciseId;References:Id;"`
	ExerciseMusclePortion []ExerciseMusclePortion `json:"omitempty,musclesAndRoles"`
	Link                  string                  `json:"omitempty,link"`
	CreatedAt             time.Time               `json:"omitempty,createdAt" gorm:"column:createdAt;autoCreateTime"`
	UpdatedAt             time.Time               `json:"omitempty,updatedAt" gorm:"column:updatedAt;autoUpdateTime"`
}

type ExerciseMusclePortion struct {
	Id              int           `json:"omitempty,-"        gorm:"primaryKey;autoIncrement;"`
	ExerciseId      int           `json:"omitempty,-"`
	Exercise        Exercise      `json:"omitempty,-"`
	MusclePortionId int           `json:"omitempty,-" gorm:"column:muscle_id"`
	MusclePortion   MusclePortion `json:"omitempty,muscle"`
	Role            string        `json:"omitempty,role"`
	CreatedAt       time.Time     `json:"omitempty,-" gorm:"column:createdAt;autoCreateTime"`
	UpdatedAt       time.Time     `json:"omitempty,-" gorm:"column:updatedAt;autoUpdateTime"`
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
