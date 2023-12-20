package repo_mock

import (
	"final-project-kelompok-1/model"

	"github.com/stretchr/testify/mock"
)

type QuestionRepoMock struct {
	mock.Mock
}

func (s *QuestionRepoMock) Create(payload model.Question) (model.Question, error) {
	args := s.Called(payload)
	return args.Get(0).(model.Question), args.Error(1)
}

func (s *QuestionRepoMock) GetById(id string) (model.Question, error) {
	args := s.Called(id)
	return args.Get(0).(model.Question), args.Error(1)
}

func (s *QuestionRepoMock) GetImagePath(id string) (string, error) {
	args := s.Called(id)
	return args.Get(0).(string), args.Error(1)
}

func (s *QuestionRepoMock) GetByStudentId(id string) (model.Question, error) {
	args := s.Called(id)
	return args.Get(0).(model.Question), args.Error(1)
}

func (s *QuestionRepoMock) GetAll() ([]model.Question, error) {
	args := s.Called()
	return args.Get(0).([]model.Question), args.Error(1)
}

func (s *QuestionRepoMock) Update(payload model.Question, id string) (model.Question, error) {
	args := s.Called(payload, id)
	return args.Get(0).(model.Question), args.Error(1)
}

func (s *QuestionRepoMock) Answer(payload model.Question, id string) (model.Question, error) {
	args := s.Called(payload, id)
	return args.Get(0).(model.Question), args.Error(1)
}

func (s *QuestionRepoMock) Delete(id string) (model.Question, error) {
	args := s.Called(id)

	return args.Get(0).(model.Question), args.Error(1)
}
