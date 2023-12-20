package repo_mock

import (
	"final-project-kelompok-1/model"

	"github.com/stretchr/testify/mock"
)

type CourseRepoMock struct {
	mock.Mock
}

func (s *CourseRepoMock) Create(payload model.Course) (model.Course, error) {
	args := s.Called(payload)
	return args.Get(0).(model.Course), args.Error(1)
}

func (s *CourseRepoMock) GetById(id string) (model.Course, error) {
	args := s.Called(id)
	return args.Get(0).(model.Course), args.Error(1)
}

func (u *CourseRepoMock) GetAll() ([]model.Course, error) {
	args := u.Called()
	return args.Get(0).([]model.Course), args.Error(1)
}

func (s *CourseRepoMock) Update(payload model.Course, id string) (model.Course, error) {
	args := s.Called(payload, id)
	return args.Get(0).(model.Course), args.Error(1)
}

func (s *CourseRepoMock) Delete(id string) (model.Course, error) {
	args := s.Called(id)

	return args.Get(0).(model.Course), args.Error(1)
}
