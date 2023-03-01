package models

import "time"

type Aluno struct {
	Id        int       `json:"omitempty,id"        gorm:"primary_key -all"`
	Name      string    `json:"omitempty,name"      gorm:"-all"`
	Nickname  string    `json:"omitempty,nickname"  gorm:"-all"`
	Email     string    `json:"omitempty,email"     gorm:"-all"`
	Password  string    `json:"omitempty,password"  gorm:"-all"`
	Cellphone string    `json:"omitempty,cellphone" gorm:"-all"`
	CreatedAt time.Time `json:"omitempty,createdAt" gorm:"column:createdAt;autoCreateTime"`
	UpdatedAt time.Time `json:"omitempty,updatedAt" gorm:"column:updatedAt;autoUpdateTime"`
}

func (u Aluno) GetId() int {
	return u.Id
}

func (u Aluno) GetName() string {
	return u.Name
}

func (u Aluno) GetNickname() string {
	return u.Nickname
}

func (u Aluno) GetEmail() string {
	return u.Email
}
func (u Aluno) GetPassword() string {
	return u.Password
}
func (u Aluno) GetCellphone() string {
	return u.Cellphone
}
func (u Aluno) GetCreatedAt() time.Time {
	return u.CreatedAt
}
func (u Aluno) GetUpdatedAt() time.Time {
	return u.UpdatedAt
}
