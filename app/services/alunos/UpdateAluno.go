package service

import (
	"api/app/models"
	"api/app/repository"
)

type UpdateAlunoService struct {
	AlunoRepository *repository.AlunosRepository
	AlunoParams     *models.Aluno
}

func (s *UpdateAlunoService) Main() error {
	var queryAluno models.Aluno
	error := s.AlunoRepository.First(s.AlunoParams.Id, &queryAluno)

	if error != nil {
		return error
	}

	if s.AlunoParams.Cellphone == "" {
		s.AlunoParams.Cellphone = queryAluno.Cellphone
	}
	if s.AlunoParams.Email == "" {
		s.AlunoParams.Email = queryAluno.Email
	}
	if s.AlunoParams.Name == "" {
		s.AlunoParams.Name = queryAluno.Name
	}
	if s.AlunoParams.Nickname == "" {
		s.AlunoParams.Nickname = queryAluno.Nickname
	}
	if s.AlunoParams.Password == "" {
		s.AlunoParams.Password = queryAluno.Password
	}

	s.AlunoRepository.Update(s.AlunoParams)

	return nil
}
