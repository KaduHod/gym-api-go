package models

import "time"

type UsersPermissions struct {
	Id           int64     `json:id            gorm:"primary_key -all"`
	UserId       int64     `json:user_id       gorm:"-all"`
	PermissionId int64     `json:permission_id gorm:"-all"`
	CreatedAt    time.Time `json:created_at    gorm:"-all"`
	UpdatedAt    time.Time `json:updated_at    gorm:"-all"`
}
