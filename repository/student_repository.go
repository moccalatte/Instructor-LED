package repository

import (
	"database/sql"
	"time"

	"final-project-kelompok-1/model"
	"final-project-kelompok-1/utils/common"
)

type StudentRepository interface {
	Create(payload model.Student) (model.Student, error)
	GetById(id string) (model.Student, error)
	Update(payload model.Student, id string) (model.Student, error)
	Delete(id string) (model.Student, error)
}

type studentRepository struct {
	db *sql.DB
}

func (s *studentRepository) Create(payload model.Student) (model.Student, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return model.Student{}, err
	}
	var student model.Student
	err = tx.QueryRow(common.CreateStudent,
		payload.Fullname,
		payload.BirthDate,
		payload.BirthPlace,
		payload.Address,
		payload.Education,
		payload.Institution,
		payload.Job,
		payload.Email,
		payload.Password,
		time.Now(),
		false).Scan(
		&student.StudentID,
		&student.Fullname,
		&student.BirthDate,
		&student.BirthPlace,
		&student.Address,
		&student.Education,
		&student.Institution,
		&student.Job,
		&student.Email,
		&student.Password,
		&student.CreatedAt,
		&student.UpdatedAt,
		&student.IsDeleted,
	)
	if err != nil {
		return model.Student{}, tx.Rollback()
	}

	if err := tx.Commit(); err != nil {
		return model.Student{}, err
	}

	return student, nil

}

func (s *studentRepository) GetById(id string) (model.Student, error) {
	var student model.Student
	err := s.db.QueryRow(common.GetStudentByid, id).Scan(
		&student.StudentID,
		&student.Fullname,
		&student.BirthDate,
		&student.BirthPlace,
		&student.Address,
		&student.Education,
		&student.Institution,
		&student.Job,
		&student.Email,
		&student.Password,
		&student.CreatedAt,
		&student.UpdatedAt,
		&student.IsDeleted,
	)
	if err != nil {
		return model.Student{}, err
	}
	return student, nil
}

func (s *studentRepository) Update(payload model.Student, id string) (model.Student, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return model.Student{}, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	var student model.Student
	err = tx.QueryRow(common.UpdateStudentbyId,
		payload.Fullname,
		payload.BirthDate,
		payload.BirthPlace,
		payload.Address,
		payload.Education,
		payload.Institution,
		payload.Job,
		payload.Email,
		payload.Password,
		time.Now(),
		false,
		id).Scan(
			&student.StudentID,
			&student.Fullname,
			&student.BirthDate,
			&student.BirthPlace,
			&student.Address,
			&student.Education,
			&student.Institution,
			&student.Job,
			&student.Email,
			&student.Password,
			&student.CreatedAt,
			&student.UpdatedAt,
			&student.IsDeleted,
	)
	if err != nil {
		return model.Student{}, tx.Rollback()
	}

	if err := tx.Commit(); err != nil {
		return model.Student{}, err
	}

	return student, nil

}

func (s *studentRepository) Delete(id string) (model.Student, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return model.Student{}, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	var student model.Student
	err = tx.QueryRow(common.DeleteStudentById,
		true,
		id).Scan(
			&student.StudentID,
			&student.Fullname,
			&student.BirthDate,
			&student.BirthPlace,
			&student.Address,
			&student.Education,
			&student.Institution,
			&student.Job,
			&student.Email,
			&student.Password,
			&student.CreatedAt,
			&student.UpdatedAt,
			&student.IsDeleted,
	)
	if err != nil {
		return model.Student{}, tx.Rollback()
	}

	if err := tx.Commit(); err != nil {
		return model.Student{}, err
	}

	return student, nil

}

func NewStudentRepository(db *sql.DB) StudentRepository {
	return &studentRepository{db: db}
}