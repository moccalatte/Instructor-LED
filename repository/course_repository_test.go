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

type CourseRepositoryTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	sqlmock sqlmock.Sqlmock
	repo    CourseRepository
}

func (suite *CourseRepositoryTestSuite) SetupTest() {
	db, sqlmock, err := sqlmock.New()
	assert.NoError(suite.T(), err)
	suite.mockDb = db
	suite.sqlmock = sqlmock
	suite.repo = NewCourseRepository(suite.mockDb)
}

func TestCourseRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(CourseRepositoryTestSuite))
}

func (suite *CourseRepositoryTestSuite) TestCreateCourse() {
	dummy := model.Course{
		CourseID:    "435687868",
		CourseName:  "go db ",
		Description: "go db bla bla la",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		IsDeleted:   false,
	}

	suite.sqlmock.ExpectBegin()
	rows := sqlmock.NewRows([]string{"course_id", "course_name", "description", "created_at", "updated_at", "is_deleted"}).
		AddRow(dummy.CourseID, dummy.CourseName, dummy.Description, dummy.CreatedAt, dummy.UpdatedAt, dummy.IsDeleted)
	suite.sqlmock.ExpectQuery("insert into course").WillReturnRows(rows)
	suite.sqlmock.ExpectCommit()


	actual, err := suite.repo.Create(dummy)

	
	assert.Nil(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), dummy.CourseID, actual.CourseID, "CourseID should match")
	assert.Equal(suite.T(), dummy.CourseName, actual.CourseName, "CourseName should match")
	assert.Equal(suite.T(), dummy.Description, actual.Description, "Description should match")
	assert.Equal(suite.T(), dummy.CreatedAt, actual.CreatedAt, "CreatedAt should match")
	assert.Equal(suite.T(), dummy.UpdatedAt, actual.UpdatedAt, "UpdatedAt should match")
	assert.Equal(suite.T(), dummy.IsDeleted, actual.IsDeleted, "IsDeleted should match")
}
