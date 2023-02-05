package models

import (
	"time"
)

type User struct {
	Id        uint      `json:id        gorm:"-all"`
	Name      string    `json:name      gorm:"-all"`
	Dickname  string    `json:nickname  gorm:"-all"`
	Email     string    `json:email     gorm:"-all"`
	Password  string    `json:password  gorm:"-all"`
	cellphone string    `json:cellphone gorm:"-all"`
	createdAt time.Time `json:createdAt gorm:"-all"`
	updatedAt time.Time `json:updatedAt gorm:"-all"`
}
