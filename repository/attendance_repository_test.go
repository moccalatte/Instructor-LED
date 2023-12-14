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
	suite.Run(t, new(CourseDetailRepositoryTestSuite))
}

func (suite *AttendanceRepositoryTestSuite) TestCreate() {
	dummy := model.Attendance{
		AttendanceID:      "098709687657645654567586778",
		SessionID:         "0980687gtr65fr65r",
		StudentID:         "098098gv6vtrfr6d6r",
		AttendanceStudent: true,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
		IsDeleted:         false,
	}
	suite.sqlmock.ExpectBegin()
	rows := sqlmock.NewRows([]string{" attendance_id", "session_id", "student_id", "attendance_student", "created_at", "updated_at", "is_deleted"}).AddRow(dummy.AttendanceID, dummy.SessionID, dummy.StudentID, dummy.AttendanceStudent, dummy.CreatedAt, dummy.UpdatedAt, dummy.IsDeleted)
	suite.sqlmock.ExpectQuery("insert into attendance").WillReturnRows(rows)
	suite.sqlmock.ExpectCommit()

	actual, err := suite.repo.Create(dummy)
	assert.Nil(suite.T(), err, "Error should be nil")

	assert.Equal(suite.T(), dummy.AttendanceID, actual.AttendanceID, "AttendanceID should match")
	assert.Equal(suite.T(), dummy.SessionID, actual.SessionID, "SessionID should match")
	assert.Equal(suite.T(), dummy.StudentID, actual.StudentID, "StudentID should match")
	assert.Equal(suite.T(), dummy.AttendanceID, actual.AttendanceID, "AttendanceID should match")
	assert.Equal(suite.T(), dummy.CreatedAt, actual.CreatedAt, "CreatedAt should match")
	assert.Equal(suite.T(), dummy.UpdatedAt, actual.UpdatedAt, "UpdatedAt should match")
	assert.Equal(suite.T(), dummy.IsDeleted, actual.IsDeleted, "IsDeleted should match")
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

func (suite *AttendanceRepositoryTestSuite) TestUpdated() {

	dummyPayload := model.Attendance{
		SessionID:         "hhhhhhhhhhhhhhhhhhhhhh",
		StudentID:         "nnnnnnnnnnnnnnnnnnnnnn",
		AttendanceStudent: true,
		UpdatedAt:         time.Now(),
	}

	dummyResult := model.Attendance{
		AttendanceID:      "098709687657645654567586778",
		SessionID:         dummyPayload.SessionID,
		StudentID:         dummyPayload.StudentID,
		AttendanceStudent: dummyPayload.AttendanceStudent,
		CreatedAt:         time.Now(),
		UpdatedAt:         dummyPayload.UpdatedAt,
		IsDeleted:         false,
	}

	query := "update attendance set session_id=\\$1 ,student_id=\\$2, attendance_student=\\$3, updated_at = \\$4, is_deleted = \\$5 where attendance_id = \\$6 returning attendance_id, session_id ,student_id, attendance_student, created_at, updated_at, is_deleted;"

	rows := sqlmock.NewRows([]string{"attendance_id", "session_id", "student_id", "attendance_student", "created_at", "updated_at", "is_deleted"}).
		AddRow(dummyResult.AttendanceID, dummyResult.SessionID, dummyResult.StudentID, dummyResult.AttendanceStudent, dummyResult.CreatedAt, dummyResult.UpdatedAt, dummyResult.IsDeleted)

	suite.sqlmock.ExpectQuery(query).
		WithArgs(
			dummyPayload.SessionID,
			dummyPayload.StudentID,
			dummyPayload.AttendanceStudent,
			dummyPayload.IsDeleted,
			dummyPayload.UpdatedAt,
			dummyResult.AttendanceID,
		).
		WillReturnRows(rows)

	actual, err := suite.repo.Update(dummyPayload, dummyResult.AttendanceID)

	assert.Nil(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), dummyResult.AttendanceID, actual.AttendanceID, "AttendanceID should match")
	assert.Equal(suite.T(), dummyPayload.SessionID, actual.SessionID, "SessionID should match")
	assert.Equal(suite.T(), dummyPayload.StudentID, actual.StudentID, "StudentID should match")
	assert.Equal(suite.T(), dummyPayload.AttendanceStudent, actual.AttendanceStudent, "AttendanceStudent should match")
	assert.Equal(suite.T(), dummyResult.CreatedAt, actual.CreatedAt, "CreatedAt should match")
	assert.WithinDuration(suite.T(), dummyResult.UpdatedAt, actual.UpdatedAt, time.Second, "UpdatedAt should be close to current time")
	assert.Equal(suite.T(), dummyResult.IsDeleted, actual.IsDeleted, "IsDeleted should match")
}

func (suite *AttendanceRepositoryTestSuite) TestDelete() {

	dummyResult := model.Attendance{
		AttendanceID:      "098709687657645654567586778",
		SessionID:         "0983485039485345",
		StudentID:         "928374298374234234",
		AttendanceStudent: true,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
		IsDeleted:         true,
	}

	query := "update attendance set is_deleted = \\$1 where attendance_id = \\$2 returning attendance_id, session_id ,student_id, attendance_student, created_at, updated_at, is_deleted;"

	rows := sqlmock.NewRows([]string{"attendance_id", "session_id", "student_id", "attendance_student", "created_at", "updated_at", "is_deleted"}).
		AddRow(dummyResult.AttendanceID, dummyResult.SessionID, dummyResult.StudentID, dummyResult.AttendanceStudent, dummyResult.CreatedAt, dummyResult.UpdatedAt, dummyResult.IsDeleted)

	suite.sqlmock.ExpectQuery(query).
		WithArgs(
			true,
			dummyResult.AttendanceID,
		).
		WillReturnRows(rows)

	actual, err := suite.repo.Delete(dummyResult.AttendanceID)

	assert.Nil(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), dummyResult.AttendanceID, actual.AttendanceID, "AttendanceID should match")
	assert.Equal(suite.T(), dummyResult.SessionID, actual.SessionID, "SessionID should match")
	assert.Equal(suite.T(), dummyResult.StudentID, actual.StudentID, "StudentID should match")
	assert.Equal(suite.T(), dummyResult.AttendanceStudent, actual.AttendanceStudent, "AttendanceStudent should match")
	assert.Equal(suite.T(), dummyResult.CreatedAt, actual.CreatedAt, "CreatedAt should match")
	assert.WithinDuration(suite.T(), dummyResult.UpdatedAt, actual.UpdatedAt, time.Second, "UpdatedAt should be close to current time")
	assert.Equal(suite.T(), dummyResult.IsDeleted, actual.IsDeleted, "IsDeleted should match")
}

// func (suite *AttendanceRepositoryTestSuite) TestFindAll() {

// 	dummyResults := []model.Attendance{
// 		{
// 			AttendanceID:      "098709687657645654567586778",
// 			SessionID:         "session1",
// 			StudentID:         "student1",
// 			AttendanceStudent: true,
// 			CreatedAt:         time.Now(),
// 			UpdatedAt:         time.Now(),
// 			IsDeleted:         false,
// 		},
// 		{
// 			AttendanceID:      "098709687657645654567586779",
// 			SessionID:         "session2",
// 			StudentID:         "student2",
// 			AttendanceStudent: false,
// 			CreatedAt:         time.Now(),
// 			UpdatedAt:         time.Now(),
// 			IsDeleted:         false,
// 		},
// 	}

// 	query := "select * from attendance where is_deleted = $1;"

// 	rows := sqlmock.NewRows([]string{"attendance_id", "session_id", "student_id", "attendance_student", "created_at", "updated_at", "is_deleted"})

// 	for _, result := range dummyResults {
// 		rows.AddRow(result.AttendanceID, result.SessionID, result.StudentID, result.AttendanceStudent, result.CreatedAt, result.UpdatedAt, result.IsDeleted)
// 	}

// 	suite.sqlmock.ExpectQuery(query).
// 		WithArgs(false).
// 		WillReturnRows(rows)

// 	actual, err := suite.repo.FindAll()

// 	assert.Nil(suite.T(), err, "Error should be nil")
// 	assert.Len(suite.T(), actual, len(dummyResults), "Number of results should match")

// 	for i, expected := range dummyResults {
// 		assert.Equal(suite.T(), expected.AttendanceID, actual[i].AttendanceID, "AttendanceID should match")
// 		assert.Equal(suite.T(), expected.SessionID, actual[i].SessionID, "SessionID should match")
// 		assert.Equal(suite.T(), expected.StudentID, actual[i].StudentID, "StudentID should match")
// 		assert.Equal(suite.T(), expected.AttendanceStudent, actual[i].AttendanceStudent, "AttendanceStudent should match")
// 		assert.Equal(suite.T(), expected.CreatedAt, actual[i].CreatedAt, "CreatedAt should match")
// 		assert.WithinDuration(suite.T(), expected.UpdatedAt, actual[i].UpdatedAt, time.Second, "UpdatedAt should be close to current time")
// 		assert.Equal(suite.T(), expected.IsDeleted, actual[i].IsDeleted, "IsDeleted should match")
// 	}
// }
