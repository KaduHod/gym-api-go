package models

import "time"

type Personal struct {
	Id        int       `json:id        gorm:"primary_key -all"`
	Name      string    `json:name      gorm:"-all"`
	Nickname  string    `json:nickname  gorm:"-all"`
	Email     string    `json:email     gorm:"-all"`
	Password  string    `json:password  gorm:"-all"`
	cellphone string    `json:cellphone gorm:"-all"`
	createdAt time.Time `json:createdAt gorm:"-all"`
	updatedAt time.Time `json:updatedAt gorm:"-all"`
}
