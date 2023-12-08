package usecase

import (
	"final-project-kelompok-1/model"
	"final-project-kelompok-1/model/dto"
	"final-project-kelompok-1/repository"
	"fmt"
)

type SessionUseCase interface {
	FindSessionByID(id string) (model.Session, error)
	AddSession(payload dto.SessionRequestDto) (model.Session, error)
	// UpdateSession(payload dto.SessionRequestDto) (model.Session, error)
	// DeleteSession(payload dto.SessionRequestDto) (model.Session, error)
}

type sessionUseCase struct {
	repo      repository.SessionRepository
	trainerUC UserUseCase
}

func (s *sessionUseCase) FindSessionByID(id string) (model.Session, error) {
	session, err := s.repo.GetById(id)
	if err != nil {
		return model.Session{}, fmt.Errorf("Session with ID  not found")
	}

	return session, nil
}

func (s *sessionUseCase) AddSession(payload dto.SessionRequestDto) (model.Session, error) {
	trainer, err := s.trainerUC.FindUserByID(payload.TrainerId)

	if err != nil {
		return model.Session{}, err
	}
	newSession := model.Session{
		TrainerAdminID: trainer,
		IsDeleted:      payload.IsDeleted,
	}

	session, err := s.repo.CreateSession(newSession)
	if err != nil {
		return model.Session{}, fmt.Errorf("failed to create transaksi : %s", err.Error())
	}

	return session, nil
}

func NewSessionUseCase(repo repository.SessionRepository, trainerUC UserUseCase) SessionUseCase {
	return &sessionUseCase{repo: repo, trainerUC: trainerUC}
}
