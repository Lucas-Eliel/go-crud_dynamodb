package services_test

import (
	"testing"

	"crud_dynamodb/internal/models"
	services "crud_dynamodb/internal/services"
	mock_repositories "crud_dynamodb/test/unit/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetUsuarios(t *testing.T) {
	//Criar cenário
	ctrl := gomock.NewController(t)
	mockRepository := mock_repositories.NewMockRepository(ctrl)
	var users []models.User
	users = append(users, models.User{ID: "1", Nome: "João"}, models.User{ID: "2", Nome: "Pedro"})
	mockRepository.EXPECT().FindAll().Return(users)

	//Chamar o metodo a ser testado
	serv := services.NewService(mockRepository)
	resp := serv.GetUsuarios()

	//Realizar as verficações
	assert.NotEmpty(t, resp)
}
