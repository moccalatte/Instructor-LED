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

func (suite *CourseRepositoryTestSuite) TestGetById() {
	dummy := model.Course{
		CourseID:    "7634836923742983742",
		CourseName:  "nananananan",
		Description: "blabla",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		IsDeleted:   false,
	}

	query := "select \\* from course where course_id = \\$1;"
	courseID := "7634836923742983742"

	rows := sqlmock.NewRows([]string{"course_id", "course_name", "description", "created_at", "updated_at", "is_deleted"}).AddRow(
		dummy.CourseID, dummy.CourseName, dummy.Description, dummy.CreatedAt, dummy.UpdatedAt, dummy.IsDeleted,
	)

	suite.sqlmock.ExpectQuery(query).WithArgs(courseID).WillReturnRows(rows)

	actual, err := suite.repo.GetById(courseID)

	assert.Nil(suite.T(), err, "Error should be nill")
	assert.Equal(suite.T(), dummy.CourseID, actual.CourseID, "CourseId")
	assert.Equal(suite.T(), dummy.CourseName, actual.CourseName, "CourseName")
	assert.Equal(suite.T(), dummy.Description, actual.Description, "Description")
	assert.Equal(suite.T(), dummy.CreatedAt, actual.CreatedAt, "CreatedAt should match")
	assert.Equal(suite.T(), dummy.UpdatedAt, actual.UpdatedAt, "UpdatedAt should match")
	assert.Equal(suite.T(), dummy.IsDeleted, actual.IsDeleted, "IsDeleted should match")
}

func(suite *CourseRepositoryTestSuite)TestCreate(){

}

func(suite *CourseRepositoryTestSuite)TestUpdate(){

}

func(suite *CourseRepositoryTestSuite)TestDelete(){

}

func(suite *CourseRepositoryTestSuite)TestFindAll(){
	
}