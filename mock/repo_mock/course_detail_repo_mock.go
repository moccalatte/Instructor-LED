package repo_mock

import (
	"final-project-kelompok-1/model"

	"github.com/stretchr/testify/mock"
)

type CourseDetailRepoMock struct {
	mock.Mock
}

func (s *CourseDetailRepoMock) Create(payload model.CourseDetail) (model.CourseDetail, error) {
	args := s.Called(payload)
	return args.Get(0).(model.CourseDetail), args.Error(1)
}

func (s *CourseDetailRepoMock) GetById(id string) (model.CourseDetail, error) {
	args := s.Called(id)
	return args.Get(0).(model.CourseDetail), args.Error(1)
}

func (s *CourseDetailRepoMock) Update(payload model.CourseDetail, id string) (model.CourseDetail, error) {
	args := s.Called(payload, id)
	return args.Get(0).(model.CourseDetail), args.Error(1)
}

func (s *CourseDetailRepoMock) Delete(id string) (model.CourseDetail, error) {
	args := s.Called(id)

	return args.Get(0).(model.CourseDetail), args.Error(1)
}
