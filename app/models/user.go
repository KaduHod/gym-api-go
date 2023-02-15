package models

import (
	"time"
)

type User struct {
	Id        int       `json:"id"        gorm:"primary_key -all"`
	Name      string    `json:"name"      gorm:"-all"`
	Nickname  string    `json:"nickname"  gorm:"-all"`
	Email     string    `json:"email"     gorm:"-all"`
	Password  string    `json:"password"  gorm:"-all"`
	Cellphone string    `json:"cellphone" gorm:"-all"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:createdAt;autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updatedAt;autoUpdateTime"`
}

type UserTypes interface {
	User | Aluno | Personal
}

type UserType interface {
	User | Aluno | Personal
	GetId() int
	GetName() string
	GetNickname() string
	GetEmail() string
	GetPassword() string
	GetCellphone() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}

func (u User) GetId() int {
	return u.Id
}

func (u User) GetName() string {
	return u.Name
}

func (u User) GetNickname() string {
	return u.Nickname
}

func (u User) GetEmail() string {
	return u.Email
}
func (u User) GetPassword() string {
	return u.Password
}
func (u User) GetCellphone() string {
	return u.Cellphone
}
func (u User) GetCreatedAt() time.Time {
	return u.CreatedAt
}
func (u User) GetUpdatedAt() time.Time {
	return u.UpdatedAt
}
