package usecase

import (
	"final-project-kelompok-1/model"
	"final-project-kelompok-1/model/dto"
	"final-project-kelompok-1/repository"
)

type StudentUseCase interface {
	AddStudent(payload dto.StudentRequestDto) (model.Student, error)
	FindStudentByID(id int) (model.Student, error)
	UpdateStudent(payload dto.StudentRequestDto, id int) (model.Student, error)
	DeleteStudent(id int) (model.Student, error)
}

type studentUseCase struct {
	repo repository.StudentRepository
}

func (s *studentUseCase) AddStudent(payload dto.StudentRequestDto) (model.Student, error) {

}

func (s *studentUseCase) FindStudentByID(id int) (model.Student, error) {

}

func (s *studentUseCase) UpdateStudent(payload dto.StudentRequestDto, id int) (model.Student, error) {

}

func (s *studentUseCase) DeleteStudent(id int) (model.Student, error) {

}

func NewStudentUseCase(repo repository.StudentRepository) StudentUseCase {
	return &studentUseCase{repo: repo}
}
