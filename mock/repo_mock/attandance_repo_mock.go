package repo_mock

import (
	"final-project-kelompok-1/model"

	"github.com/stretchr/testify/mock"
)

type AttendanceRepoMock struct {
	mock.Mock
}

func (a *AttendanceRepoMock) Create(payload model.Attendance) (model.Attendance, error) {
	args := a.Called(payload)
	return args.Get(0).(model.Attendance), args.Error(1)
}

func (a *AttendanceRepoMock) GetById(id string) (model.Attendance, error) {
	args := a.Called(id)
	return args.Get(0).(model.Attendance), args.Error(1)
}

func (a *AttendanceRepoMock) GetAll() ([]model.Attendance, error) {
	args := a.Called()
	return args.Get(0).([]model.Attendance), args.Error(1)
}

func (a *AttendanceRepoMock) GetBySessionId(id string) (model.Attendance, error) {
	args := a.Called(id)
	return args.Get(0).(model.Attendance), args.Error(1)
}

func (a *AttendanceRepoMock) Update(payload model.Attendance, id string) (model.Attendance, error) {
	args := a.Called(payload, id)
	return args.Get(0).(model.Attendance), args.Error(1)
}

func (a *AttendanceRepoMock) Delete(id string) (model.Attendance, error) {
	args := a.Called(id)
	return args.Get(0).(model.Attendance), args.Error(1)
}
