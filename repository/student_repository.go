package repository

import (
	"database/sql"
	"fmt"
	"time"

	"final-project-kelompok-1/model"
	"final-project-kelompok-1/utils/common"
)

type StudentRepository interface {
	Create(payload model.Student) (model.Student, error)
	GetById(id string) (model.Student, error)
	GetAll() ([]model.Student, error)
	Update(payload model.Student, id string) (model.Student, error)
	Delete(id string) (model.Student, error)
	GetByStudentEmail(email string) (model.Student, error)
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
		"student",
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
		&student.Role,
		&student.IsDeleted,
	)

	if err != nil {
		fmt.Print("Error student in repo : ", err.Error())
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
		&student.Role,
	)
	if err != nil {
		return model.Student{}, err
	}
	return student, nil
}

func (s *studentRepository) GetAll() ([]model.Student, error) {
	var students []model.Student

	rows, err := s.db.Query(common.GetAllStudent)

	if err != nil {
		return students, err
	}
	for rows.Next() {
		var student model.Student
		err := rows.Scan(
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
			&student.IsDeleted,
		)

		if err != nil {
			fmt.Println("error in repo :", err.Error())
			return students, nil
		}

		students = append(students, student)
	}

	return students, nil
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
		"student",
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
		&student.Role,
	)
	if err != nil {
		fmt.Println("Error in repo student update : ", err.Error())
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

func (s *studentRepository) GetByStudentEmail(email string) (model.Student, error) {
	var student model.Student
	err := s.db.QueryRow(common.GetByStudentEmail, email).Scan(
		&student.StudentID,
		&student.Fullname,
		&student.BirthDate,
		&student.BirthPlace,
		&student.Address,
		&student.Education,
		&student.Institution,
		&student.Job,
		&student.Role,
		&student.Email,
		&student.Password,
		&student.CreatedAt,
		&student.UpdatedAt,
		&student.IsDeleted,
	)
	if err != nil {
		fmt.Println("Error in repo : ", err.Error())
		return model.Student{}, err
	}
	return student, nil
}

func NewStudentRepository(db *sql.DB) StudentRepository {
	return &studentRepository{db: db}
}
