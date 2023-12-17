// mock/controller_mock/attendance_controller_mock.go

package controller_mock

import (
	"final-project-kelompok-1/controller"
	"final-project-kelompok-1/model/dto"
	"github.com/stretchr/testify/mock"
)

// AttendanceControllerMock is a mock implementation of the AttendanceController interface for testing purposes.
type AttendanceControllerMock struct {
	mock.Mock
}

// CreateHandler is a mock implementation for the CreateHandler method.
func (m *AttendanceControllerMock) CreateHandler(ctx *dto.AttendanceRequestDto) (interface{}, error) {
	args := m.Called(ctx)
	return args.Get(0), args.Error(1)
}

// GetHandlerByID is a mock implementation for the GetHandlerByID method.
func (m *AttendanceControllerMock) GetHandlerByID(attendanceID string) (interface{}, error) {
	args := m.Called(attendanceID)
	return args.Get(0), args.Error(1)
}

// GetHandlerAll is a mock implementation for the GetHandlerAll method.
func (m *AttendanceControllerMock) GetHandlerAll() (interface{}, error) {
	args := m.Called()
	return args.Get(0), args.Error(1)
}

// UpdateHandler is a mock implementation for the UpdateHandler method.
func (m *AttendanceControllerMock) UpdateHandler(ctx *dto.AttendanceRequestDto, attendanceID string) (interface{}, error) {
	args := m.Called(ctx, attendanceID)
	return args.Get(0), args.Error(1)
}

// DeleteHandler is a mock implementation for the DeleteHandler method.
func (m *AttendanceControllerMock) DeleteHandler(attendanceID string) (interface{}, error) {
	args := m.Called(attendanceID)
	return args.Get(0), args.Error(1)
}
