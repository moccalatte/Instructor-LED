package repository

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"final-project-kelompok-1/model"
)

type StudentRepositoryTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	sqlmock sqlmock.Sqlmock
	repo    StudentRepository
}

func (suite *StudentRepositoryTestSuite) SetupTest() {
	db, sqlmock, err := sqlmock.New()
	assert.NoError(suite.T(), err)
	suite.mockDb = db
	suite.sqlmock = sqlmock
	suite.repo = NewStudentRepository(suite.mockDb)
}

func TestStudentRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(StudentRepositoryTestSuite))
}

func (suite *StudentRepositoryTestSuite) TesCreateStudent_Succes() {
	dummy := model.Student{
		StudentID:   "stdn1212",
		Fullname:    "golangungil",
		BirthDate:   "2004-12-10",
		BirthPlace:  "Medan",
		Address:     "Jakarta",
		Education:   "Strata 1",
		Institution: "ITICM",
		Job:         "IT Consultan",
		Email:       "Noiyutyt@gmail.com",
		Password:    "1212112",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now().Add(24 * 24 * time.Hour),
		IsDeleted:   false,
	}

	suite.sqlmock.ExpectBegin()

	rows := sqlmock.NewRows([]string{"student_id ,fullname, birth_date, birth_place, address, education, institution, job, email, password, created_at, updated_at, is_deleted"}).AddRow(dummy.StudentID, dummy.Fullname, dummy.BirthDate, dummy.BirthPlace, dummy.Address, dummy.Education, dummy.Institution, dummy.Job, dummy.Email, dummy.Password, dummy.CreatedAt, dummy.UpdatedAt, dummy.IsDeleted)
	suite.sqlmock.ExpectQuery("insert into student").WillReturnRows(rows)
	suite.sqlmock.ExpectCommit()
	actual, err := suite.repo.Create(dummy)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), dummy.StudentID, actual.StudentID)
}

func (suite *StudentRepositoryTestSuite) TesCreateStudent_Failed() {
	dummy := model.Student{
		StudentID:   "stdn1212",
		Fullname:    "golangungil",
		BirthDate:   "2004-12-10",
		BirthPlace:  "Medan",
		Address:     "Jakarta",
		Education:   "Strata 1",
		Institution: "ITICM",
		Job:         "IT Consultan",
		Email:       "Noiyutyt@gmail.com",
		Password:    "1212112",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now().Add(24 * 24 * time.Hour),
		IsDeleted:   false,
	}

	suite.sqlmock.ExpectBegin().WillReturnError(errors.New("error begin"))
	_, err := suite.repo.Create(dummy)
	assert.Error(suite.T(), err)

	suite.sqlmock.ExpectBegin()
	suite.sqlmock.ExpectQuery("insert into student").WillReturnError(errors.New("insert failed"))
	_, err = suite.repo.Create(dummy)
	assert.Error(suite.T(), err)

	suite.sqlmock.ExpectBegin()
	rows := sqlmock.NewRows([]string{"student_id ,fullname, birth_date, birth_place, address, education, institution, job, email, password, created_at, updated_at, is_deleted"}).AddRow(dummy.StudentID, dummy.Fullname, dummy.BirthDate, dummy.BirthPlace, dummy.Address, dummy.Education, dummy.Institution, dummy.Job, dummy.Email, dummy.Password, dummy.CreatedAt, dummy.UpdatedAt, dummy.IsDeleted)

	suite.sqlmock.ExpectQuery("insert into student").WillReturnRows(rows)
	suite.sqlmock.ExpectCommit()
	_, err = suite.repo.Create(dummy)
	assert.Error(suite.T(), err)
	assert.Contains(suite.T(), err.Error(), "some error")
}
