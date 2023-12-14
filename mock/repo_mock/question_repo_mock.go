package repo_mock

import (
	"final-project-kelompok-1/model"

	"github.com/stretchr/testify/mock"
)

type QuestionRepoMock struct {
	mock.Mock
}

func (q *QuestionRepoMock) Create(payload model.Question) (model.Question, error) {
	args := q.Called(payload)
	return args.Get(0).(model.Question), args.Error(1)
}

func (q *QuestionRepoMock) GetById(id string) (model.Question, error) {
	args := q.Called(id)
	return args.Get(0).(model.Question), args.Error(1)
}

func (q *QuestionRepoMock) GetByStudentId(id string) (model.Question, error) {
	args := q.Called(id)
	return args.Get(0).(model.Question), args.Error(1)
}

func (q *QuestionRepoMock) Update(payload model.Question, id string) (model.Question, error) {
	args := q.Called(payload, id)
	return args.Get(0).(model.Question), args.Error(1)
}

func (q *QuestionRepoMock) Delete(id string) (model.Question, error) {
	args := q.Called(id)
	return args.Get(0).(model.Question), args.Error(1)
}

func (q *QuestionRepoMock) Answer(payload model.Question, id string) (model.Question, error) {
	args := q.Called(payload, id)
	return args.Get(0).(model.Question), args.Error(1)
}
