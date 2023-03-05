package models

import "time"

type UsersPermissions struct {
	Id           int64     `json:"id" gorm:"primary_key"`
	UserId       int       `json:"user_id" `
	PermissionId int       `json:"permission_id"`
	CreatedAt    time.Time `gorm:"column:createdAt"`
	UpdatedAt    time.Time `gorm:"column:updatedAt"`
}
