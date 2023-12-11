package repository

import (
	"database/sql"
	"fmt"
	"time"

	"final-project-kelompok-1/model"
	"final-project-kelompok-1/utils/common"
)

type SessionRepository interface {
	Create(payload model.Session) (model.Session, error)
	GetById(id string) (model.Session, error)
	Update(payload model.Session, id string) (model.Session, error)
	Delete(id string) (model.Session, error)
}

type sessionRepository struct {
	db *sql.DB
}

func (s *sessionRepository) Create(payload model.Session) (model.Session, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return model.Session{}, err
	}
	var session model.Session
	err = tx.QueryRow(common.CreateSesion,
		payload.Title,
		payload.Description,
		payload.SessionDate,
		payload.SessionTime,
		payload.SessionLink,
		payload.TrainerID,
		time.Now(),
		time.Now(),
		false).Scan(
		&session.SessionID,
		&session.Title,
		&session.Description,
		&session.SessionDate,
		&session.SessionTime,
		&session.SessionLink,
		&session.TrainerID,
		&session.CreatedAt,
		&session.UpdatedAt,
		&session.IsDeleted,
	)
	fmt.Print(err, "SESSION REPO")
	if err != nil {
		return model.Session{}, tx.Rollback()
	}

	if err := tx.Commit(); err != nil {
		return model.Session{}, err
	}

	return session, nil
}

func (s *sessionRepository) GetById(id string) (model.Session, error) {
	var session model.Session
	err := s.db.QueryRow(common.GetSessionById, id).Scan(
		&session.SessionID,
		&session.Title,
		&session.Description,
		&session.SessionDate,
		&session.SessionTime,
		&session.SessionLink,
		&session.TrainerID,
		&session.CreatedAt,
		&session.UpdatedAt,
		&session.IsDeleted,
	)
	if err != nil {
		return model.Session{}, err
	}
	return session, nil
}

func (s *sessionRepository) Update(payload model.Session, id string) (model.Session, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return model.Session{}, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	var session model.Session
	err = tx.QueryRow(common.UpdateSessionById,
		payload.Title,
		payload.Description,
		payload.SessionDate,
		payload.SessionTime,
		payload.SessionLink,
		payload.TrainerID,
		time.Now(),
		false,
		id).Scan(
		&session.SessionID,
		&session.Title,
		&session.Description,
		&session.SessionDate,
		&session.SessionTime,
		&session.SessionLink,
		&session.TrainerID,
		&session.CreatedAt,
		&session.UpdatedAt,
		&session.IsDeleted,
	)
	if err != nil {
		return model.Session{}, tx.Rollback()
	}

	if err := tx.Commit(); err != nil {
		return model.Session{}, err
	}

	return session, nil
}

func (s *sessionRepository) Delete(id string) (model.Session, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return model.Session{}, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	var session model.Session
	err = tx.QueryRow(common.DeleteSessionById,
		true,
		id).Scan(
		&session.SessionID,
		&session.Title,
		&session.Description,
		&session.SessionDate,
		&session.SessionTime,
		&session.SessionLink,
		&session.TrainerID,
		&session.CreatedAt,
		&session.UpdatedAt,
		&session.IsDeleted,
	)
	if err != nil {
		return model.Session{}, tx.Rollback()
	}

	if err := tx.Commit(); err != nil {
		return model.Session{}, err
	}

	return session, nil
}

func NewSessionRepository(db *sql.DB) SessionRepository {
	return &sessionRepository{db: db}
}
