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
  
  func (suite *UserRepositoryTestSuite) TestCreate() {
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
	  AddRow(
		dummy.UserID, 
		dummy.Fullname, 
		dummy.Role, 
		dummy.Email, 
		dummy.Password, 
		dummy.CreatedAt, 
		dummy.UpdatedAt, 
		dummy.IsDeleted)
	suite.sqlmock.ExpectQuery("insert into users").
	WithArgs(dummy.Fullname, 
		dummy.Role, dummy.Email, dummy.Password, sqlmock.AnyArg(), sqlmock.AnyArg(), dummy.IsDeleted).
	  WillReturnRows(rows)
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

func (suite *UserRepositoryTestSuite) TestGetAll() {
	dummyUsers := []model.Users{
		{
		  UserID:    "1",
		  Fullname:  "John Doe",
		  Role:      "Admin",
		  Email:     "john.doe@example.com",
		  Password:  "hashed_password_1",
		  IsDeleted: false,
		},
		{
		  UserID:    "2",
		  Fullname:  "Jane Doe",
		  Role:      "User",
		  Email:     "jane.doe@example.com",
		  Password:  "hashed_password_2",
		  IsDeleted: false,
		},
	  }
	  rows := sqlmock.NewRows([]string{"user_id", "fullname", "role", "email", "password", "is_deleted"}).
    AddRow(dummyUsers[0].UserID, dummyUsers[0].Fullname, dummyUsers[0].Role, dummyUsers[0].Email, dummyUsers[0].Password, dummyUsers[0].IsDeleted).
    AddRow(dummyUsers[1].UserID, dummyUsers[1].Fullname, dummyUsers[1].Role, dummyUsers[1].Email, dummyUsers[1].Password, dummyUsers[1].IsDeleted)

  suite.sqlmock.ExpectQuery("SELECT user_id, fullname, role, email, password, is_deleted FROM users WHERE is_deleted \\= false ORDER BY fullname ASC").WillReturnRows(rows)

  actualUsers, err := suite.repo.GetAll()
  assert.Nil(suite.T(), err)
  assert.NoError(suite.T(), err)
  assert.Equal(suite.T(), dummyUsers, actualUsers)
}

func (suite *UserRepositoryTestSuite) TestUpdate() {
	dummy := model.Users{
		UserID:    "sdsdsdsdsdsdsd",
		Fullname:  "Joko Santoso",
		Role:      "admin",
		Email:     "chril@example.com",
		Password:  "230104",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsDeleted: false,
	  }
	
	  query := "update users set fullname \\= \\$1,role \\= \\$2,email \\= \\$3,password \\= \\$4, updated_at \\= \\$5, is_deleted \\= \\$6 where user_id \\= \\$7 returning user_id, fullname, role, email, password, created_at, updated_at, is_deleted"
	  suite.sqlmock.ExpectBegin()
	
	  rows := sqlmock.NewRows([]string{"user_id", "fullname", "role", "email", "password", "created_at", "updated_at", "is_deleted"}).
		AddRow(dummy.UserID, dummy.Fullname, dummy.Role, dummy.Email, dummy.Password, dummy.CreatedAt, dummy.UpdatedAt, dummy.IsDeleted)
	  suite.sqlmock.ExpectQuery(query).WithArgs(dummy.Fullname, dummy.Role, dummy.Email, dummy.Password, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnRows(rows)
	  suite.sqlmock.ExpectCommit()
	
	  actual, err := suite.repo.Update(dummy, dummy.UserID)
	  assert.Nil(suite.T(), err)
	  assert.NoError(suite.T(), err)
	  assert.Equal(suite.T(), dummy, actual)
}

func (suite *UserRepositoryTestSuite) TestDelete() {
	dummy := model.Users{
		UserID:    "sdsdsdsdsdsdsd",
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
		AddRow(dummy.UserID, dummy.Fullname, dummy.Role, dummy.Email, dummy.Password, dummy.CreatedAt, dummy.UpdatedAt, dummy.IsDeleted)
	  suite.sqlmock.ExpectQuery("update users set is_deleted \\= \\$1 where user_id \\= \\$2 returning user_id, fullname, role, email, password, created_at, updated_at, is_deleted").WithArgs(true, dummy.UserID).
		WillReturnRows(rows)
	  suite.sqlmock.ExpectCommit()
	
	  actual, err := suite.repo.Delete(dummy.UserID)
	  assert.Nil(suite.T(), err)
	  assert.NoError(suite.T(), err)
	  assert.Equal(suite.T(), dummy, actual)
	}


