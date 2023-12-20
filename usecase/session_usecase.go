package usecase

import (
	"final-project-kelompok-1/model"
	"final-project-kelompok-1/model/dto"
	"final-project-kelompok-1/repository"
	"fmt"
)

type SessionUseCase interface {
	AddSession(payload dto.SessionRequestDto) (model.Session, error)
	FindSessionById(id string) (model.Session, error)
	GetAllSession() ([]model.Session, error)
	Update(payload dto.SessionRequestDto, id string) (model.Session, error)
	UpdateNote(payload dto.SessionRequestDto, id string) (model.Session, error)
	Delete(id string) (model.Session, error)
}

type sessionUseCase struct {
	repo repository.SessionRepository
}

func (s *sessionUseCase) AddSession(payload dto.SessionRequestDto) (model.Session, error) {
	newSession := model.Session{
		Title:       payload.Title,
		Description: payload.Description,
		SessionDate: payload.SessionDate,
		SessionTime: payload.SessionTime,
		SessionLink: payload.SessionLink,
		TrainerID:   payload.TrainerID,
		Note:        payload.Note,
	}

	createsSession, err := s.repo.Create(newSession)

	if err != nil {
		return model.Session{}, fmt.Errorf("failed to save data : %s", err.Error())
	}
	return createsSession, nil
}

func (s *sessionUseCase) FindSessionById(id string) (model.Session, error) {
	session, err := s.repo.GetById(id)

	if err != nil {
		return model.Session{}, fmt.Errorf("failed to get data session by id : %s", err.Error())
	}
	return session, nil
}

func (s *sessionUseCase) GetAllSession() ([]model.Session, error) {
	var sliceSession []model.Session
	sessionData, err := s.repo.GetAllSession()
	if err != nil {
		return nil, fmt.Errorf("failed to find all data : %s", err.Error())
	}
	return append(sliceSession, sessionData...), nil
}

func (s *sessionUseCase) Update(payload dto.SessionRequestDto, id string) (model.Session, error) {
	sessions := model.Session{
		Title:       payload.Title,
		Description: payload.Description,
		SessionDate: payload.SessionDate,
		SessionTime: payload.SessionTime,
		SessionLink: payload.SessionLink,
		TrainerID:   payload.TrainerID,
		Note:        payload.Note,
	}

	session, err := s.repo.Update(sessions, id)

	if err != nil {
		return model.Session{}, fmt.Errorf("failed to Update Session : %s", err.Error())
	}

	return session, nil
}

func (s *sessionUseCase) UpdateNote(payload dto.SessionRequestDto, id string) (model.Session, error) {
	sessions := model.Session{
		Note: payload.Note,
	}

	session, err := s.repo.UpdateNote(sessions, id)

	if err != nil {
		return model.Session{}, fmt.Errorf("failed to Update Session : %s", err.Error())
	}

	return session, nil
}

func (s *sessionUseCase) Delete(id string) (model.Session, error) {
	session, err := s.repo.Delete(id)

	if err != nil {
		return model.Session{}, fmt.Errorf("failed to delete data : %s", err.Error())
	}

	return session, nil
}

func NewSession(repo repository.SessionRepository) SessionUseCase {
	return &sessionUseCase{repo: repo}
}
