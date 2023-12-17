package repository

import (
	"database/sql"
	"final-project-kelompok-1/model"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
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

func (suite *UserRepositoryTestSuite) TestGetById() {
	dummy := model.Users{
		UserID:    "80980986875676768798797",
		Fullname:  "Gopan andika andre rizki taulani",
		Role:      "owner",
		Email:     "apa2ajalah@gmail.com",
		Password:  "satusatu",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsDeleted: false,
	}
	query := "select \\* from users where user_id = \\$1;"
	user_id := "80980986875676768798797"

	rows := sqlmock.NewRows([]string{"user_id", "fullname", "role", "email", "password", "created_at", "updated_at", "is_deleted"}).
		AddRow(dummy.UserID, dummy.Fullname, dummy.Role, dummy.Email, dummy.Password, dummy.CreatedAt, dummy.UpdatedAt, dummy.IsDeleted)

	suite.sqlmock.ExpectQuery(query).WithArgs(user_id).WillReturnRows(rows)

	actual, err := suite.repo.GetById(user_id)

	assert.Nil(suite.T(), err, "error should be nil")
	assert.NoError(suite.T(), err, "error should be nil")
	assert.Equal(suite.T(), dummy.UserID, actual.UserID, "UserID should match")
	assert.Equal(suite.T(), dummy.Fullname, actual.Fullname, "Fullname should match")
	assert.Equal(suite.T(), dummy.Role, actual.Role, "Role should match")
	assert.Equal(suite.T(), dummy.Email, actual.Email, "Email should match")
	assert.Equal(suite.T(), dummy.Password, actual.Password, "Password should match")
	assert.WithinDuration(suite.T(), dummy.CreatedAt, actual.CreatedAt, time.Second, "CreatedAt should match")
	assert.WithinDuration(suite.T(), dummy.UpdatedAt, actual.UpdatedAt, time.Second, "UpdatedAt should match")
	assert.Equal(suite.T(), dummy.IsDeleted, actual.IsDeleted, "IsDeleted should match")
}

func (suite *UserRepositoryTestSuite) TestCreate() {

}
func (suite *UserRepositoryTestSuite) TestGetAll() {

}
func (suite *UserRepositoryTestSuite) TestUpdate() {

}

func (suite *UserRepositoryTestSuite) TestDelete() {

}

func (suite *UserRepositoryTestSuite) TestGetByUsername() {
	dummy := model.Users{
		UserID:    "12345687678",
		Fullname:  "andikairfanandrerizki",
		Role:      "admin",
		Email:     "ccccc@gmail.com",
		Password:  "1232131",
		IsDeleted: false,
	}

	query := "select * from users where fullname = $1 OR email = $1 returning user_id, fullname, role, email, password, created_at, updated_at, is_deleted;"
	email := "ccccc@gmail.com"

	// Set expectations for the mock
	rows := sqlmock.NewRows([]string{"user_id", "fullname", "role", "email", "password", "created_at", "updated_at", "is_deleted"}).
		AddRow(dummy.UserID, dummy.Fullname, dummy.Role, dummy.Email, dummy.Password, dummy.CreatedAt, dummy.UpdatedAt, dummy.IsDeleted)
	suite.sqlmock.ExpectQuery(query).WithArgs(email).WillReturnRows(rows)

	// Call the method being tested
	actual, err := suite.repo.GetByUsername(email)

	// Assertions
	assert.Nil(suite.T(), err, "error should be nil")
	assert.Equal(suite.T(), dummy.UserID, actual.UserID, "UserID should match")
	assert.Equal(suite.T(), dummy.Fullname, actual.Fullname, "Fullname should match")
	assert.Equal(suite.T(), dummy.Role, actual.Role, "Role should match")
	assert.Equal(suite.T(), dummy.Email, actual.Email, "Email should match")
	assert.Equal(suite.T(), dummy.Password, actual.Password, "Password should match")
	assert.Equal(suite.T(), dummy.IsDeleted, actual.IsDeleted, "IsDeleted should match")

	// Ensure all expectations were met
	assert.NoError(suite.T(), suite.sqlmock.ExpectationsWereMet())
}
