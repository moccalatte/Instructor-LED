package usecase

import (
	"final-project-kelompok-1/model"
	"final-project-kelompok-1/model/dto"
	"final-project-kelompok-1/repository"
	"fmt"
)

type StudentUseCase interface {
	AddStudent(payload dto.StudentRequestDto) (model.Student, error)
	FindStudentByID(id string) (model.Student, error)
	GetAllStudent() ([]model.Student, error)
	UpdateStudent(payload dto.StudentRequestDto, id string) (model.Student, error)
	DeleteStudent(id string) (model.Student, error)
}

type studentUseCase struct {
	repo repository.StudentRepository
}

func (s *studentUseCase) AddStudent(payload dto.StudentRequestDto) (model.Student, error) {
	newStudent := model.Student{
		Fullname:    payload.Fullname,
		BirthDate:   payload.BirthDate,
		BirthPlace:  payload.BirthPlace,
		Address:     payload.Address,
		Education:   payload.Education,
		Institution: payload.Institution,
		Job:         payload.Job,
		Email:       payload.Email,
		Password:    payload.Password,
	}

	createdStudent, err := s.repo.Create(newStudent)

	if err != nil {
		return model.Student{}, fmt.Errorf("failed to save data: %s", err.Error())
	}

	return createdStudent, nil
}

func (s *studentUseCase) FindStudentByID(id string) (model.Student, error) {
	student, err := s.repo.GetById(id)

	if err != nil {
		return model.Student{}, fmt.Errorf("failed get data by id : %s", err.Error())
	}
	return student, nil
}

func (s *studentUseCase) GetAllStudent() ([]model.Student, error) {
	studentAll, err := s.repo.GetAll()

	if err != nil {
		fmt.Println("Error Get All Data in use case : ", err.Error())
		return studentAll, fmt.Errorf("failed to find data : %v", err.Error())
	}

	return studentAll, nil
}

func (s *studentUseCase) UpdateStudent(payload dto.StudentRequestDto, id string) (model.Student, error) {
	newStudent := model.Student{
		Fullname: payload.Fullname,

		BirthDate:   payload.BirthDate,
		BirthPlace:  payload.BirthPlace,
		Address:     payload.Address,
		Education:   payload.Education,
		Institution: payload.Institution,
		Job:         payload.Job,
		Email:       payload.Email,
		Password:    payload.Password,
	}

	updatedStudent, err := s.repo.Update(newStudent, id)

	if err != nil {
		return model.Student{}, fmt.Errorf("failed update data by id : %s", err.Error())
	}

	return updatedStudent, nil
}

func (s *studentUseCase) DeleteStudent(id string) (model.Student, error) {
	deleteStudent, err := s.repo.Delete(id)

	if err != nil {
		return model.Student{}, fmt.Errorf("failed to deleted data : %s", err.Error())
	}

	return deleteStudent, nil
}

func NewStudentUseCase(repo repository.StudentRepository) StudentUseCase {
	return &studentUseCase{repo: repo}
}