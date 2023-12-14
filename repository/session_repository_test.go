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

type SessionRepositoryTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	sqlmock sqlmock.Sqlmock
	repo    SessionRepository
}

func (suite *SessionRepositoryTestSuite) SetupTest() {
	db, sqlmock, err := sqlmock.New()
	assert.NoError(suite.T(), err)
	suite.mockDb = db
	suite.sqlmock = sqlmock
	suite.repo = NewSessionRepository(suite.mockDb)
}
func TestSessionRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(SessionRepositoryTestSuite))
}

func (suite *SessionRepositoryTestSuite) TestCreateSession_Success() {
	dummy := model.Session{
		Title:       "Fundamental GO",
		Description: "ini adalah awal dari bahasa go",
		SessionDate: "2023-12-18",
		SessionTime: "18:00 WIB",
		SessionLink: "https://gogogogogogo.com",
		TrainerID:   "098708767t6087f760g760g8760g",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now().Add(26 * 24 * time.Hour),
		IsDeleted:   false,
	}

	suite.sqlmock.ExpectBegin()

	rows := sqlmock.NewRows([]string{"session_id", "title", "description", "session_date", "session_time", "session_link", "trainer_id", "created_at", "updated_at", "is_deleted"}).
		AddRow(dummy.SessionID, dummy.Title, dummy.Description, dummy.SessionDate, dummy.SessionTime, dummy.SessionLink, dummy.TrainerID, dummy.CreatedAt, dummy.UpdatedAt, dummy.IsDeleted)
	suite.sqlmock.ExpectQuery("insert into session").WillReturnRows(rows)
	suite.sqlmock.ExpectCommit()

	actual, err := suite.repo.Create(dummy)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), dummy.SessionID, actual.SessionID)
}

func (suite *SessionRepositoryTestSuite) TestGetById() {
	dummy := model.Session{
		SessionID:   "qwerty",
		Title:       "Goo",
		Description: "Go adalaha .....",
		SessionDate: "2023-12-10",
		SessionTime: "18:00",
		SessionLink: "https://auth/lop",
		TrainerID:   "0980986867",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		IsDeleted:   false,
	}
	query := "select \\* from session where session_id = \\$1;"
	sessionID := "qwerty"

	rows := sqlmock.NewRows([]string{"session_id", "title", "description", "session_date", "session_time", "session_link", "trainer_id", "created_at", "updated_at", "is_deleted"}).AddRow(
		dummy.SessionID,
		dummy.Title,
		dummy.Description,
		dummy.SessionDate,
		dummy.SessionTime,
		dummy.SessionLink,
		dummy.TrainerID,
		dummy.CreatedAt,
		dummy.UpdatedAt,
		dummy.IsDeleted,
	)
	suite.sqlmock.ExpectQuery(query).WithArgs(sessionID).WillReturnRows(rows)

	actual, err := suite.repo.GetById(sessionID)

	assert.Nil(suite.T(), err, "Error should be nill")
	assert.Equal(suite.T(), dummy.SessionID, actual.SessionID, "SessionID should match")
	assert.Equal(suite.T(), dummy.Title, actual.Title, "Title should match")
	assert.Equal(suite.T(), dummy.Description, actual.Description, "Description should match")
	assert.Equal(suite.T(), dummy.SessionDate, actual.SessionDate, "SessionDate should match")
	assert.Equal(suite.T(), dummy.SessionTime, actual.SessionTime, "SessionTime should match")
	assert.Equal(suite.T(), dummy.SessionLink, actual.SessionLink, "SessoinLink should match")
	assert.Equal(suite.T(),dummy.TrainerID,actual.TrainerID,"TrinerID should match")
	assert.Equal(suite.T(),dummy.CreatedAt,actual.CreatedAt,"CreatedAt should match")
	assert.Equal(suite.T(),dummy.UpdatedAt,actual.UpdatedAt,"UpdatedAt should match")
	assert.Equal(suite.T(),dummy.IsDeleted,actual.IsDeleted,"Title should match")
}
