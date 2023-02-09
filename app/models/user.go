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
	CreatedAt time.Time `gorm:"column:createdAt;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updatedAt;autoUpdateTime"`
}
