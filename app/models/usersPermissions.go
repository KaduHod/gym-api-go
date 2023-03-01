package models

import "time"

type UsersPermissions struct {
	Id           int64     `json:"id,omitempty" gorm:"primary_key"`
	UserId       int       `json:"user_id,omitempty" `
	PermissionId int       `json:"permission_id,omitempty"`
	CreatedAt    time.Time `gorm:"column:createdAt"`
	UpdatedAt    time.Time `gorm:"column:updatedAt"`
}
