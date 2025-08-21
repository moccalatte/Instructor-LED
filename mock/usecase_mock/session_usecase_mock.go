package usecase_mock

import (
	"instructor-led/model"
	"instructor-led/model/dto"

	"github.com/stretchr/testify/mock"
)

type MockSessionUseCase struct {
	mock.Mock
}

func (m *MockSessionUseCase) AddSession(payload dto.SessionRequestDto) (model.Session, error) {
	args := m.Called(payload)
	return args.Get(0).(model.Session), args.Error(1)
}

func (m *MockSessionUseCase) FindSessionById(id string) (model.Session, error) {
	args := m.Called(id)
	return args.Get(0).(model.Session), args.Error(1)
}

func (m *MockSessionUseCase) GetAllSession() ([]model.Session, error) {
	args := m.Called()
	return args.Get(0).([]model.Session), args.Error(1)
}

func (m *MockSessionUseCase) Update(payload dto.SessionRequestDto, id string) (model.Session, error) {
	args := m.Called(payload, id)
	return args.Get(0).(model.Session), args.Error(1)
}

func (m *MockSessionUseCase) UpdateNote(payload dto.SessionRequestDto, id string) (model.Session, error) {
	args := m.Called(payload, id)
	return args.Get(0).(model.Session), args.Error(1)
}

func (m *MockSessionUseCase) Delete(id string) (model.Session, error) {
	args := m.Called(id)
	return args.Get(0).(model.Session), args.Error(1)
}
