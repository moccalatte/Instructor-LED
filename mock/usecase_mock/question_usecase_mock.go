package usecase_mock

import (
	"final-project-kelompok-1/model"
	"final-project-kelompok-1/model/dto"

	"github.com/stretchr/testify/mock"
)

type MockQuestionUseCase struct {
	mock.Mock
}

func (m *MockQuestionUseCase) AddQuestion(payload dto.QuestionRequestDto) (model.Question, error) {
	args := m.Called(payload)
	return args.Get(0).(model.Question), args.Error(1)
}

func (m *MockQuestionUseCase) FindQuestionById(id string) (model.Question, error) {
	args := m.Called(id)
	return args.Get(0).(model.Question), args.Error(1)
}

func (m *MockQuestionUseCase) FindQuestionByStudentId(id string) (model.Question, error) {
	args := m.Called(id)
	return args.Get(0).(model.Question), args.Error(1)
}

func (m *MockQuestionUseCase) GetAllQuestion() ([]model.Question, error) {
	args := m.Called()
	return args.Get(0).([]model.Question), args.Error(1)
}

func (m *MockQuestionUseCase) Update(payload dto.QuestionRequestDto, id string) (model.Question, error) {
	args := m.Called(payload, id)
	return args.Get(0).(model.Question), args.Error(1)
}

func (m *MockQuestionUseCase) Delete(id string) (model.Question, error) {
	args := m.Called(id)
	return args.Get(0).(model.Question), args.Error(1)
}

func (m *MockQuestionUseCase) Answer(payload dto.QuestionRequestDto, id string) (model.Question, error) {
	args := m.Called(payload, id)
	return args.Get(0).(model.Question), args.Error(1)
}

func (m *MockQuestionUseCase) GetImagePath(questionID string) (string, error) {
	args := m.Called(questionID)
	return args.String(0), args.Error(1)
}
