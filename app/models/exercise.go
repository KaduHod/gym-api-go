package models

import "time"

type Exercise struct {
	Id             int             `json:"id"        gorm:"primary_key -all"`
	Name           string          `json:"name"`
	Force          string          `json:"force"`
	Execution      string          `json:"execution"`
	Mechanic       string          `json:"mechanic"`
	MusclePortions []MusclePortion `json:"muscles" gorm:"many2many:exercise_musclePortion;ForeignKey:Id;joinForeignKey:exercise_id"`
	Link           string          `json:"link"`
	CreatedAt      time.Time       `json:"createdAt" gorm:"column:createdAt;autoCreateTime"`
	UpdatedAt      time.Time       `json:"updatedAt" gorm:"column:updatedAt;autoUpdateTime"`
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
