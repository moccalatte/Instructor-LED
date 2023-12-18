package repository

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CsvRepositoryTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	sqlmock sqlmock.Sqlmock
	repo    CsvRepository
}

func (suite *CsvRepositoryTestSuite) SetupTest() {
	db, sqlmock, err := sqlmock.New()
	assert.NoError(suite.T(), err)
	suite.mockDb = db
	suite.sqlmock = sqlmock
	suite.repo = NewCsv(suite.mockDb)
}

func TestCsvRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(CsvRepositoryTestSuite))
}

func (c *CsvRepositoryTestSuite) TestCsvStart() {

	expectedSessionIDs := []string{"sessionID1", "sessionID2"}

	c.sqlmock.ExpectQuery("SELECT session_id FROM session").WillReturnRows(sqlmock.NewRows([]string{"sessionId"}).AddRow(expectedSessionIDs[0]).AddRow(expectedSessionIDs[1]))


	result, err := c.repo.CsvStart()

	
	assert.NoError(c.T(), err)
	assert.Equal(c.T(), expectedSessionIDs, result)

	assert.NoError(c.T(), c.sqlmock.ExpectationsWereMet())
}


