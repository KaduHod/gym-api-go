package service

import (
	"api/app/models"
	"api/app/repository"
)

type UpdateAlunoService struct {
	AlunoRepository      *repository.AlunosRepository
	PermissionRepository *repository.PermissionRepository
	AlunoParams          *models.Aluno
}

func (s *UpdateAlunoService) Main() {

}
