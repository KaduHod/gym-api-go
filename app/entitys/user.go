package entitys

import (
	"time"
)

type UserE struct {
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
	UserE | AlunoE | PersonalE
}

type UserType interface {
	UserE | AlunoE | PersonalE
	GetId() int
	GetName() string
	GetNickname() string
	GetEmail() string
	GetPassword() string
	GetCellphone() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}

func (u UserE) GetId() int {
	return u.Id
}

func (u UserE) GetName() string {
	return u.Name
}

func (u UserE) GetNickname() string {
	return u.Nickname
}

func (u UserE) GetEmail() string {
	return u.Email
}
func (u UserE) GetPassword() string {
	return u.Password
}
func (u UserE) GetCellphone() string {
	return u.Cellphone
}
func (u UserE) GetCreatedAt() time.Time {
	return u.CreatedAt
}
func (u UserE) GetUpdatedAt() time.Time {
	return u.UpdatedAt
}
