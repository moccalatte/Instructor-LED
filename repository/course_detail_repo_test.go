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

type CourseDetailRepositoryTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	sqlmock sqlmock.Sqlmock
	repo    CourseDetailRepository
}

func (suite *CourseDetailRepositoryTestSuite) SetupTest() {
	db, sqlmock, err := sqlmock.New()
	assert.NoError(suite.T(), err)
	suite.mockDb = db
	suite.sqlmock = sqlmock
	suite.repo = NewCourseDetailRepository(suite.mockDb)
}

func TestCourseDetailRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(CourseDetailRepositoryTestSuite))
}

func (suite *CourseDetailRepositoryTestSuite) TestGetById() {
	dummy := model.CourseDetail{
		CourseDetailID: "09809328402934",
		CourseID:       "0980239842342",
		CourseChapter:  "golang database",
		CourseContent:  "Api connect datbase",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		IsDeleted:      false,
	}
	query := "select \\* from course_detail where course_detail_id = \\$1;"
	CourseDetailID := "09809328402934"
	rows := sqlmock.NewRows([]string{"course_detail_id", "course_id", "course_chapter", "course_content", "created_at", "updated_at", "is_deleted"}).AddRow(dummy.CourseDetailID, dummy.CourseID, dummy.CourseChapter, dummy.CourseContent, dummy.CreatedAt, dummy.UpdatedAt, dummy.IsDeleted)

	suite.sqlmock.ExpectQuery(query).WithArgs(CourseDetailID).WillReturnRows(rows)

	actual, err := suite.repo.GetById(CourseDetailID)

	assert.Nil(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), dummy.CourseDetailID, actual.CourseDetailID, "CourseDetailID should match")
	assert.Equal(suite.T(), dummy.CourseID, actual.CourseID, "CourseID should match")
	assert.Equal(suite.T(), dummy.CourseChapter, actual.CourseChapter, "CourseChapter should match")
	assert.Equal(suite.T(), dummy.CourseContent, actual.CourseContent, "CourseContent should match")
	assert.Equal(suite.T(), dummy.CreatedAt, actual.CreatedAt, "CreatedAt should match")
	assert.Equal(suite.T(), dummy.UpdatedAt, actual.UpdatedAt, "UpdatedAt should match")
	assert.Equal(suite.T(), dummy.IsDeleted, actual.IsDeleted, "IsDeleted should match")

}

func (suite *CourseDetailRepositoryTestSuite)TestGetAll(){

}

func (suite *CourseDetailRepositoryTestSuite)TestDelete(){

}

func (suite *CourseDetailRepositoryTestSuite)TestUpdate(){

}
