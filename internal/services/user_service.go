package services

import (
	"crud_dynamodb/internal/models"
	"crud_dynamodb/internal/repositories"
)

// Interface
type Service interface {
	GetUsuarios() []models.User
	GetUsuario(id string) models.User
	PostUsuario(user models.User)
	DeleteUsuario(id string)
}

// Object
type ServiceImpl struct {
	repository repositories.Repository
}

func NewService(repository repositories.Repository) Service {
	return ServiceImpl{repository}
}

func (s ServiceImpl) GetUsuarios() []models.User {
	return s.repository.FindAll()
}

func (s ServiceImpl) GetUsuario(id string) models.User {
	return s.repository.FindById(id)
}

func (s ServiceImpl) PostUsuario(user models.User) {
	s.repository.Create(user)
}

func (s ServiceImpl) DeleteUsuario(id string) {
	s.repository.Delete(id)
}
