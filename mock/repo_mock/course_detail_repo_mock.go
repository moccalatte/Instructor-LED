package repo_mock

import (
	"final-project-kelompok-1/model"

	"github.com/stretchr/testify/mock"
)

type CourseDetailRepoMock struct {
	mock.Mock
}

func (c *CourseDetailRepoMock) Create(payload model.CourseDetail) (model.CourseDetail, error) {
	args := c.Called(payload)
	return args.Get(0).(model.CourseDetail), args.Error(1)
}

func (c *CourseDetailRepoMock) GetById(id string) (model.CourseDetail, error) {
	args := c.Called(id)
	return args.Get(0).(model.CourseDetail), args.Error(1)
}

func (c *CourseDetailRepoMock) Update(payload model.CourseDetail, id string) (model.CourseDetail, error) {
	args := c.Called(payload, id)
	return args.Get(0).(model.CourseDetail), args.Error(1)
}

func (c *CourseDetailRepoMock) Delete(id string) (model.CourseDetail, error) {
	args := c.Called(id)
	return args.Get(0).(model.CourseDetail), args.Error(1)
}
