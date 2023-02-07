package repository

import (
	"api/app/models"

	"gorm.io/gorm"
)

type PermissionRepository struct {
	Db *gorm.DB
}

var PermissionTypes = map[string]int{
	"Aluno":    1,
	"Personal": 2,
	"Admin":    3,
}

func NewPermissionRepository(Db *gorm.DB) PermissionRepository {
	return PermissionRepository{
		Db: Db,
	}
}

func (r *PermissionRepository) CreateAluno(userId int) {
	permissionAluno := models.UsersPermissions{
		UserId:       userId,
		PermissionId: PermissionTypes["Aluno"],
	}
	r.Db.Create(&permissionAluno)
}

func (r *PermissionRepository) CreatePersonal(userId int) error {
	permission := models.UsersPermissions{
		PermissionId: PermissionTypes["Personal"],
		UserId:       userId,
	}
	r.Db.Create(&permission)
	return nil
}
