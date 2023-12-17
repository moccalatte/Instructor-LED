package usecase_mock

import (
	"final-project-kelompok-1/model"
	"final-project-kelompok-1/model/dto"

	"github.com/stretchr/testify/mock"
)

type MockUserUseCase struct {
	mock.Mock
}

func (m *MockUserUseCase) AddUser(payload dto.UserRequestDto) (model.Users, error) {
	args := m.Called(payload)
	return args.Get(0).(model.Users), args.Error(1)
}

func (m *MockUserUseCase) FindUserByID(id string) (model.Users, error) {
	args := m.Called(id)
	return args.Get(0).(model.Users), args.Error(1)
}

func (m *MockUserUseCase) GetAllUser() ([]model.Users, error) {
	args := m.Called()
	return args.Get(0).([]model.Users), args.Error(1)
}

func (m *MockUserUseCase) UpdateUser(payload dto.UserRequestDto, id string) (model.Users, error) {
	args := m.Called(payload, id)
	return args.Get(0).(model.Users), args.Error(1)
}

func (m *MockUserUseCase) DeleteUser(id string) (model.Users, error) {
	args := m.Called(id)
	return args.Get(0).(model.Users), args.Error(1)
}

func (m *MockUserUseCase) RegisterNewUser(payload model.Users) (model.Users, error) {
	args := m.Called(payload)
	return args.Get(0).(model.Users), args.Error(1)
}

func (m *MockUserUseCase) FindByUsernamePassword(email string, password string) (model.Users, error) {
	args := m.Called(email, password)
	return args.Get(0).(model.Users), args.Error(1)
}
