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
	query := "select \\* from users where user_id = \\$1 returning user_id, fullname, role, email, password, created_at, updated_at, is_deleted;"
	user_id := "80980986875676768798797"

	rows := sqlmock.NewRows([]string{"user_id", "fullname", "role", "email", "password", "created_at", "updated_at", "is_deleted"}).
		AddRow(dummy.UserID, dummy.Fullname, dummy.Role, dummy.Email, dummy.Password, dummy.CreatedAt, dummy.UpdatedAt, dummy.IsDeleted)

	suite.sqlmock.ExpectQuery(query).WithArgs(user_id).WillReturnRows(rows)

	actual, err := suite.repo.GetById(user_id)

	assert.Nil(suite.T(), err, "error should be nil")
	assert.Equal(suite.T(), dummy.UserID, actual.UserID, "UserID should match")
	assert.Equal(suite.T(), dummy.Fullname, actual.Fullname, "Fullname should match")
	assert.Equal(suite.T(), dummy.Role, actual.Role, "Role should match")
	assert.Equal(suite.T(), dummy.Email, actual.Email, "Email should match")
	assert.Equal(suite.T(), dummy.Password, actual.Password, "Password should match")
	assert.WithinDuration(suite.T(), dummy.CreatedAt, actual.CreatedAt, time.Second, "CreatedAt should match")
	assert.WithinDuration(suite.T(), dummy.UpdatedAt, actual.UpdatedAt, time.Second, "UpdatedAt should match")
	assert.Equal(suite.T(), dummy.IsDeleted, actual.IsDeleted, "IsDeleted should match")
}

// func (suite *UserRepositoryTestSuite) TestUpdate() {
// 	dummyPayload := model.Users{
// 		Fullname:  "Parto",
// 		Role:      "boss",
// 		Email:     "guakanbos@gmail.com",
// 		Password:  "22222222",
// 		UpdatedAt: time.Now(),
// 		IsDeleted: false,
// 	}

// 	dummyResult := model.Users{
// 		UserID:    "78237492837492387429387423",
// 		Fullname:  dummyPayload.Fullname,
// 		Role:      dummyPayload.Role,
// 		Email:     dummyPayload.Email,
// 		Password:  dummyPayload.Password,
// 		UpdatedAt: time.Now(),
// 		CreatedAt: time.Now(),
// 		IsDeleted: false,
// 	}
// 	query := "update users set fullname = \\$1,role = \\$2,email = \\$3,password = \\$4, updated_at = \\$5, is_deleted = \\$6 where user_id = \\$7 returning user_id, fullname, role, email, password, created_at, updated_at, is_deleted;"

// 	rows := sqlmock.NewRows([]string{"user_id", "fullname", "role", "email", "password", "created_at", "updated_at", "is_deleted"}).
// 		AddRow(dummyResult.UserID, dummyResult.Fullname, dummyResult.Role, dummyResult.Email, dummyResult.Password, dummyResult.CreatedAt, dummyResult.UpdatedAt, dummyResult.IsDeleted)

// 	suite.sqlmock.ExpectQuery(query).WithArgs(
// 		dummyPayload.Fullname,
// 		dummyPayload.Role,
// 		dummyPayload.Email,
// 		dummyPayload.Password,
// 		dummyPayload.UpdatedAt,
// 		dummyPayload.IsDeleted,
// 		dummyResult.UserID,
// 	).WillReturnRows(rows)

// 	actual, err := suite.repo.Update(dummyPayload, dummyResult.UserID)

// 	assert.Nil(suite.T(), err, "Error should be nil")
// 	assert.Equal(suite.T(), dummyResult.UserID, actual.UserID, "UserID should match")
// 	assert.Equal(suite.T(), dummyPayload.Fullname, actual.Fullname, "Fullname should match")
// 	assert.Equal(suite.T(), dummyPayload.Role, actual.Role, "Role should match")
// 	assert.Equal(suite.T(), dummyPayload.Email, actual.Email, "UserID should match")
// 	assert.Equal(suite.T(), dummyPayload.Password, actual.Password, "Password should match")
// 	assert.Equal(suite.T(), dummyResult.CreatedAt, actual.CreatedAt, "CreatedAt should match")
// 	assert.Equal(suite.T(), dummyPayload.UpdatedAt, actual.UpdatedAt, "UpdatedAt should match")
// 	assert.Equal(suite.T(), dummyPayload.IsDeleted, actual.IsDeleted, "UserID should match")
// }
