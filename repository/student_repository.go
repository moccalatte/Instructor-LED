package repository

import (
	"database/sql"
	"final-project-kelompok-1/model"
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
}

func (s *studentRepository) GetById(id string) (model.Student, error) {
}

func (s *studentRepository) Update(payload model.Student, id string) (model.Student, error) {
}

func (s *studentRepository) Delete(id string) (model.Student, error) {
}

func NewStudentRepository(db *sql.DB) StudentRepository {
	return &studentRepository{db: db}
}
