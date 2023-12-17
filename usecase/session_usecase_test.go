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

type SessionUseCaseTestSuite struct {
	suite.Suite
	srm *repomock.SessionRepoMock
	su  SessionUseCase
}

func (suite *SessionUseCaseTestSuite) SetupTest() {
	suite.srm = new(repomock.SessionRepoMock)
	suite.su = NewSession(suite.srm)
}

func TestSessionUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(SessionUseCaseTestSuite))
}

func (suite *SessionUseCaseTestSuite) TestAddSessionSuccess() {
	payload := dto.SessionRequestDto{
		Title:       "Session 1",
		Description: "Description of Session 1",
		SessionDate: "2023-12-31",
		SessionTime: "12:00 PM",
		SessionLink: "https://example.com/session1",
		TrainerID:   "trainer-1",
	}

	expectedSession := model.Session{
		Title:       payload.Title,
		Description: payload.Description,
		SessionDate: payload.SessionDate,
		SessionTime: payload.SessionTime,
		SessionLink: payload.SessionLink,
		TrainerID:   payload.TrainerID,
	}

	suite.srm.On("Create", mock.AnythingOfType("model.Session")).Return(expectedSession, nil)

	addedSession, err := suite.su.AddSession(payload)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedSession, addedSession)
	suite.srm.AssertExpectations(suite.T())
}

func (suite *SessionUseCaseTestSuite) TestAddSessionError() {
	payload := dto.SessionRequestDto{
		Title:       "Session 1",
		Description: "Description of Session 1",
		SessionDate: "2023-12-31",
		SessionTime: "12:00 PM",
		SessionLink: "https://example.com/session1",
		TrainerID:   "trainer-1",
	}

	expectedError := errors.New("some error")

	suite.srm.On("Create", mock.AnythingOfType("model.Session")).Return(model.Session{}, expectedError)

	addedSession, err := suite.su.AddSession(payload)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.Session{}, addedSession)
	suite.srm.AssertExpectations(suite.T())
}

func (suite *SessionUseCaseTestSuite) TestFindSessionByIdSuccess() {
	targetID := "some-id"

	expectedSession := model.Session{
		Title:       "Session 1",
		Description: "Description of Session 1",
		SessionDate: "2023-12-31",
		SessionTime: "12:00 PM",
		SessionLink: "https://example.com/session1",
		TrainerID:   "trainer-1",
	}

	suite.srm.On("GetById", targetID).Return(expectedSession, nil)

	resultSession, err := suite.su.FindSessionById(targetID)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedSession, resultSession)
	suite.srm.AssertExpectations(suite.T())
}

func (suite *SessionUseCaseTestSuite) TestFindSessionByIdError() {
	invalidID := "non-existent-id"

	expectedError := errors.New("data not found")
	suite.srm.On("GetById", invalidID).Return(model.Session{}, expectedError)

	resultSession, err := suite.su.FindSessionById(invalidID)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.Session{}, resultSession)
	suite.srm.AssertExpectations(suite.T())
}

func (suite *SessionUseCaseTestSuite) TestGetAllSession_Success() {
	// Mock repository response
	expectedSessions := []model.Session{
		{
			SessionID:   "1",
			Title:       "Introduction to Go",
			Description: "Learn the basics of Go programming language",
			SessionDate: "2023-01-01",
		},
	}
	suite.srm.On("GetAllSession").Return(expectedSessions, nil)

	// Call the method being tested
	sessions, err := suite.su.GetAllSession()

	// Assertions
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedSessions, sessions)

	// Assert that the expected method was called
	suite.srm.AssertExpectations(suite.T())
}

func (suite *SessionUseCaseTestSuite) TestGetAllSession_ErrorFromRepository() {
	// Mock repository response with an error
	suite.srm.On("GetAllSession").Return([]model.Session{}, errors.New("repository error"))

	// Call the method being tested
	sessions, err := suite.su.GetAllSession()

	// Assertions
	assert.Error(suite.T(), err)
	assert.Empty(suite.T(), sessions)

	// Assert that the expected method was called
	suite.srm.AssertExpectations(suite.T())
}

func (suite *SessionUseCaseTestSuite) TestUpdateSessionSuccess() {
	targetID := "some-id"

	updatePayload := dto.SessionRequestDto{
		Title:       "Updated Session 1",
		Description: "Updated Description of Session 1",
		SessionDate: "2023-12-31",
		SessionTime: "12:00 PM",
		SessionLink: "https://example.com/updated-session1",
		TrainerID:   "updated-trainer-1",
	}

	updatedSession := model.Session{
		Title:       updatePayload.Title,
		Description: updatePayload.Description,
		SessionDate: updatePayload.SessionDate,
		SessionTime: updatePayload.SessionTime,
		SessionLink: updatePayload.SessionLink,
		TrainerID:   updatePayload.TrainerID,
	}

	suite.srm.On("Update", mock.AnythingOfType("model.Session"), targetID).Return(updatedSession, nil)

	resultSession, err := suite.su.Update(updatePayload, targetID)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), updatedSession, resultSession)
	suite.srm.AssertExpectations(suite.T())
}

func (suite *SessionUseCaseTestSuite) TestUpdateSessionError() {
	targetID := "some-id"

	updatePayload := dto.SessionRequestDto{
		Title: "Updated Session 1",
	}

	expectedError := errors.New("failed update data by id : SomeError")
	suite.srm.On("Update", mock.AnythingOfType("model.Session"), targetID).Return(model.Session{}, expectedError)

	_, err := suite.su.Update(updatePayload, targetID)

	assert.Error(suite.T(), err)
	assert.EqualError(suite.T(), err, "failed to Update Session : failed update data by id : SomeError")
	suite.srm.AssertExpectations(suite.T())
}

func (suite *SessionUseCaseTestSuite) TestUpdateSessionNoteSuccess() {
	targetID := "some-id"

	updatePayload := dto.SessionRequestDto{
		Note: "mwehehehehehe",
	}

	updatedSession := model.Session{
		Note: updatePayload.Note,
	}

	suite.srm.On("UpdateNote", mock.AnythingOfType("model.Session"), targetID).Return(updatedSession, nil)

	resultSession, err := suite.su.UpdateNote(updatePayload, targetID)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), updatedSession, resultSession)
	suite.srm.AssertExpectations(suite.T())
}

func (suite *SessionUseCaseTestSuite) TestUpdateSessionNoteError() {
	targetID := "some-id"

	updatePayload := dto.SessionRequestDto{
		Note: "jjjjjjjjjjj",
	}

	expectedError := errors.New("failed update data by id : SomeError")
	suite.srm.On("UpdateNote", mock.AnythingOfType("model.Session"), targetID).Return(model.Session{}, expectedError)

	_, err := suite.su.UpdateNote(updatePayload, targetID)

	assert.Error(suite.T(), err)
	assert.EqualError(suite.T(), err, "failed to Update Session : failed update data by id : SomeError")
	suite.srm.AssertExpectations(suite.T())
}

func (suite *SessionUseCaseTestSuite) TestDeleteSessionSuccess() {
	targetID := "some-id"

	deletedSession := model.Session{
		Title:       "Session 1",
		Description: "Description of Session 1",
		SessionDate: "2023-12-31",
		SessionTime: "12:00 PM",
		SessionLink: "https://example.com/session1",
		TrainerID:   "trainer-1",
	}

	suite.srm.On("Delete", targetID).Return(deletedSession, nil)

	resultSession, err := suite.su.Delete(targetID)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), deletedSession, resultSession)
	suite.srm.AssertExpectations(suite.T())
}

func (suite *SessionUseCaseTestSuite) TestDeleteSessionError() {
	targetID := "some-id"

	expectedError := errors.New("failed to delete data")
	suite.srm.On("Delete", targetID).Return(model.Session{}, expectedError)

	resultSession, err := suite.su.Delete(targetID)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.Session{}, resultSession)
	suite.srm.AssertExpectations(suite.T())
}
