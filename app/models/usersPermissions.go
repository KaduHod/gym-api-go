package models

import "time"

type UsersPermissions struct {
	Id           int64     `json:id            gorm:"primary_key -all"`
	UserId       int       `json:user_id       gorm:"-all"`
	PermissionId int       `json:permission_id gorm:"-all"`
	CreatedAt    time.Time `gorm:"column:createdAt"`
	UpdatedAt    time.Time `gorm:"column:updatedAt"`
}
