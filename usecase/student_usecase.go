package usecase

import (
	"errors"
	"final-project-kelompok-1/model"
	"final-project-kelompok-1/model/dto"
	"final-project-kelompok-1/repository"
	"final-project-kelompok-1/utils/common"
	"fmt"
)

type StudentUseCase interface {
	AddStudent(payload dto.StudentRequestDto) (model.Student, error)
	FindStudentByID(id string) (model.Student, error)
	GetAllStudent() ([]model.Student, error)
	UpdateStudent(payload dto.StudentRequestDto, id string) (model.Student, error)
	DeleteStudent(id string) (model.Student, error)
	FindByEmailPassword(email string, password string) (model.Student, error)
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

	newPassword, err := common.GeneratePasswordHash(payload.Password)
	if err != nil {
		return model.Student{}, err
	}

	newStudent.Password = newPassword
	createdStudent, err := s.repo.Create(newStudent)

	if err != nil {
		fmt.Println("Error student  in Usecase :", err.Error())
		return model.Student{}, fmt.Errorf("failed to save data: %s", err.Error())
	}

	return createdStudent, nil
}

func (s *studentUseCase) FindStudentByID(id string) (model.Student, error) {
	student, err := s.repo.GetById(id)

	if err != nil {
		fmt.Println("Error in usecase student : ", err.Error())
		return model.Student{}, fmt.Errorf("failed get data by id : %s", err.Error())
	}
	return student, nil
}

func (s *studentUseCase) GetAllStudent() ([]model.Student, error) {
	var sliceStudent []model.Student
	userData, err := s.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to find all data : %s", err.Error())
	}
	return append(sliceStudent, userData...), nil
}

func (s *studentUseCase) UpdateStudent(payload dto.StudentRequestDto, id string) (model.Student, error) {
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
	newPassword, err := common.GeneratePasswordHash(payload.Password)
	if err != nil {
		fmt.Println("Error for generate password : ", err.Error())
		return model.Student{}, err
	}

	newStudent.Password = newPassword
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

func (s *studentUseCase) FindByEmailPassword(email string, password string) (model.Student, error) {
	student, err := s.repo.GetByStudentEmail(email)
	fmt.Println(student)

	if err != nil {
		fmt.Println("Error in usecase : ", err.Error())
		return model.Student{}, errors.New("invalid email or password")
	}

	if err := common.ComparePasswordHash(student.Password, password); err != nil {
		return model.Student{}, err
	}

	student.Password = ""
	return student, nil
}

func NewStudentUseCase(repo repository.StudentRepository) StudentUseCase {
	return &studentUseCase{repo: repo}
}
