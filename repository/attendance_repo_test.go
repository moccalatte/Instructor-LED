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

type AttendanceRepositoryTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	sqlmock sqlmock.Sqlmock
	repo    AttendanceRepository
}

func (suite *AttendanceRepositoryTestSuite) SetupTest() {
	db, sqlmock, err := sqlmock.New()
	assert.NoError(suite.T(), err)
	suite.mockDb = db
	suite.sqlmock = sqlmock
	suite.repo = NewAttendanceRepository(suite.mockDb)
}

func TestAttendanceRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(AttendanceRepositoryTestSuite))
}

func (suite *AttendanceRepositoryTestSuite) TestGetById() {

	dummy := model.Attendance{
		AttendanceID:      "098709687657645654567586778",
		SessionID:         "0980687gtr65fr65r",
		StudentID:         "098098gv6vtrfr6d6r",
		AttendanceStudent: true,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
		IsDeleted:         false,
	}

	query := "select \\* from attendance where attendance_id = \\$1;"
	attendanceID := "098709687657645654567586778"

	rows := sqlmock.NewRows([]string{"attendance_id", "session_id", "student_id", "attendance_student", "created_at", "updated_at", "is_deleted"}).
		AddRow(dummy.AttendanceID, dummy.SessionID, dummy.StudentID, dummy.AttendanceStudent, dummy.CreatedAt, dummy.UpdatedAt, dummy.IsDeleted)

	suite.sqlmock.ExpectQuery(query).WithArgs(attendanceID).WillReturnRows(rows)

	actual, err := suite.repo.GetById(attendanceID)

	assert.Nil(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), dummy.AttendanceID, actual.AttendanceID, "AttendanceID should match")
	assert.Equal(suite.T(), dummy.SessionID, actual.SessionID, "SessionID should match")
	assert.Equal(suite.T(), dummy.StudentID, actual.StudentID, "StudentID should match")
	assert.Equal(suite.T(), dummy.AttendanceID, actual.AttendanceID, "AttendanceID should match")
	assert.Equal(suite.T(), dummy.CreatedAt, actual.CreatedAt, "CreatedAt should match")
	assert.Equal(suite.T(), dummy.UpdatedAt, actual.UpdatedAt, "UpdatedAt should match")
	assert.Equal(suite.T(), dummy.IsDeleted, actual.IsDeleted, "IsDeleted should match")
}

func (suite *AttendanceRepositoryTestSuite) TestGetAll() {
	dummyResult := []model.Attendance{
		{
			AttendanceID:      "098709687657645654567586778",
			SessionID:         "0980687gtr65fr65r",
			StudentID:         "098098gv6vtrfr6d6r",
			AttendanceStudent: true,
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
			IsDeleted:         false,
		},
		{
			AttendanceID:      "098709687657645654567586778",
			SessionID:         "0980687gtr65fr65r",
			StudentID:         "098098gv6vtrfr6d6r",
			AttendanceStudent: true,
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
			IsDeleted:         false,
		},
	}

	query := "select \\* from course_detail where is_deleted = \\$1;"

	rows := sqlmock.NewRows([]string{"attendance_id", "session_id", "student_id", "attendance_student", "created_at", "updated_at", "is_deleted"})

	for _, result := range dummyResult {
		rows.AddRow(
			result.AttendanceID,
			result.SessionID,
			result.StudentID,
			result.AttendanceStudent,
			result.CreatedAt,
			result.UpdatedAt,
			result.IsDeleted,
		)
	}
	suite.sqlmock.ExpectQuery(query).WithArgs().WillReturnRows(rows)

	actual, err := suite.repo.GetAll()

	assert.Nil(suite.T(), err)
	assert.Len(suite.T(), actual, len(dummyResult))
}

func (suite *AttendanceRepositoryTestSuite) TestGetBySessionID() {
	dummy := model.Attendance{
		AttendanceID:      "098709687657645654567586778",
		SessionID:         "0980687gtr65fr65r",
		StudentID:         "098098gv6vtrfr6d6r",
		AttendanceStudent: true,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
		IsDeleted:         false,
	}

	query := "select \\* from attendance where session_id = \\$1;"
	SessionID := "0980687gtr65fr65r"

	rows := sqlmock.NewRows([]string{"attendance_id", "session_id", "student_id", "attendance_student", "created_at", "updated_at", "is_deleted"}).
		AddRow(dummy.AttendanceID, dummy.SessionID, dummy.StudentID, dummy.AttendanceStudent, dummy.CreatedAt, dummy.UpdatedAt, dummy.IsDeleted)

	suite.sqlmock.ExpectQuery(query).WithArgs(SessionID).WillReturnRows(rows)

	actual, err := suite.repo.GetBySessionId(SessionID)

	assert.Nil(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), dummy.AttendanceID, actual.AttendanceID, "AttendanceID should match")
	assert.Equal(suite.T(), dummy.SessionID, actual.SessionID, "SessionID should match")
	assert.Equal(suite.T(), dummy.StudentID, actual.StudentID, "StudentID should match")
	assert.Equal(suite.T(), dummy.AttendanceID, actual.AttendanceID, "AttendanceID should match")
	assert.Equal(suite.T(), dummy.CreatedAt, actual.CreatedAt, "CreatedAt should match")
	assert.Equal(suite.T(), dummy.UpdatedAt, actual.UpdatedAt, "UpdatedAt should match")
	assert.Equal(suite.T(), dummy.IsDeleted, actual.IsDeleted, "IsDeleted should match")
}
func (suite *AttendanceRepositoryTestSuite) TestUpdate() {

}
func (suite *AttendanceRepositoryTestSuite) TestDelete() {

}
func (suite *AttendanceRepositoryTestSuite) TestGetlAll() {

}

func TestCreateAttendance_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	repo := NewAttendanceRepository(db)

	payload := model.Attendance{
		SessionID:         "session123",
		StudentID:         "student123",
		AttendanceStudent: true,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
		IsDeleted:         false,
	}

	expectedResult := model.Attendance{
		AttendanceID:      "attendance123",
		SessionID:         payload.SessionID,
		StudentID:         payload.StudentID,
		AttendanceStudent: payload.AttendanceStudent,
		CreatedAt:         payload.CreatedAt,
		UpdatedAt:         payload.UpdatedAt,
		IsDeleted:         payload.IsDeleted,
	}

	mock.ExpectBegin()

	rows := sqlmock.NewRows([]string{
		"attendance_id", "session_id", "student_id", "attendance_student",
		"created_at", "updated_at", "is_deleted",
	}).AddRow(
		expectedResult.AttendanceID, expectedResult.SessionID, expectedResult.StudentID,
		expectedResult.AttendanceStudent, expectedResult.CreatedAt, expectedResult.UpdatedAt,
		expectedResult.IsDeleted,
	)

	mock.ExpectQuery("insert into attendance").WithArgs(
		payload.SessionID, payload.StudentID, payload.AttendanceStudent, sqlmock.AnyArg(), payload.IsDeleted,
	).WillReturnRows(rows)

	mock.ExpectCommit()

	result, err := repo.Create(payload)

	assert.NoError(t, err, "Error should be nil")
	assert.Equal(t, expectedResult, result, "Objects should match")

	assert.NoError(t, mock.ExpectationsWereMet(), "Not all expectations were met")
}
