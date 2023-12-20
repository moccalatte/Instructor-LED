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

func (suite *SessionRepositoryTestSuite) TestGetById() {
	dummy := model.Session{
		SessionID:   "123123123",
		Title:       "Fundamental GO",
		Description: "ini adalah awal dari bahasa go",
		SessionDate: "2023-12-18",
		SessionTime: "18:00 WIB",
		SessionLink: "https://gogogogogogo.com",
		TrainerID:   "098708767t6087f760g760g8760g",
		Note:        "okokokok",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now().Add(26 * 24 * time.Hour),
		IsDeleted:   false,
	}

	query := "select \\* from session where session_id = \\$1;"
	sessionID := "123123123"

	rows := sqlmock.NewRows([]string{"session_id", "title", "description", "session_date", "session_time", "session_link", "trainer_id", "note", "created_at", "updated_at", "is_deleted"}).AddRow(
		dummy.SessionID,
		dummy.Title,
		dummy.Description,
		dummy.SessionDate,
		dummy.SessionTime,
		dummy.SessionLink,
		dummy.TrainerID,
		dummy.Note,
		dummy.CreatedAt,
		dummy.UpdatedAt,
		dummy.IsDeleted,
	)

	suite.sqlmock.ExpectQuery(query).WithArgs(sessionID).WillReturnRows(rows)

	actual, err := suite.repo.GetById(sessionID)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), dummy.SessionID, actual.SessionID)
	assert.Equal(suite.T(), dummy.Title, actual.Title)
	assert.Equal(suite.T(), dummy.Description, actual.Description)
	assert.Equal(suite.T(), dummy.SessionDate, actual.SessionDate)
	assert.Equal(suite.T(), dummy.SessionTime, actual.SessionTime)
	assert.Equal(suite.T(), dummy.SessionLink, actual.SessionLink)
	assert.Equal(suite.T(), dummy.TrainerID, actual.TrainerID)
	assert.Equal(suite.T(), dummy.Note, actual.Note)
	assert.Equal(suite.T(), dummy.CreatedAt, actual.CreatedAt)
	assert.Equal(suite.T(), dummy.UpdatedAt, actual.UpdatedAt)
	assert.Equal(suite.T(), dummy.IsDeleted, actual.IsDeleted)

}

func (suite *SessionRepositoryTestSuite) TestDelete() {
	dummy := model.Session{
		SessionID:   "123123123",
		Title:       "Fundamental GO",
		Description: "ini adalah awal dari bahasa go",
		SessionDate: "2023-12-18",
		SessionTime: "18:00 WIB",
		SessionLink: "https://gogogogogogo.com",
		TrainerID:   "098708767t6087f760g760g8760g",
		Note:        "okokokok",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now().Add(26 * 24 * time.Hour),
		IsDeleted:   true,
	}

	query := "update session set is_deleted = \\$1 where session_id = \\$2 returning session_id, title, description, session_date, session_time, session_link, trainer_id, note,created_at, updated_at, is_deleted;"

	suite.sqlmock.ExpectBegin()
	rows := sqlmock.NewRows([]string{"session_id", "title", "description", "session_date", "session_time", "session_link", "trainer_id", "note", "created_at", "updated_at", "is_deleted"}).AddRow(
		dummy.SessionID,
		dummy.Title,
		dummy.Description,
		dummy.SessionDate,
		dummy.SessionTime,
		dummy.SessionLink,
		dummy.TrainerID,
		dummy.Note,
		dummy.CreatedAt,
		dummy.UpdatedAt,
		dummy.IsDeleted,
	)
	suite.sqlmock.ExpectQuery(query).WithArgs(
		true,
		dummy.SessionID,
	).WillReturnRows(rows)

	suite.sqlmock.ExpectCommit()

	actual, err := suite.repo.Delete(dummy.SessionID)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), dummy.SessionID, actual.SessionID)
	assert.Equal(suite.T(), dummy.Title, actual.Title)
	assert.Equal(suite.T(), dummy.Description, actual.Description)
	assert.Equal(suite.T(), dummy.SessionDate, actual.SessionDate)
	assert.Equal(suite.T(), dummy.SessionTime, actual.SessionTime)
	assert.Equal(suite.T(), dummy.SessionLink, actual.SessionLink)
	assert.Equal(suite.T(), dummy.TrainerID, actual.TrainerID)
	assert.Equal(suite.T(), dummy.Note, actual.Note)
	assert.Equal(suite.T(), dummy.CreatedAt, actual.CreatedAt)
	assert.Equal(suite.T(), dummy.UpdatedAt, actual.UpdatedAt)
	assert.Equal(suite.T(), dummy.IsDeleted, actual.IsDeleted)

}

func (suite *SessionRepositoryTestSuite) TestGetAllSession() {
	dummyResults := []model.Session{
		{
			SessionID:   "098092384092834",
			Title:       "Fundamental GO",
			Description: "ini adalah awal dari bahasa go",
			SessionDate: "2023-12-18",
			SessionTime: "18:00 WIB",
			SessionLink: "https://gogogogogogo.com",
			TrainerID:   "098708767t6087f760g760g8760g",
			Note:        "okokokok",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now().Add(26 * 24 * time.Hour),
			IsDeleted:   false,
		},
		{
			SessionID:   "987982340238423423",
			Title:       "Fundamental GO",
			Description: "ini adalah awal dari bahasa go",
			SessionDate: "2023-12-18",
			SessionTime: "18:00 WIB",
			SessionLink: "https://gogogogogogo.com",
			TrainerID:   "098708767t6087f760g760g8760g",
			Note:        "okokokok",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now().Add(26 * 24 * time.Hour),
			IsDeleted:   false,
		},
	}

	query := `SELECT session_id, title, description, session_date, session_time, session_link, trainer_id, note,is_deleted FROM session WHERE is_deleted = false ORDER BY session_date ASC;`

	rows := sqlmock.NewRows([]string{"session_id", "title", "description", "session_date", "session_time", "session_link", "trainer_id", "note", "is_deleted"})

	for _, result := range dummyResults {
		rows.AddRow(result.SessionID, result.Title, result.Description, result.SessionDate, result.SessionTime, result.SessionLink, result.TrainerID, result.Note, result.IsDeleted)
	}

	suite.sqlmock.ExpectQuery(query).
		WithArgs().
		WillReturnRows(rows)

	actual, err := suite.repo.GetAllSession()

	assert.Nil(suite.T(), err, "Error should be nil")
	assert.Len(suite.T(), actual, len(dummyResults), "Number of results should match")

	for i, expected := range dummyResults {
		assert.Equal(suite.T(), expected.SessionID, actual[i].SessionID, "SessionID should match")
		assert.Equal(suite.T(), expected.Title, actual[i].Title, "Tittle should match")
		assert.Equal(suite.T(), expected.Description, actual[i].Description, "Description should match")
		assert.Equal(suite.T(), expected.SessionDate, actual[i].SessionDate, "SessionDate should match")
		assert.Equal(suite.T(), expected.SessionTime, actual[i].SessionTime, "SessionTime should match")
		assert.Equal(suite.T(), expected.SessionLink, actual[i].SessionLink, "SessionLink should match")
		assert.Equal(suite.T(), expected.TrainerID, actual[i].TrainerID, "TrainerID should match")
		assert.Equal(suite.T(), expected.Note, actual[i].Note, "IsDeleted should match")
		assert.Equal(suite.T(), expected.IsDeleted, actual[i].IsDeleted, "IsDeleted should match")
	}
}

// func (suite *SessionRepositoryTestSuite) TestUpdateQ() {
// 	dummy := model.Session{
// 		SessionID:   "123123123",
// 		Title:       "Fundamental GO",
// 		Description: "ini adalah awal dari bahasa go",
// 		SessionDate: "2023-12-18",
// 		SessionTime: "18:00 WIB",
// 		SessionLink: "https://gogogogogogo.com",
// 		TrainerID:   "098708767t6087f760g760g8760g",
// 		Note:        "okokokok",
// 		CreatedAt:   time.Now(),
// 		UpdatedAt:   time.Now().Add(26 * 24 * time.Hour),
// 		IsDeleted:   false,
// 	}

// 	query := `update session set title = \\$1, description = \\$2, session_date = \\$3, session_time = \\$4, session_link = \\$5, trainer_id = \\$6,  note = \\$7,updated_at = \\$8, is_deleted = \\$9 where session_id = \\$10 returning session_id, title, description, session_date, session_time, session_link, trainer_id, note,created_at, updated_at, is_deleted;`

// 	suite.sqlmock.ExpectBegin()

// 	rows := sqlmock.NewRows([]string{"session_id", "title", "description", "session_date", "session_time", "session_link", "trainer_id", "note", "created_at", "updated_at", "is_deleted"}).AddRow(
// 		dummy.SessionID,
// 		dummy.Title,
// 		dummy.Description,
// 		dummy.SessionDate,
// 		dummy.SessionTime,
// 		dummy.SessionLink,
// 		dummy.TrainerID,
// 		dummy.Note,
// 		dummy.CreatedAt,
// 		dummy.UpdatedAt,
// 		dummy.IsDeleted,
// 	)
// 	suite.sqlmock.ExpectCommit()
// 	suite.sqlmock.ExpectQuery(query).WithArgs(
// 		dummy.Title,
// 		dummy.Description,
// 		dummy.SessionDate,
// 		dummy.SessionTime,
// 		dummy.SessionLink,
// 		dummy.TrainerID,
// 		dummy.Note,
// 		dummy.CreatedAt,
// 		dummy.UpdatedAt,
// 		dummy.IsDeleted,
// 		dummy.SessionID,
// 	).WillReturnRows(rows)

// 	actual, err := suite.repo.Update(dummy, dummy.SessionID)
// 	fmt.Println(actual)
// 	fmt.Println(dummy)

// 	assert.Nil(suite.T(), err)
// 	assert.NoError(suite.T(), err)
// 	assert.Equal(suite.T(), dummy.SessionID, actual.SessionID)
// 	assert.Equal(suite.T(), dummy.Title, actual.Title)
// 	assert.Equal(suite.T(), dummy.Description, actual.Description)
// 	assert.Equal(suite.T(), dummy.SessionDate, actual.SessionDate)
// 	assert.Equal(suite.T(), dummy.SessionTime, actual.SessionTime)
// 	assert.Equal(suite.T(), dummy.SessionLink, actual.SessionLink)
// 	assert.Equal(suite.T(), dummy.TrainerID, actual.TrainerID)
// 	assert.Equal(suite.T(), dummy.Note, actual.Note)
// 	assert.Equal(suite.T(), dummy.CreatedAt, actual.CreatedAt)
// 	assert.Equal(suite.T(), dummy.UpdatedAt, actual.UpdatedAt)
// 	assert.Equal(suite.T(), dummy.IsDeleted, actual.IsDeleted)

// }
func (suite *SessionRepositoryTestSuite) TestUpdateNote() {

}

func (suite *SessionRepositoryTestSuite) TestCreate() {

}
