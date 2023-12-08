package repository

import (
	"database/sql"
	"final-project-kelompok-1/model"
	"final-project-kelompok-1/utils/common"
	"fmt"
)

type SessionRepository interface {
	GetById(id string) (model.Session, error)
	CreateSession(payload model.Session) (model.Session, error)
	UpdateSession(payload model.Session) (model.Session, error)
	DeleteSession(payload model.Session) (model.Session, error)
}

type sessionRepository struct {
	db *sql.DB
}

func (s *sessionRepository) GetById(id string) (model.Session, error) {
	var session model.Session

	err := s.db.QueryRow(common.GetSession, id).Scan(
		&session.SessionID,
		&session.Date,
		&session.TrainerAdminID,
	)

	if err != nil {
		return model.Session{}, err
	}

	return session, nil
}

func (s *sessionRepository) CreateSession(payload model.Session) (model.Session, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return model.Session{}, err
	}

	var session model.Session
	err = tx.QueryRow(common.CreateSession, payload.Date, payload.TrainerAdminID).Scan(
		&session.SessionID,
		&session.Date,
		&session.TrainerAdminID,
	)

	if err != nil {
		fmt.Println("Error Insert Session : ", err)
		return model.Session{}, tx.Rollback()
	}

	if err := tx.Commit(); err != nil {
		return model.Session{}, err
	}

	return session, nil
}

func (s *sessionRepository) UpdateSession(payload model.Session) (model.Session, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return model.Session{}, err
	}

	var session model.Session
	err = tx.QueryRow(common.UpdateSessionById,
		payload.Date,
		payload.TrainerAdminID,
		payload.SessionID).Scan(
		&session.Date,
		&session.TrainerAdminID,
		&session.SessionID,
	)

	if err != nil {
		fmt.Println("Error Update Session : ", err)
		return model.Session{}, tx.Rollback()
	}

	if err := tx.Commit(); err != nil {
		return model.Session{}, err

	}

	return session, nil
}

func (s *sessionRepository) DeleteSession(payload model.Session) (model.Session, error) {
	tx, err := s.db.Begin()

	if err != nil {
		return model.Session{}, err
	}

	var session model.Session
	err = tx.QueryRow(common.DeleteSessionById, 1, payload.SessionID).Scan(
		&session.IsDeleted,
		&session.SessionID,
	)

	if err != nil {
		fmt.Println("Error Delete Session : ", err)
		return model.Session{}, err
	}

	return session, nil
}

func NewSessionRepository(db *sql.DB) SessionRepository {
	return &sessionRepository{db: db}
}
