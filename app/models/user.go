package models

import (
	"time"
)

type User struct {
	Id        int       `json:id        gorm:"primary_key -all"`
	Name      string    `json:name      gorm:"-all"`
	Nickname  string    `json:nickname  gorm:"-all"`
	Email     string    `json:email     gorm:"-all"`
	Password  string    `json:password  gorm:"-all"`
	Cellphone string    `json:cellphone gorm:"-all"`
	CreatedAt time.Time `json:createdAt gorm:"-all"`
	UpdatedAt time.Time `json:updatedAt gorm:"-all"`
}
