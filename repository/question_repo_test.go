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

func (suite *QuestionRepositoryTestSuite) TestGetById() {
	dummy := model.Question{
		QuestionID:  "12123123",
		SessionID:   "0980980809",
		StudentID:   "09798798798",
		TrainerID:   "98787987978",
		Title:       "hujan",
		Description: "yuisuydisyifywerwerwerwer",
		CourseID:    "98797343324",
		Image:       "gambarnaruto.jpg",
		Answer:      "",
		Status:      "belum terjawab",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		IsDeleted:   false,
	}
	query := "select \\* from question where question_id = \\$1;"
	questionID := "12123123"

	rows := sqlmock.NewRows([]string{"question_id", "session_id", "student_id", "trainer_id", "title", "description", "course_id", "image", "answer", "status", "created_at", "updated_at", "is_deleted"}).
		AddRow(dummy.QuestionID, dummy.SessionID, dummy.StudentID, dummy.TrainerID, dummy.Title, dummy.Description, dummy.CourseID, dummy.Image, dummy.Answer, dummy.Status, dummy.CreatedAt, dummy.UpdatedAt, dummy.IsDeleted)

	suite.sqlmock.ExpectQuery(query).WithArgs(questionID).WillReturnRows(rows)

	actual, err := suite.repo.GetById(questionID)

	assert.Nil(suite.T(), err, "Error should be nill")
	assert.Equal(suite.T(), dummy.QuestionID, actual.QuestionID, "QuestionID")
	assert.Equal(suite.T(), dummy.SessionID, actual.SessionID, "SessionID")
	assert.Equal(suite.T(), dummy.StudentID, actual.StudentID, "StudentID")
	assert.Equal(suite.T(), dummy.TrainerID, actual.TrainerID, "TrainerID")
	assert.Equal(suite.T(), dummy.Title, actual.Title, "Tittle")
	assert.Equal(suite.T(), dummy.Description, actual.Description, "Description")
	assert.Equal(suite.T(), dummy.CourseID, actual.CourseID, "CourseID")
	assert.Equal(suite.T(), dummy.Image, actual.Image, "Image")
	assert.Equal(suite.T(), dummy.Answer, actual.Answer, "Answer")
	assert.Equal(suite.T(), dummy.Status, actual.Status, "Status")
	assert.Equal(suite.T(), dummy.CreatedAt, actual.CreatedAt, "CreatedAt")
	assert.Equal(suite.T(), dummy.UpdatedAt, actual.UpdatedAt, "UpdatedAt")
	assert.Equal(suite.T(), dummy.IsDeleted, actual.IsDeleted, "IsDeleted")
}


func (suite *QuestionRepositoryTestSuite) TestByStudentId() {
	dummy := model.Question{
		QuestionID:  "12123123",
		SessionID:   "0980980809",
		StudentID:   "09798798798",
		TrainerID:   "98787987978",
		Title:       "hujan",
		Description: "yuisuydisyifywerwerwerwer",
		CourseID:    "98797343324",
		Image:       "gambarnaruto.jpg",
		Answer:      "",
		Status:      "belum terjawab",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		IsDeleted:   false,
	}
	query := "select \\* from question where student_id = \\$1;"
	studentID := "0980980809"

	rows := sqlmock.NewRows([]string{"question_id", "session_id", "student_id", "trainer_id", "title", "description", "course_id", "image", "answer", "status", "created_at", "updated_at", "is_deleted"}).
		AddRow(dummy.QuestionID, dummy.SessionID, dummy.StudentID, dummy.TrainerID, dummy.Title, dummy.Description, dummy.CourseID, dummy.Image, dummy.Answer, dummy.Status, dummy.CreatedAt, dummy.UpdatedAt, dummy.IsDeleted)

	suite.sqlmock.ExpectQuery(query).WithArgs(studentID).WillReturnRows(rows)

	actual, err := suite.repo.GetByStudentId(studentID)

	assert.Nil(suite.T(), err, "Error should be nill")
	assert.Equal(suite.T(), dummy.QuestionID, actual.QuestionID, "QuestionID")
	assert.Equal(suite.T(), dummy.SessionID, actual.SessionID, "SessionID")
	assert.Equal(suite.T(), dummy.StudentID, actual.StudentID, "StudentID")
	assert.Equal(suite.T(), dummy.TrainerID, actual.TrainerID, "TrainerID")
	assert.Equal(suite.T(), dummy.Title, actual.Title, "Tittle")
	assert.Equal(suite.T(), dummy.Description, actual.Description, "Description")
	assert.Equal(suite.T(), dummy.CourseID, actual.CourseID, "CourseID")
	assert.Equal(suite.T(), dummy.Image, actual.Image, "Image")
	assert.Equal(suite.T(), dummy.Answer, actual.Answer, "Answer")
	assert.Equal(suite.T(), dummy.Status, actual.Status, "Status")
	assert.Equal(suite.T(), dummy.CreatedAt, actual.CreatedAt, "CreatedAt")
	assert.Equal(suite.T(), dummy.UpdatedAt, actual.UpdatedAt, "UpdatedAt")
	assert.Equal(suite.T(), dummy.IsDeleted, actual.IsDeleted, "IsDeleted")
}

func (suite *QuestionRepositoryTestSuite) TestGetall() {
	dummResult := []model.Question{
		{
			QuestionID:  "12123123",
			SessionID:   "0980980809",
			StudentID:   "09798798798",
			TrainerID:   "98787987978",
			Title:       "hujan",
			Description: "yuisuydisyifywerwerwerwer",
			CourseID:    "98797343324",
			Image:       "gambarnaruto.jpg",
			Answer:      "",
			Status:      "belum terjawab",
			IsDeleted:   false,
		}, {
			QuestionID:  "12123023",
			SessionID:   "0980980809",
			StudentID:   "09798798798",
			TrainerID:   "98787987978",
			Title:       "hujan",
			Description: "yuisuydisyifywerwerwerwer",
			CourseID:    "98797343324",
			Image:       "gambarnaruto.jpg",
			Answer:      "",
			Status:      "belum terjawab",
			IsDeleted:   false,
		},
	}

	query := "SELECT question_id,session_id, student_id, trainer_id, title, description, course_id, image, answer, status, is_deleted FROM question WHERE is_deleted = false;"

	rows := sqlmock.NewRows([]string{"question_id", "session_id", "student_id", "trainer_id", "title", "description", "course_id", "image", "answer", "status", "is_deleted"})

	for _, result := range dummResult {
		rows.AddRow(
			result.QuestionID,
			result.SessionID,
			result.StudentID,
			result.TrainerID,
			result.Title,
			result.Description,
			result.CourseID,
			result.Image,
			result.Answer,
			result.Status,
			result.IsDeleted,
		)

	}
	suite.sqlmock.ExpectQuery(query).WithArgs().WillReturnRows(rows)
	actual, err := suite.repo.GetAll()

	assert.Nil(suite.T(), err, "Error should be nil")
	assert.Len(suite.T(), actual, len(dummResult), "Number of results should match")

	for i, expected := range dummResult {
		assert.Equal(suite.T(), expected.QuestionID, actual[i].QuestionID)
		assert.Equal(suite.T(), expected.SessionID, actual[i].SessionID)
		assert.Equal(suite.T(), expected.StudentID, actual[i].StudentID)
		assert.Equal(suite.T(), expected.TrainerID, actual[i].TrainerID)
		assert.Equal(suite.T(), expected.Title, actual[i].Title)
		assert.Equal(suite.T(), expected.Description, actual[i].Description)
		assert.Equal(suite.T(), expected.CourseID, actual[i].CourseID)
		assert.Equal(suite.T(), expected.Image, actual[i].Image)
		assert.Equal(suite.T(), expected.Answer, actual[i].Answer)
		assert.Equal(suite.T(), expected.Status, actual[i].Status)
		assert.Equal(suite.T(), expected.IsDeleted, actual[i].IsDeleted)

	}

}



func (suite *QuestionRepositoryTestSuite) TestDelete() {
	dummy := model.Question{
		QuestionID:  "12123123",
		SessionID:   "0980980809",
		StudentID:   "09798798798",
		TrainerID:   "98787987978",
		Title:       "hujan",
		Description: "yuisuydisyifywerwerwerwer",
		CourseID:    "98797343324",
		Image:       "gambarnaruto.jpg",
		Answer:      "iuiooououo",
		Status:      "belum terjawab",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		IsDeleted:   true,
	}
	query := "update question set is_deleted \\= \\$1 where question_id \\= \\$2 returning question_id, session_id, student_id, trainer_id, title, description, course_id, image, answer, status, created_at, updated_at, is_deleted;"

	suite.sqlmock.ExpectBegin()
	rows := sqlmock.NewRows([]string{"question_id", "session_id", "student_id", "trainer_id", "title", "description", "course_id", "image", "answer", "status", "created_at", "updated_at", "is_deleted"}).
		AddRow(
			dummy.QuestionID,
			dummy.SessionID,
			dummy.StudentID,
			dummy.TrainerID,
			dummy.Title,
			dummy.Description,
			dummy.CourseID,
			dummy.Image,
			dummy.Answer,
			dummy.Status,
			dummy.CreatedAt,
			dummy.UpdatedAt,
			dummy.IsDeleted)
	suite.sqlmock.ExpectQuery(query).WithArgs(
		true,
		dummy.QuestionID,
	).WillReturnRows(rows)

	suite.sqlmock.ExpectCommit()

	actual, err := suite.repo.Delete(dummy.QuestionID)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), dummy.QuestionID, actual.QuestionID)
	assert.Equal(suite.T(), dummy.SessionID, actual.SessionID, "SessionID")
	assert.Equal(suite.T(), dummy.StudentID, actual.StudentID, "StudentID")
	assert.Equal(suite.T(), dummy.TrainerID, actual.TrainerID, "TrainerID")
	assert.Equal(suite.T(), dummy.Title, actual.Title, "Tittle")
	assert.Equal(suite.T(), dummy.Description, actual.Description, "Description")
	assert.Equal(suite.T(), dummy.CourseID, actual.CourseID, "CourseID")
	assert.Equal(suite.T(), dummy.Image, actual.Image, "Image")
	assert.Equal(suite.T(), dummy.Answer, actual.Answer, "Answer")
	assert.Equal(suite.T(), dummy.Status, actual.Status, "Status")
	assert.Equal(suite.T(), dummy.CreatedAt, actual.CreatedAt, "CreatedAt")
	assert.Equal(suite.T(), dummy.UpdatedAt, actual.UpdatedAt, "UpdatedAt")
	assert.Equal(suite.T(), dummy.IsDeleted, actual.IsDeleted, "IsDeleted")

}

