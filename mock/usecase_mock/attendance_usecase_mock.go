package usecase_mock

import (
	"instructor-led/model"
	"instructor-led/model/dto"

	"github.com/stretchr/testify/mock"
)

type MockAttendanceUseCase struct {
	mock.Mock
}

func (m *MockAttendanceUseCase) AddAttendance(payload dto.AttendanceRequestDto) (model.Attendance, error) {
	args := m.Called(payload)
	return args.Get(0).(model.Attendance), args.Error(1)
}

func (m *MockAttendanceUseCase) FindAttendanceByID(id string) (model.Attendance, error) {
	args := m.Called(id)
	return args.Get(0).(model.Attendance), args.Error(1)
}

func (m *MockAttendanceUseCase) GetAllAttendance() ([]model.Attendance, error) {
	args := m.Called()
	return args.Get(0).([]model.Attendance), args.Error(1)
}

func (m *MockAttendanceUseCase) FindAttendanceBySessionId(id string) (model.Attendance, error) {
	args := m.Called(id)
	return args.Get(0).(model.Attendance), args.Error(1)
}

func (m *MockAttendanceUseCase) UpdateAttendance(payload dto.AttendanceRequestDto, id string) (model.Attendance, error) {
	args := m.Called(payload, id)
	return args.Get(0).(model.Attendance), args.Error(1)
}

func (m *MockAttendanceUseCase) Delete(id string) (model.Attendance, error) {
	args := m.Called(id)
	return args.Get(0).(model.Attendance), args.Error(1)
}
