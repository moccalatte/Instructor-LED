package repository

import (
	"database/sql"
	"time"

	"final-project-kelompok-1/model"
	"final-project-kelompok-1/utils/common"
)

type QuestionRepository interface {
	Create(payload model.Question) (model.Question, error)
	GetById(id string) (model.Question, error)
	GetByStudentId(id string) (model.Question, error)
	Update(payload model.Question, id string) (model.Question, error)
	Delete(id string) (model.Question, error)
	Answer(payload model.Question, id string) (model.Question, error)
}
type questionRepository struct {
	db *sql.DB
}

func (q *questionRepository) Create(payload model.Question) (model.Question, error) {
	tx, err := q.db.Begin()
	if err != nil {
		return model.Question{}, err
	}
	var question model.Question
	err = tx.QueryRow(common.CreateQuestion,
		payload.SessionID,
		payload.StudentID,
		payload.TrainerID,
		payload.Title,
		payload.Description,
		payload.CourseID,
		payload.Image,
		payload.Answer,
		payload.Status,
		time.Now(),
		false).Scan(
		&question.QuestionID,
		&question.SessionID,
		&question.StudentID,
		&question.TrainerID,
		&question.Title,
		&question.Description,
		&question.CourseID,
		&question.Image,
		&question.Answer,
		&question.Status,
		&question.CreatedAt,
		&question.UpdatedAt,
		&question.IsDeleted,
	)

	if err != nil {
		return model.Question{}, tx.Rollback()
	}

	if err := tx.Commit(); err != nil {
		return model.Question{}, err
	}

	return question, nil
}

func (q *questionRepository) GetById(id string) (model.Question, error) {
	var question model.Question
	err := q.db.QueryRow(common.GetQuestionById, id).Scan(
		&question.QuestionID,
		&question.SessionID,
		&question.StudentID,
		&question.TrainerID,
		&question.Title,
		&question.Description,
		&question.CourseID,
		&question.Image,
		&question.Answer,
		&question.Status,
		&question.CreatedAt,
		&question.UpdatedAt,
		&question.IsDeleted,
	)

	if err != nil {
		return model.Question{}, err
	}

	return question, nil
}

func (q *questionRepository) GetByStudentId(id string) (model.Question, error) {
	var question model.Question
	err := q.db.QueryRow(common.GetQuestionByStudentId, id).Scan(
		&question.QuestionID,
		&question.SessionID,
		&question.StudentID,
		&question.TrainerID,
		&question.Title,
		&question.Description,
		&question.CourseID,
		&question.Image,
		&question.Answer,
		&question.Status,
		&question.CreatedAt,
		&question.UpdatedAt,
		&question.IsDeleted,
	)

	if err != nil {
		return model.Question{}, err
	}

	return question, nil
}

func (q *questionRepository) Update(payload model.Question, id string) (model.Question, error) {
	tx, err := q.db.Begin()
	if err != nil {
		return model.Question{}, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	var question model.Question
	err = tx.QueryRow(
		common.UpdateQuestionById,
		payload.SessionID,
		payload.StudentID,
		payload.TrainerID,
		payload.Title,
		payload.Description,
		payload.CourseID,
		payload.Image,
		payload.Answer,
		payload.Status,
		time.Now(),
		false,
		id).Scan(
		&question.QuestionID,
		&question.SessionID,
		&question.StudentID,
		&question.TrainerID,
		&question.Title,
		&question.Description,
		&question.CourseID,
		&question.Image,
		&question.Answer,
		&question.Status,
		&question.CreatedAt,
		&question.UpdatedAt,
		&question.IsDeleted,
	)

	if err != nil {
		return model.Question{}, tx.Rollback()
	}

	if err := tx.Commit(); err != nil {
		return model.Question{}, err
	}

	return question, nil

}

func (q *questionRepository) Delete(id string) (model.Question, error) {
	tx, err := q.db.Begin()
	if err != nil {
		return model.Question{}, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	var question model.Question
	err = tx.QueryRow(
		common.DeleteQuestionById,
		true,
		id).Scan(
		&question.QuestionID,
		&question.SessionID,
		&question.StudentID,
		&question.TrainerID,
		&question.Title,
		&question.Description,
		&question.CourseID,
		&question.Image,
		&question.Answer,
		&question.Status,
		&question.CreatedAt,
		&question.UpdatedAt,
		&question.IsDeleted,
	)
	if err != nil {
		return model.Question{}, tx.Rollback()
	}

	if err := tx.Commit(); err != nil {
		return model.Question{}, err
	}

	return question, nil

}

func (q *questionRepository) Answer(payload model.Question, id string) (model.Question, error) {
	tx, err := q.db.Begin()
	if err != nil {
		return model.Question{}, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	var question model.Question
	err = tx.QueryRow(
		common.AnswerQuestionById,
		payload.Answer,
		time.Now(),
		id).Scan(
		&question.QuestionID,
		&question.SessionID,
		&question.StudentID,
		&question.TrainerID,
		&question.Title,
		&question.Description,
		&question.CourseID,
		&question.Image,
		&question.Answer,
		&question.Status,
		&question.CreatedAt,
		&question.UpdatedAt,
		&question.IsDeleted,
	)
	if err != nil {
		return model.Question{}, tx.Rollback()
	}

	if err := tx.Commit(); err != nil {
		return model.Question{}, err
	}

	return question, nil

}

func NewQuestionRepository(db *sql.DB) QuestionRepository {
	return &questionRepository{db: db}
}
