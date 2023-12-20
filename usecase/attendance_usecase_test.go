package usecase

import (
	"errors"
	repomock "final-project-kelompok-1/mock/repo_mock"
	"final-project-kelompok-1/model"
	"final-project-kelompok-1/model/dto"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type AttendanceUseCaseTestSuite struct {
	suite.Suite
	arm *repomock.AttendanceRepoMock
	au  AttendanceUseCase
}

func (suite *AttendanceUseCaseTestSuite) SetupTest() {
	suite.arm = new(repomock.AttendanceRepoMock)
	suite.au = NewAttendanceUseCase(suite.arm)
}

func TestAttendanceUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(AttendanceUseCaseTestSuite))
}

func (suite *AttendanceUseCaseTestSuite) TestAddAttendanceSuccess() {
	payload := dto.AttendanceRequestDto{
		SessionID:         "session-1",
		StudentID:         "student-1",
		AttendanceStudent: true,
	}

	expectedAttendance := model.Attendance{
		SessionID:         payload.SessionID,
		StudentID:         payload.StudentID,
		AttendanceStudent: payload.AttendanceStudent,
	}

	suite.arm.On("Create", mock.AnythingOfType("model.Attendance")).Return(expectedAttendance, nil)

	addedAttendance, err := suite.au.AddAttendance(payload)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedAttendance, addedAttendance)
	suite.arm.AssertExpectations(suite.T())
}

func (suite *AttendanceUseCaseTestSuite) TestAddAttendanceError() {
	payload := dto.AttendanceRequestDto{
		SessionID:         "session-1",
		StudentID:         "student-1",
		AttendanceStudent: true,
	}

	expectedError := errors.New("some error")

	suite.arm.On("Create", mock.AnythingOfType("model.Attendance")).Return(model.Attendance{}, expectedError)

	addedAttendance, err := suite.au.AddAttendance(payload)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.Attendance{}, addedAttendance)
	suite.arm.AssertExpectations(suite.T())
}

func (suite *AttendanceUseCaseTestSuite) TestFindAttendanceByIDSuccess() {
	targetID := "some-id"

	expectedAttendance := model.Attendance{
		SessionID:         "session-1",
		StudentID:         "student-1",
		AttendanceStudent: true,
	}

	suite.arm.On("GetById", targetID).Return(expectedAttendance, nil)

	resultAttendance, err := suite.au.FindAttendanceByID(targetID)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedAttendance, resultAttendance)
	suite.arm.AssertExpectations(suite.T())
}

func (suite *AttendanceUseCaseTestSuite) TestFindAttendanceByIDError() {
	invalidID := "non-existent-id"

	expectedError := errors.New("data not found")
	suite.arm.On("GetById", invalidID).Return(model.Attendance{}, expectedError)

	resultAttendance, err := suite.au.FindAttendanceByID(invalidID)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.Attendance{}, resultAttendance)
	suite.arm.AssertExpectations(suite.T())
}

func (suite *AttendanceUseCaseTestSuite) TestFindAttendanceBySessionIDSuccess() {
	targetID := "some-id"

	expectedAttendance := model.Attendance{
		SessionID:         "session-1",
		StudentID:         "student-1",
		AttendanceStudent: true,
	}

	suite.arm.On("GetBySessionId", targetID).Return(expectedAttendance, nil)

	resultAttendance, err := suite.au.FindAttendanceBySessionId(targetID)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedAttendance, resultAttendance)
	suite.arm.AssertExpectations(suite.T())
}

func (suite *AttendanceUseCaseTestSuite) TestFindAttendanceBySessionIDError() {
	invalidID := "non-existent-id"

	expectedError := errors.New("data not found")
	suite.arm.On("GetBySessionId", invalidID).Return(model.Attendance{}, expectedError)

	resultAttendance, err := suite.au.FindAttendanceBySessionId(invalidID)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.Attendance{}, resultAttendance)
	suite.arm.AssertExpectations(suite.T())
}

func (suite *AttendanceUseCaseTestSuite) TestFindAttendanceBySessionIdSuccess() {
	targetID := "some-id"

	expectedAttendance := model.Attendance{
		SessionID:         "session-1",
		StudentID:         "student-1",
		AttendanceStudent: true,
	}

	suite.arm.On("GetBySessionId", targetID).Return(expectedAttendance, nil)

	resultAttendance, err := suite.au.FindAttendanceBySessionId(targetID)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedAttendance, resultAttendance)
	suite.arm.AssertExpectations(suite.T())
}

func (suite *AttendanceUseCaseTestSuite) TestFindAttendanceBySessionIdError() {
	invalidID := "non-existent-id"

	expectedError := errors.New("data not found")
	suite.arm.On("GetBySessionId", invalidID).Return(model.Attendance{}, expectedError)

	resultAttendance, err := suite.au.FindAttendanceBySessionId(invalidID)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.Attendance{}, resultAttendance)
	suite.arm.AssertExpectations(suite.T())
}

func (suite *AttendanceUseCaseTestSuite) TestGetAllAttendance_Success() {
	expectedAttendance := []model.Attendance{
		{
			AttendanceID:      "1",
			SessionID:         "C1",
			StudentID:         "S1",
			AttendanceStudent: true,
		},
	}
	suite.arm.On("GetAll").Return(expectedAttendance, nil)

	attendance, err := suite.au.GetAllAttendance()

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedAttendance, attendance)

	suite.arm.AssertExpectations(suite.T())
}

func (suite *AttendanceUseCaseTestSuite) TestGetAllAttendance_ErrorFromRepository() {
	suite.arm.On("GetAll").Return([]model.Attendance{}, errors.New("repository error"))

	attendance, err := suite.au.GetAllAttendance()

	assert.Error(suite.T(), err)
	assert.Empty(suite.T(), attendance)

	suite.arm.AssertExpectations(suite.T())
}

func (suite *AttendanceUseCaseTestSuite) TestUpdateAttendanceSuccess() {
	targetID := "some-id"

	updatePayload := dto.AttendanceRequestDto{
		SessionID:         "updated-session-1",
		StudentID:         "updated-student-1",
		AttendanceStudent: false,
	}

	updatedAttendance := model.Attendance{
		SessionID:         updatePayload.SessionID,
		StudentID:         updatePayload.StudentID,
		AttendanceStudent: updatePayload.AttendanceStudent,
	}

	suite.arm.On("Update", mock.AnythingOfType("model.Attendance"), targetID).Return(updatedAttendance, nil)

	resultAttendance, err := suite.au.UpdateAttendance(updatePayload, targetID)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), updatedAttendance, resultAttendance)
	suite.arm.AssertExpectations(suite.T())
}

func (suite *AttendanceUseCaseTestSuite) TestUpdateAttendanceError() {
	targetID := "some-id"

	updatePayload := dto.AttendanceRequestDto{
		SessionID: "updated-session-1",
	}

	expectedError := errors.New("failed update data by id : SomeError")
	suite.arm.On("Update", mock.AnythingOfType("model.Attendance"), targetID).Return(model.Attendance{}, expectedError)

	_, err := suite.au.UpdateAttendance(updatePayload, targetID)

	assert.Error(suite.T(), err)
	assert.EqualError(suite.T(), err, "failed to Update Attendance : failed update data by id : SomeError")
	suite.arm.AssertExpectations(suite.T())
}

func (suite *AttendanceUseCaseTestSuite) TestDeleteAttendanceSuccess() {
	targetID := "some-id"

	deletedAttendance := model.Attendance{
		SessionID:         "session-1",
		StudentID:         "student-1",
		AttendanceStudent: true,
	}

	suite.arm.On("Delete", targetID).Return(deletedAttendance, nil)

	resultAttendance, err := suite.au.Delete(targetID)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), deletedAttendance, resultAttendance)
	suite.arm.AssertExpectations(suite.T())
}

func (suite *AttendanceUseCaseTestSuite) TestDeleteAttendanceError() {
	targetID := "some-id"

	expectedError := errors.New("failed to delete data")
	suite.arm.On("Delete", targetID).Return(model.Attendance{}, expectedError)

	resultAttendance, err := suite.au.Delete(targetID)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.Attendance{}, resultAttendance)
	suite.arm.AssertExpectations(suite.T())
}
