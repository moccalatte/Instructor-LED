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

func (suite *CourseDetailRepositoryTestSuite) TestCreate() {
	dummy := model.CourseDetail{
		CourseDetailID: "0980909898098234",
		CourseID:       "oiuouoiytiuyu",
		CourseChapter:  "golang web api",
		CourseContent:  "go",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		IsDeleted:      false,
	}
	suite.sqlmock.ExpectBegin()
	rows := sqlmock.NewRows([]string{"course_detail_id", "course_id", "course_chapter", "course_content", "created_at", "updated_at", "is_deleted"}).AddRow(dummy.CourseDetailID, dummy.CourseID, dummy.CourseChapter, dummy.CourseContent, dummy.CreatedAt, dummy.UpdatedAt, dummy.IsDeleted)
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

// func (suite *CourseDetailRepositoryTestSuite) TestUpdate() {
// 	dummyPayload := model.CourseDetail{
// 		CourseID:      "1234",
// 		CourseChapter: "gogo",
// 		CourseContent: "api",
// 		UpdatedAt:     time.Now(),
// 	}

// 	var dummyResult = model.CourseDetail{
// 		CourseDetailID: "1232333",
// 		CourseID:       dummyPayload.CourseID,
// 		CourseChapter:  dummyPayload.CourseChapter,
// 		CourseContent:  dummyPayload.CourseContent,
// 		CreatedAt:      time.Now(),
// 		UpdatedAt:      dummyPayload.UpdatedAt,
// 		IsDeleted:      false,
// 	}

// 	query := "update course_detail set course_id = \\$1, course_chapter = \\$2, course_content = \\$3, updated_at = \\$4, is_deleted = \\$5 where course_detail_id = \\$6 returning course_detail_id, course_id, course_chapter, course_content, created_at, updated_at, is_deleted;"

// 	rows := sqlmock.NewRows([]string{"course_detail_id", "course_id", "course_chapter", "course_content", "created_at", "updated_at", "is_deleted"}).
// 		AddRow(dummyResult.CourseDetailID, dummyPayload.CourseID, dummyPayload.CourseChapter, dummyPayload.CourseContent, dummyResult.CreatedAt, dummyPayload.UpdatedAt, dummyResult.IsDeleted)
// 	fmt.Println("error")
// 	suite.sqlmock.ExpectQuery(query).WithArgs(
// 		dummyPayload.CourseID,
// 		dummyPayload.CourseChapter,
// 		dummyPayload.CourseContent,
// 		dummyPayload.UpdatedAt,
// 		dummyResult.CourseDetailID,
// 	).WillReturnRows(rows)

// 	actual, err := suite.repo.Update(dummyPayload, dummyResult.CourseDetailID)

// 	assert.Nil(suite.T(), err, "Error should be nill")
// 	assert.Equal(suite.T(), dummyResult.CourseDetailID, actual.CourseDetailID, "CourseDetailId")
// 	assert.Equal(suite.T(), dummyPayload.CourseID, actual.CourseID, "CourseID")
// 	assert.Equal(suite.T(), dummyPayload.CourseChapter, actual.CourseChapter, "CourseChpater")
// 	assert.Equal(suite.T(), dummyPayload.CourseContent, actual.CourseContent, "CourseConten")
// 	assert.Equal(suite.T(), dummyResult.CreatedAt, actual.CreatedAt, "CreatedAt")
// 	assert.Equal(suite.T(), dummyPayload.UpdatedAt, actual.UpdatedAt, "Updated")
// 	assert.Equal(suite.T(), dummyPayload.IsDeleted, actual.IsDeleted, "IsDeleted")

// }
