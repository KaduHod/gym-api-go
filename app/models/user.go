package models

import (
	"encoding/json"
	"time"
)

type User struct {
	Id         int              `json:"id,omitempty"        gorm:"primary_key -all"`
	Name       string           `json:"name,omitempty"      gorm:"-all"`
	Nickname   string           `json:"nickname,omitempty"  gorm:"-all"`
	Email      string           `json:"email,omitempty"     gorm:"-all"`
	Password   string           `json:"password,omitempty"  gorm:"-all"`
	Cellphone  string           `json:"cellphone,omitempty" gorm:"-all"`
	Permission UsersPermissions `json:"permission"`
	CreatedAt  time.Time        `json:"createdAt,omitempty" gorm:"column:createdAt;autoCreateTime"`
	UpdatedAt  time.Time        `json:"updatedAt,omitempty" gorm:"column:updatedAt;autoUpdateTime"`
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

func JsonToUser[T UserType](jsonString string) *T {
	var user T
	json.Unmarshal([]byte(jsonString), &user)
	return &user
}

func JsonToUsers[T UserType](jsonString string) []*T {
	var users []*T

	if err := json.Unmarshal([]byte(jsonString), &users); err != nil {
		panic(err)
	}

	return users
}
