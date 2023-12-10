package repository

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"final-project-kelompok-1/model"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	sqlmock sqlmock.Sqlmock
	repo    UserRepository
}

func (suite *UserRepositoryTestSuite) SetupTest() {
	db, sqlmock, err := sqlmock.New()
	assert.NoError(suite.T(), err)
	suite.mockDb = db
	suite.sqlmock = sqlmock
	suite.repo = NewUserRepository(suite.mockDb)
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}

func (suite *UserRepositoryTestSuite) TesCreateUser_Succes() {
	dummy := model.Users{
		UserID:    "sdsdsdsdsdsdsd",
		Fullname:  "Joko santoso",
		Role:      "admin",
		Email:     "chril.com",
		Password:  "230104",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now().Add(3 * 24 * time.Hour),
		IsDeleted: false,
	}

	suite.sqlmock.ExpectBegin()

	rows := sqlmock.NewRows([]string{"user_id, fullname, role, email, password, created_at, updated_at, is_deleted"}).AddRow(dummy.UserID, dummy.Fullname, dummy.Role, dummy.Email, dummy.Password, dummy.CreatedAt, dummy.UpdatedAt, dummy.IsDeleted)
	suite.sqlmock.ExpectQuery("insert into users").WillReturnRows(rows)
	suite.sqlmock.ExpectCommit()
	actual, err := suite.repo.Create(dummy)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), dummy.UserID, actual.UserID)
}

func (suite *UserRepositoryTestSuite) TesCreateUser_Failed() {
	dummy := model.Users{
		UserID:    "sdsdsdsdsdsdsd",
		Fullname:  "Joko santoso",
		Role:      "admin",
		Email:     "chril.com",
		Password:  "230104",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now().Add(3 * 24 * time.Hour),
		IsDeleted: false,
	}

	suite.sqlmock.ExpectBegin().WillReturnError(fmt.Errorf("begin error"))
	suite.sqlmock.ExpectQuery("insert into users").WillReturnError(fmt.Errorf("some error"))
	suite.sqlmock.ExpectRollback()
	_, err := suite.repo.Create(dummy)
	assert.Nil(suite.T(), err)
	assert.Error(suite.T(), err)
	assert.Contains(suite.T(), err.Error(), "some error")
}
