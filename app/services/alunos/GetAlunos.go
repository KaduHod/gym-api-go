package service

import (
	"api/app/models"
	"api/app/repository"
	"net/url"
)

type ListAlunosService struct {
	AlunoRepository *repository.AlunosRepository
	Params          url.Values
}

func (s *ListAlunosService) Main() *[]models.User {
	return s.AlunoRepository.FindAll(s.Params)
}
