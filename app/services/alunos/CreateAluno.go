package service

import (
	"api/app/models"
	"api/app/repository"
)

type CreateAlunoService struct {
	AlunoRepository      *repository.AlunosRepository
	PermissionRepository *repository.PermissionRepository
	AlunoParams          models.User
}

func (s *CreateAlunoService) Main() (models.User, error) {
	alunoCreated, err := s.AlunoRepository.Create(s.AlunoParams)
	if err != nil {
		return alunoCreated, err
	}
	s.PermissionRepository.CreateAluno(alunoCreated.Id)
	return alunoCreated, nil
}
