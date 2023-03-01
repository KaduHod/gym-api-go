package models

import "time"

type UsersPermissions struct {
	Id           int64     `json:"omitempty,id          " gorm:"primary_key -all"`
	UserId       int       `json:"omitempty,user_id     " gorm:"-all"`
	PermissionId int       `json:"omitempty,permission_id" gorm:"-all"`
	CreatedAt    time.Time `gorm:"column:createdAt"`
	UpdatedAt    time.Time `gorm:"column:updatedAt"`
}
