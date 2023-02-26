package models

import "time"

type Exercise struct {
	Id                    int                     `json:"id"        gorm:"primaryKey;autoIncrement;"`
	Name                  string                  `json:"name"`
	Force                 string                  `json:"force"`
	Execution             string                  `json:"execution"`
	Mechanic              string                  `json:"mechanic"`
	MusclePortions        []MusclePortion         `json:"muscles" gorm:"many2many:exercise_muscle_portions;foreignKey:Id;joinForeignKey:ExerciseId;References:Id;"`
	ExerciseMusclePortion []ExerciseMusclePortion `json:"musclesAndRoles"`
	Link                  string                  `json:"link"`
	CreatedAt             time.Time               `json:"createdAt" gorm:"column:createdAt;autoCreateTime"`
	UpdatedAt             time.Time               `json:"updatedAt" gorm:"column:updatedAt;autoUpdateTime"`
}

type ExerciseMusclePortion struct {
	Id              int           `json:"-"        gorm:"primaryKey;autoIncrement;"`
	ExerciseId      int           `json:"-"`
	Exercise        Exercise      `json:"-"`
	MusclePortionId int           `json:"-" gorm:"column:muscle_id"`
	MusclePortion   MusclePortion `json:"muscle"`
	Role            string        `json:"role"`
	CreatedAt       time.Time     `json:"-" gorm:"column:createdAt;autoCreateTime"`
	UpdatedAt       time.Time     `json:"-" gorm:"column:updatedAt;autoUpdateTime"`
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
