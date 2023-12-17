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
	assert.Equal(suite.T(), dummy.Title, actual.Title)
	assert.Equal(suite.T(), dummy.Description, actual.Description)
	assert.Equal(suite.T(), dummy.SessionDate, actual.SessionDate)
	assert.Equal(suite.T(), dummy.SessionTime, actual.SessionTime)
	assert.Equal(suite.T(), dummy.SessionLink, actual.SessionLink)
	assert.Equal(suite.T(), dummy.TrainerID, actual.TrainerID)
	assert.Equal(suite.T(), dummy.CreatedAt, actual.CreatedAt)
	assert.Equal(suite.T(), dummy.UpdatedAt, actual.UpdatedAt)
	assert.Equal(suite.T(), dummy.IsDeleted, actual.IsDeleted)
}

func(suite *SessionRepositoryTestSuite)TestCreate(){

}

func(suite *SessionRepositoryTestSuite)TestGetById(){

}
func(suite *SessionRepositoryTestSuite)TestGetAllSession(){

}
func(suite *SessionRepositoryTestSuite)TestUpdateQ(){

}

func(suite *SessionRepositoryTestSuite)TestUpdateNote(){

}
func(suite *SessionRepositoryTestSuite)TestDelete(){

}

func(suite *SessionRepositoryTestSuite)TestFindAll(){
	
}