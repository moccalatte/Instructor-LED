package repository

import (
	"database/sql"
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

func (suite *UserRepositoryTestSuite) TestCreateUser_Success() {
	dummy := model.Users{
		UserID:    "sdsdsdsdsdsdsd",
		Fullname:  "Joko Santoso",
		Role:      "admin",
		Email:     "chril@example.com",
		Password:  "230104",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now().Add(3 * 24 * time.Hour),
		IsDeleted: false,
	}

	suite.sqlmock.ExpectBegin()

	rows := sqlmock.NewRows([]string{"user_id", "fullname", "role", "email", "password", "created_at", "updated_at", "is_deleted"}).
		AddRow(dummy.UserID, dummy.Fullname, dummy.Role, dummy.Email, dummy.Password, dummy.CreatedAt, dummy.UpdatedAt, dummy.IsDeleted)
	suite.sqlmock.ExpectQuery("insert into users").WillReturnRows(rows)
	suite.sqlmock.ExpectCommit()

	actual, err := suite.repo.Create(dummy)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), dummy.UserID, actual.UserID)
}

func (suite *UserRepositoryTestSuite) TestGetUserById_Success() {
	UserID := "809808098098"
	dummyDB := model.Users{
		UserID:    "809808098098",
		Fullname:  "Joko Santoso",
		Role:      "admin",
		Email:     "chril@example.com",
		Password:  "230104",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsDeleted: false,
	}

	suite.sqlmock.ExpectBegin()

	rows := sqlmock.NewRows([]string{"user_id", "fullname", "role", "email", "password", "created_at", "updated_at", "is_deleted"}).
		AddRow(dummyDB.UserID, dummyDB.Fullname, dummyDB.Role, dummyDB.Email, dummyDB.Password, dummyDB.CreatedAt, dummyDB.UpdatedAt, dummyDB.IsDeleted)
	suite.sqlmock.ExpectQuery("select * from users where user_id = $1").WithArgs(UserID).WillReturnRows(rows)

	suite.sqlmock.ExpectCommit()

	actual, err := suite.repo.GetById(UserID)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), dummyDB.UserID, actual.UserID)
	assert.Equal(suite.T(), dummyDB.Fullname, actual.Fullname)
	assert.Equal(suite.T(), dummyDB.Role, actual.Role)
	assert.Equal(suite.T(), dummyDB.Email, actual.Email)
	assert.Equal(suite.T(), dummyDB.Password, actual.Password)
	assert.Equal(suite.T(), dummyDB.CreatedAt, actual.CreatedAt)
	assert.Equal(suite.T(), dummyDB.UpdatedAt, actual.UpdatedAt)
	assert.Equal(suite.T(), dummyDB.IsDeleted, actual.IsDeleted)
}
