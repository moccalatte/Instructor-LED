package repo_mock

import (
	"final-project-kelompok-1/model"

	"github.com/stretchr/testify/mock"
)

type SessionRepoMock struct {
	mock.Mock
}

func (s *SessionRepoMock) Create(payload model.Session) (model.Session, error) {
	args := s.Called(payload)
	return args.Get(0).(model.Session), args.Error(1)
}

func (s *SessionRepoMock) GetById(id string) (model.Session, error) {
	args := s.Called(id)
	return args.Get(0).(model.Session), args.Error(1)
}

func (s *SessionRepoMock) GetAllSession() ([]model.Session, error) {
	args := s.Called()
	return args.Get(0).([]model.Session), args.Error(1)
}

func (s *SessionRepoMock) Update(payload model.Session, id string) (model.Session, error) {
	args := s.Called(payload, id)
	return args.Get(0).(model.Session), args.Error(1)
}

func (s *SessionRepoMock) UpdateNote(payload model.Session, id string) (model.Session, error) {
	args := s.Called(payload, id)
	return args.Get(0).(model.Session), args.Error(1)
}

func (s *SessionRepoMock) Delete(id string) (model.Session, error) {
	args := s.Called(id)
	return args.Get(0).(model.Session), args.Error(1)
}
