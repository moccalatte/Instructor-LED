package usecase_mock

import (
	"final-project-kelompok-1/model"
	"final-project-kelompok-1/model/dto"

	"github.com/stretchr/testify/mock"
)

type MockStudentUseCase struct {
	mock.Mock
}

func (m *MockStudentUseCase) AddStudent(payload dto.StudentRequestDto) (model.Student, error) {
	args := m.Called(payload)
	return args.Get(0).(model.Student), args.Error(1)
}

func (m *MockStudentUseCase) FindStudentByID(id string) (model.Student, error) {
	args := m.Called(id)
	return args.Get(0).(model.Student), args.Error(1)
}

func (m *MockStudentUseCase) GetAllStudent() ([]model.Student, error) {
	args := m.Called()
	return args.Get(0).([]model.Student), args.Error(1)
}

func (m *MockStudentUseCase) UpdateStudent(payload dto.StudentRequestDto, id string) (model.Student, error) {
	args := m.Called(payload, id)
	return args.Get(0).(model.Student), args.Error(1)
}

func (m *MockStudentUseCase) DeleteStudent(id string) (model.Student, error) {
	args := m.Called(id)
	return args.Get(0).(model.Student), args.Error(1)
}

func (m *MockStudentUseCase) FindByEmailPassword(email string, password string) (model.Student, error) {
	args := m.Called(email, password)
	return args.Get(0).(model.Student), args.Error(1)
}
