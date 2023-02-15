package models

import "time"

type Personal struct {
	Id        int       `json:"id"        gorm:"primary_key -all"`
	Name      string    `json:"name"      gorm:"-all"`
	Nickname  string    `json:"nickname"  gorm:"-all"`
	Email     string    `json:"email"     gorm:"-all"`
	Password  string    `json:"password"  gorm:"-all"`
	Cellphone string    `json:"cellphone" gorm:"-all"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:createdAt;autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updatedAt;autoUpdateTime"`
}

func (u Personal) GetId() int {
	return u.Id
}

func (u Personal) GetName() string {
	return u.Name
}

func (u Personal) GetNickname() string {
	return u.Nickname
}

func (u Personal) GetEmail() string {
	return u.Email
}
func (u Personal) GetPassword() string {
	return u.Password
}
func (u Personal) GetCellphone() string {
	return u.Cellphone
}
func (u Personal) GetCreatedAt() time.Time {
	return u.CreatedAt
}
func (u Personal) GetUpdatedAt() time.Time {
	return u.UpdatedAt
}
