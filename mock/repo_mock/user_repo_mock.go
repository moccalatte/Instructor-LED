package repo_mock

import (
	"final-project-kelompok-1/model"

	"github.com/stretchr/testify/mock"
)

type UserRepoMock struct {
	mock.Mock
}

func (u *UserRepoMock) Create(payload model.Users) (model.Users, error) {
	args := u.Called(payload)
	return args.Get(0).(model.Users), args.Error(1)
}

func (u *UserRepoMock) GetById(id string) (model.Users, error) {
	args := u.Called(id)
	return args.Get(0).(model.Users), args.Error(1)
}

func (u *UserRepoMock) GetAll() ([]model.Users, error) {
	args := u.Called()
	return args.Get(0).([]model.Users), args.Error(1)
}
func (u *UserRepoMock) Update(payload model.Users, id string) (model.Users, error) {
	args := u.Called(payload, id)
	return args.Get(0).(model.Users), args.Error(1)
}

func (u *UserRepoMock) Delete(id string) (model.Users, error) {
	args := u.Called(id)
	return args.Get(0).(model.Users), args.Error(1)
}

func (u *UserRepoMock) GetByUsername(username string) (model.Users, error) {
	args := u.Called(username)
	return args.Get(0).(model.Users), args.Error(1)

}
