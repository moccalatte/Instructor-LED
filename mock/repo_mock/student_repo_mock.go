package repo_mock

import (
	"final-project-kelompok-1/model"

	"github.com/stretchr/testify/mock"
)

type StudentRepoMock struct {
	mock.Mock
}

func (s *StudentRepoMock) Create(payload model.Student) (model.Student, error) {
	args := s.Called(payload)
	return args.Get(0).(model.Student), args.Error(1)
}

func (s *StudentRepoMock) GetById(id string) (model.Student, error) {
	args := s.Called(id)
	return args.Get(0).(model.Student), args.Error(1)
}

func (s *StudentRepoMock) GetAll() ([]model.Student, error) {
	args := s.Called()
	return args.Get(0).([]model.Student), args.Error(1)
}

func (s *StudentRepoMock) GetByStudentEmail(email string) (model.Student, error) {
	args := s.Called(email)
	return args.Get(0).(model.Student), args.Error(1)
}

func (s *StudentRepoMock) Update(payload model.Student, id string) (model.Student, error) {
	args := s.Called(payload, id)
	return args.Get(0).(model.Student), args.Error(1)
}

func (s *StudentRepoMock) Delete(id string) (model.Student, error) {
	args := s.Called(id)
	return args.Get(0).(model.Student), args.Error(1)
}
