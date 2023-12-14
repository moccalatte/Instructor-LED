package repo_mock

import (
	"final-project-kelompok-1/model"

	"github.com/stretchr/testify/mock"
)

type CourseRepoMock struct {
	mock.Mock
}

func (c *CourseRepoMock) Create(payload model.Course) (model.Course, error) {
	args := c.Called(payload)
	return args.Get(0).(model.Course), args.Error(1)
}

func (c *CourseRepoMock) GetById(id string) (model.Course, error) {
	args := c.Called(id)
	return args.Get(0).(model.Course), args.Error(1)
}

func (c *CourseRepoMock) Update(payload model.Course, id string) (model.Course, error) {
	args := c.Called(payload, id)
	return args.Get(0).(model.Course), args.Error(1)
}

func (c *CourseRepoMock) Delete(id string) (model.Course, error) {
	args := c.Called(id)
	return args.Get(0).(model.Course), args.Error(1)
}
