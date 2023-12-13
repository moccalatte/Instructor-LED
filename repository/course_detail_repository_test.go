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

func(suite * CourseDetailRepositoryTestSuite)TestCreate(){
	dummy := model.CourseDetail{
		CourseDetailID: "0980909898098234",
		CourseID: "oiuouoiytiuyu",
		CourseChapter: "golang web api",
		CourseContent: "go",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsDeleted: false,
	}
	suite.sqlmock.ExpectBegin()
	rows:=sqlmock.NewRows([]string{"course_detail_id", "course_id", "course_chapter", "course_content", "created_at", "updated_at", "is_deleted"}).AddRow(dummy.CourseDetailID,dummy.CourseID,dummy.CourseChapter,dummy.CourseContent,dummy.CreatedAt, dummy.UpdatedAt, dummy.IsDeleted)
	suite.sqlmock.ExpectQuery("insert into course_detail").WillReturnRows(rows)
	suite.sqlmock.ExpectCommit()

	actual, err := suite.repo.Create(dummy)

	assert.Nil(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), dummy.CourseDetailID, actual.CourseDetailID, "CourseDetailID should match")
	assert.Equal(suite.T(), dummy.CourseID, actual.CourseID, "CourseID should match")
	assert.Equal(suite.T(), dummy.CourseChapter, actual.CourseChapter, "CourseChapter should match")
	assert.Equal(suite.T(), dummy.CourseContent, actual.CourseContent, "CourseContent should match")
	assert.Equal(suite.T(), dummy.CreatedAt, actual.CreatedAt, "CreatedAt should match")
	assert.Equal(suite.T(), dummy.UpdatedAt, actual.UpdatedAt, "UpdatedAt should match")
	assert.Equal(suite.T(), dummy.IsDeleted, actual.IsDeleted, "IsDeleted should match")
}
