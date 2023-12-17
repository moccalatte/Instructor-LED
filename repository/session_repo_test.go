package repository

import (
	"database/sql"
	"testing"

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