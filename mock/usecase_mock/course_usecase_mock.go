package usecase_mock

import (
	"instructor-led/model"
	"instructor-led/model/dto"

	"github.com/stretchr/testify/mock"
)

type MockCourseUseCase struct {
	mock.Mock
}

func (m *MockCourseUseCase) AddCourse(payload dto.CourseRequestDto) (model.Course, error) {
	args := m.Called(payload)
	return args.Get(0).(model.Course), args.Error(1)
}

func (m *MockCourseUseCase) FindCourseByID(id string) (model.Course, error) {
	args := m.Called(id)
	return args.Get(0).(model.Course), args.Error(1)
}

func (m *MockCourseUseCase) GetAllCourse() ([]model.Course, error) {
	args := m.Called()
	return args.Get(0).([]model.Course), args.Error(1)
}

func (m *MockCourseUseCase) UpdateCourse(payload dto.CourseRequestDto, id string) (model.Course, error) {
	args := m.Called(payload, id)
	return args.Get(0).(model.Course), args.Error(1)
}

func (m *MockCourseUseCase) DeleteCourse(id string) (model.Course, error) {
	args := m.Called(id)
	return args.Get(0).(model.Course), args.Error(1)
}
