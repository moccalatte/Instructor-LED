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

func (suite *CourseRepositoryTestSuite) TestUpdateCourse() {
	dummyPayload := model.Course{
		CourseName:  "go db ",
		Description: "go db bla bla la",
		UpdatedAt:   time.Now(),
		IsDeleted:   false,
	}
	dummyResult := model.Course{
		CourseID:    "435687868",
		CourseName:  dummyPayload.CourseName,
		Description: dummyPayload.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   dummyPayload.UpdatedAt,
		IsDeleted:   false,
	}

	suite.sqlmock.ExpectBegin()
	query := "update course set course_name = \\$1, description = \\$2, updated_at = \\$3, is_deleted = \\$4 where course_id = \\$5 returning  course_id, course_name, description, created_at, updated_at, is_deleted;"
	rows := sqlmock.NewRows([]string{"course_id", "course_name", "description", "created_at", "updated_at", "is_deleted"}).
		AddRow(dummyResult.CourseID, dummyResult.CourseName, dummyResult.Description, dummyResult.CreatedAt, dummyResult.UpdatedAt, dummyResult.IsDeleted)
	suite.sqlmock.ExpectQuery(query).WillReturnRows(rows)
	suite.sqlmock.ExpectCommit()

	actual, err := suite.repo.Update(dummyPayload,dummyResult.CourseID)

	assert.Nil(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), dummyResult.CourseID, actual.CourseID, "CourseID should match")
	assert.Equal(suite.T(), dummyPayload.CourseName, actual.CourseName, "CourseName should match")
	assert.Equal(suite.T(), dummyPayload.Description, actual.Description, "Description should match")
	assert.Equal(suite.T(), dummyResult.CreatedAt, actual.CreatedAt, "CreatedAt should match")
	assert.Equal(suite.T(), dummyPayload.UpdatedAt, actual.UpdatedAt, "UpdatedAt should match")
	assert.Equal(suite.T(), dummyPayload.IsDeleted, actual.IsDeleted, "IsDeleted should match")
}

func (suite *CourseRepositoryTestSuite) TestDelete() {
	dummyPayload := model.Course{
		IsDeleted: true,
	}
	dummyResult := model.Course{
		CourseID:    "435687868",
		CourseName:  dummyPayload.CourseName,
		Description: dummyPayload.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		IsDeleted:   dummyPayload.IsDeleted,
	}

	suite.sqlmock.ExpectBegin()
	query := "update course set is_deleted = \\$1 where course_id = \\$2 returning  course_id, course_name, description, created_at, updated_at, is_deleted;"

	rows := sqlmock.NewRows([]string{"course_id", "course_name", "description", "created_at", "updated_at", "is_deleted"}).
		AddRow(dummyResult.CourseID, dummyResult.CourseName, dummyResult.Description, dummyResult.CreatedAt, dummyResult.UpdatedAt, dummyResult.IsDeleted)

	suite.sqlmock.ExpectQuery(query).
		WithArgs(dummyPayload.IsDeleted, dummyResult.CourseID).
		WillReturnRows(rows)

	suite.sqlmock.ExpectCommit()

	actual, err := suite.repo.Delete(dummyResult.CourseID)

	assert.Nil(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), dummyResult.CourseID, actual.CourseID, "CourseID should match")
	assert.Equal(suite.T(), dummyResult.CourseName, actual.CourseName, "CourseName should match")
	assert.Equal(suite.T(), dummyResult.Description, actual.Description, "Description should match")
	assert.Equal(suite.T(), dummyResult.CreatedAt, actual.CreatedAt, "CreatedAt should match")
	assert.Equal(suite.T(), dummyResult.UpdatedAt, actual.UpdatedAt, "UpdatedAt should match")
	assert.Equal(suite.T(), dummyPayload.IsDeleted, actual.IsDeleted, "IsDeleted should match")
}

