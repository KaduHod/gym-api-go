package service

import (
	"api/app/models"
	"api/app/repository"
)

type CreateAlunoService struct {
	AlunoRepository      *repository.AlunosRepository
	PermissionRepository *repository.PermissionRepository
	Aluno                *models.User
}

func (s *CreateAlunoService) Main() error {
	err := s.AlunoRepository.Create(s.Aluno)
	if err != nil {
		return err
	}
	s.PermissionRepository.CreateAluno(s.Aluno.Id)
	return nil
}
