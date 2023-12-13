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

type QuestionRepositoryTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	sqlmock sqlmock.Sqlmock
	repo    QuestionRepository
}

func (suite *QuestionRepositoryTestSuite) SetupTest() {
	db, sqlmock, err := sqlmock.New()
	assert.NoError(suite.T(), err)
	suite.mockDb = db
	suite.sqlmock = sqlmock
	suite.repo = NewQuestionRepository(suite.mockDb)
}

func TestQuestionRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(QuestionRepositoryTestSuite))
}

func (suite *QuestionRepositoryTestSuite) TestCreateQuestion() {

	dummy := model.Question{
		QuestionID:  "1212uywewewr2",
		SessionID:   "123d2423",
		StudentID:   "234eedwerwer",
		TrainerID:   "234234fjehe",
		Title:       "go db",
		Description: "error saat debug bagaimana caranya?",
		CourseID:    "4234423423423",
		Answer:      "belum",
		Image:       "yweiuryeiuyriwerwerwerwer.jpg",
		Status:      "belum terjawab",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		IsDeleted:   false,
	}

	suite.sqlmock.ExpectBegin()

	rows := sqlmock.NewRows([]string{"question_id", "session_id", "student_id", "trainer_id", "title", "description", "course_id", "image", "answer", "status", "created_at", "updated_at", "is_deleted"}).
		AddRow(dummy.QuestionID, dummy.SessionID, dummy.StudentID, dummy.TrainerID, dummy.Title, dummy.Description, dummy.CourseID, dummy.Image, dummy.Answer, dummy.Status, dummy.CreatedAt, dummy.UpdatedAt, dummy.IsDeleted)

	suite.sqlmock.ExpectQuery("insert into question").WillReturnRows(rows)
	suite.sqlmock.ExpectCommit()

	actual, err := suite.repo.Create(dummy)
	assert.Nil(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), dummy.QuestionID, actual.QuestionID, "QuestionID should match")
	assert.Equal(suite.T(), dummy.SessionID, actual.SessionID, "SessionID should match")
	assert.Equal(suite.T(), dummy.StudentID, actual.StudentID, "StudentID should match")
	assert.Equal(suite.T(), dummy.TrainerID, actual.TrainerID, "TrainerID should match")
	assert.Equal(suite.T(), dummy.Title, actual.Title, "Title should match")
	assert.Equal(suite.T(), dummy.Description, actual.Description, "Description should match")
	assert.Equal(suite.T(), dummy.CourseID, actual.CourseID, "CourseID should match")
	assert.Equal(suite.T(), dummy.Image, actual.Image, "Image should match")
	assert.Equal(suite.T(), dummy.Answer, actual.Answer, "Answer should match")
	assert.Equal(suite.T(), dummy.Status, actual.Status, "Status should match")
	assert.Equal(suite.T(), dummy.CreatedAt, actual.CreatedAt, "CreatedAt should match")
	assert.Equal(suite.T(), dummy.UpdatedAt, actual.UpdatedAt, "UpdatedAt should match")
	assert.Equal(suite.T(), dummy.IsDeleted, actual.IsDeleted, "IsDeleted should match")
}

