package service

import (
	"api/app/models"
	"api/app/repository"
)

type UpdateAlunoService struct {
	AlunoRepository *repository.AlunosRepository
	AlunoParams     *models.Aluno
}

func (s *UpdateAlunoService) Main() models.Aluno {
	var queryAluno models.Aluno
	s.AlunoRepository.First(s.AlunoParams.Id, &queryAluno)

	queryAluno.Name = s.AlunoParams.Name
	queryAluno.Nickname = s.AlunoParams.Nickname
	queryAluno.Cellphone = s.AlunoParams.Cellphone
	queryAluno.Email = s.AlunoParams.Email
	queryAluno.Password = s.AlunoParams.Password
	s.AlunoRepository.Update(&queryAluno)

	return queryAluno
}
