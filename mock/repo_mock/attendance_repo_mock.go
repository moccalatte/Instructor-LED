package repo_mock

// import (
// 	"final-project-kelompok-1/model"

// 	"github.com/stretchr/testify/mock"
// )

// type AttendanceRepoMock struct{
// 	mock.Mock
// }

// func (s *AttendanceRepoMock) Create(payload model.Attendance) (model.Attendance, error) {
// 	args := s.Called(payload)
// 	return args.Get(0).(model.Attendance), args.Error(1)
// }

// func (s *AttendanceRepoMock) GetById(id string) (model.Attendance, error) {
// 	args := s.Called(id)
// 	return args.Get(0).(model.Attendance), args.Error(1)
// }

// func (s *AttendanceRepoMock) Update(payload model.Attendance, id string) (model.Attendance, error) {
// 	args := s.Called(payload, id)
// 	return args.Get(0).(model.Attendance), args.Error(1)
// }

// func (s *AttendanceRepoMock) Delete(id string) (model.Attendance, error) {
// 	args := s.Called(id)
// 	return args.Get(0).(model.Attendance), args.Error(1)
// }
