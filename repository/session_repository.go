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
	UpdateNote(payload model.Session, id string) (model.Session, error)
	Delete(id string) (model.Session, error)
	GetAllSession() ([]model.Session, error)
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
		payload.Note,
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
		&session.Note,
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
		&session.Note,
	)
	if err != nil {
		return model.Session{}, err
	}
	return session, nil
}

func (s *sessionRepository) GetAllSession() ([]model.Session, error) {
	var sessions []model.Session

	rows, err := s.db.Query(common.GetAllSession)

	if err != nil {
		return sessions, err
	}
	for rows.Next() {
		var session model.Session
		err := rows.Scan(
			&session.SessionID,
			&session.Title,
			&session.Description,
			&session.SessionDate,
			&session.SessionTime,
			&session.SessionLink,
			&session.TrainerID,
			&session.Note,
			&session.IsDeleted,
		)

		if err != nil {
			fmt.Println("error in repo :", err.Error())
			return sessions, nil
		}

		sessions = append(sessions, session)
	}

	return sessions, nil
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
		payload.Note,
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
		&session.Note,
		&session.CreatedAt,
		&session.UpdatedAt,
		&session.IsDeleted,
	)
	if err != nil {
		fmt.Println("Error in repo : ", err.Error())
		return model.Session{}, tx.Rollback()
	}

	if err := tx.Commit(); err != nil {
		return model.Session{}, err
	}

	return session, nil
}

func (s *sessionRepository) UpdateNote(payload model.Session, id string) (model.Session, error) {
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
	err = tx.QueryRow(common.UpdateNote,
		payload.Note,
		time.Now(),
		id).Scan(
		&session.SessionID,
		&session.Title,
		&session.Description,
		&session.SessionDate,
		&session.SessionTime,
		&session.SessionLink,
		&session.TrainerID,
		&session.Note,
		&session.CreatedAt,
		&session.UpdatedAt,
		&session.IsDeleted,
	)
	if err != nil {
		fmt.Println("Error in repo : ", err.Error())
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
		&session.Note,
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

func (s *sessionRepository) FindAll() ([]model.Session, error) {
	rows, err := s.db.Query(common.GetAllDatas, false)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sessions []model.Session
	for rows.Next() {
		var session model.Session
		err := rows.Scan(
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
			return nil, err
		}
		sessions = append(sessions, session)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return sessions, nil
}

func NewSessionRepository(db *sql.DB) SessionRepository {
	return &sessionRepository{db: db}
}
