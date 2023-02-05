package service

import (
	"api/app/models"
	"api/app/repository"
)

type CreatePersonalService struct {
	PersonalRepository   *repository.PersonalRepository
	PermissionRepository *repository.PermissionRepository
	Personal             *models.Personal
}

func (s *CreatePersonalService) Main() error {
	err := s.createUser()
	if err != nil {
		return err
	}
	return nil
}

func (s *CreatePersonalService) createUser() error {
	err := s.PersonalRepository.Create(s.Personal)
	if err != nil {
		return err
	}
	erro := s.createPermission()
	if erro != nil {
		return erro
	}
	return nil
}

func (s *CreatePersonalService) createPermission() error {
	err := s.PermissionRepository.CreatePersonal(s.Personal.Id)
	return err
}
