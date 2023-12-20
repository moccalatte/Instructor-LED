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

type QuestionUseCaseTestSuite struct {
	suite.Suite
	qrm *repomock.QuestionRepoMock
	qu  QuestionUseCase
}

func (suite *QuestionUseCaseTestSuite) SetupTest() {
	suite.qrm = new(repomock.QuestionRepoMock)
	suite.qu = NewQuestion(suite.qrm)
}

func TestQuestionUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(QuestionUseCaseTestSuite))
}

func (suite *QuestionUseCaseTestSuite) TestAddQuestionSuccess() {
	payload := dto.QuestionRequestDto{
		SessionID:   "session-1",
		StudentID:   "student-1",
		TrainerID:   "trainer-1",
		Title:       "Question Title",
		Description: "Question Description",
		CourseID:    "course-1",
		Image:       "question-image.png",
		Answer:      "Answer to the question",
		Status:      "answered",
	}

	expectedQuestion := model.Question{
		SessionID:   payload.SessionID,
		StudentID:   payload.StudentID,
		TrainerID:   payload.TrainerID,
		Title:       payload.Title,
		Description: payload.Description,
		CourseID:    payload.CourseID,
		Image:       payload.Image,
		Answer:      payload.Answer,
		Status:      payload.Status,
	}

	suite.qrm.On("Create", mock.AnythingOfType("model.Question")).Return(expectedQuestion, nil)

	addedQuestion, err := suite.qu.AddQuestion(payload)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedQuestion, addedQuestion)
	suite.qrm.AssertExpectations(suite.T())
}

func (suite *QuestionUseCaseTestSuite) TestAddQuestionError() {
	payload := dto.QuestionRequestDto{
		SessionID:   "session-1",
		StudentID:   "student-1",
		TrainerID:   "trainer-1",
		Title:       "Question Title",
		Description: "Question Description",
		CourseID:    "course-1",
		Image:       "question-image.png",
		Answer:      "Answer to the question",
		Status:      "answered",
	}

	expectedError := errors.New("some error")

	suite.qrm.On("Create", mock.AnythingOfType("model.Question")).Return(model.Question{}, expectedError)

	addedQuestion, err := suite.qu.AddQuestion(payload)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.Question{}, addedQuestion)
	suite.qrm.AssertExpectations(suite.T())
}

func (suite *QuestionUseCaseTestSuite) TestFindQuestionByIdSuccess() {
	targetID := "some-id"

	expectedQuestion := model.Question{
		SessionID:   "session-1",
		StudentID:   "student-1",
		TrainerID:   "trainer-1",
		Title:       "Question Title",
		Description: "Question Description",
		CourseID:    "course-1",
		Image:       "question-image.png",
		Answer:      "Answer to the question",
		Status:      "answered",
	}

	suite.qrm.On("GetById", targetID).Return(expectedQuestion, nil)

	resultQuestion, err := suite.qu.FindQuestionById(targetID)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedQuestion, resultQuestion)
	suite.qrm.AssertExpectations(suite.T())
}

func (suite *QuestionUseCaseTestSuite) TestFindQuestionByIdError() {
	invalidID := "non-existent-id"

	expectedError := errors.New("data not found")
	suite.qrm.On("GetById", invalidID).Return(model.Question{}, expectedError)

	resultQuestion, err := suite.qu.FindQuestionById(invalidID)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.Question{}, resultQuestion)
	suite.qrm.AssertExpectations(suite.T())
}

func (suite *QuestionUseCaseTestSuite) TestFindQuestionByStudentIdSuccess() {
	targetID := "some-id"

	expectedQuestion := model.Question{
		SessionID:   "session-1",
		StudentID:   "student-1",
		TrainerID:   "trainer-1",
		Title:       "Question Title",
		Description: "Question Description",
		CourseID:    "course-1",
		Image:       "question-image.png",
		Answer:      "Answer to the question",
		Status:      "answered",
	}

	suite.qrm.On("GetByStudentId", targetID).Return(expectedQuestion, nil)

	resultQuestion, err := suite.qu.FindQuestionByStudentId(targetID)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedQuestion, resultQuestion)
	suite.qrm.AssertExpectations(suite.T())
}

func (suite *QuestionUseCaseTestSuite) TestFindQuestionByStudentIdError() {
	invalidID := "non-existent-id"

	expectedError := errors.New("data not found")
	suite.qrm.On("GetByStudentId", invalidID).Return(model.Question{}, expectedError)

	resultQuestion, err := suite.qu.FindQuestionByStudentId(invalidID)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.Question{}, resultQuestion)
	suite.qrm.AssertExpectations(suite.T())
}

func (suite *QuestionUseCaseTestSuite) TestGetAllQuestion_Success() {
	expectedQuestions := []model.Question{
		{
			QuestionID: "1",
			SessionID:  "1",
			StudentID:  "1",
			TrainerID:  "1",
		},
	}
	suite.qrm.On("GetAll").Return(expectedQuestions, nil)

	questions, err := suite.qu.GetAllQuestion()

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedQuestions, questions)

	suite.qrm.AssertExpectations(suite.T())
}

func (suite *QuestionUseCaseTestSuite) TestGetAllQuestion_ErrorFromRepository() {
	suite.qrm.On("GetAll").Return([]model.Question{}, errors.New("repository error"))

	questions, err := suite.qu.GetAllQuestion()

	assert.Error(suite.T(), err)
	assert.Empty(suite.T(), questions)

	suite.qrm.AssertExpectations(suite.T())
}

func (suite *QuestionUseCaseTestSuite) TestUpdateQuestionSuccess() {
	targetID := "some-id"

	updatePayload := dto.QuestionRequestDto{
		SessionID:   "updated-session-1",
		StudentID:   "updated-student-1",
		TrainerID:   "updated-trainer-1",
		Title:       "Updated Question Title",
		Description: "Updated Question Description",
		CourseID:    "updated-course-1",
		Image:       "updated-question-image.png",
		Answer:      "Updated Answer to the question",
		Status:      "updated-answered",
	}

	updatedQuestion := model.Question{
		SessionID:   updatePayload.SessionID,
		StudentID:   updatePayload.StudentID,
		TrainerID:   updatePayload.TrainerID,
		Title:       updatePayload.Title,
		Description: updatePayload.Description,
		CourseID:    updatePayload.CourseID,
		Image:       updatePayload.Image,
		Answer:      updatePayload.Answer,
		Status:      updatePayload.Status,
	}

	suite.qrm.On("Update", mock.AnythingOfType("model.Question"), targetID).Return(updatedQuestion, nil)

	resultQuestion, err := suite.qu.Update(updatePayload, targetID)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), updatedQuestion, resultQuestion)
	suite.qrm.AssertExpectations(suite.T())
}

func (suite *QuestionUseCaseTestSuite) TestUpdateQuestionError() {
	targetID := "some-id"

	updatePayload := dto.QuestionRequestDto{
		SessionID: "updated-session-1",
	}

	expectedError := errors.New("failed update data by id : SomeError")
	suite.qrm.On("Update", mock.AnythingOfType("model.Question"), targetID).Return(model.Question{}, expectedError)

	_, err := suite.qu.Update(updatePayload, targetID)

	assert.Error(suite.T(), err)
	assert.EqualError(suite.T(), err, "failed to Update Question : failed update data by id : SomeError")
	suite.qrm.AssertExpectations(suite.T())
}

func (suite *QuestionUseCaseTestSuite) TestDeleteQuestionSuccess() {
	targetID := "some-id"

	deletedQuestion := model.Question{
		SessionID:   "session-1",
		StudentID:   "student-1",
		TrainerID:   "trainer-1",
		Title:       "Question Title",
		Description: "Question Description",
		CourseID:    "course-1",
		Image:       "question-image.png",
		Answer:      "Answer to the question",
		Status:      "answered",
	}

	suite.qrm.On("Delete", targetID).Return(deletedQuestion, nil)

	resultQuestion, err := suite.qu.Delete(targetID)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), deletedQuestion, resultQuestion)
	suite.qrm.AssertExpectations(suite.T())
}

func (suite *QuestionUseCaseTestSuite) TestDeleteQuestionError() {
	targetID := "some-id"

	expectedError := errors.New("failed to delete data")
	suite.qrm.On("Delete", targetID).Return(model.Question{}, expectedError)

	resultQuestion, err := suite.qu.Delete(targetID)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.Question{}, resultQuestion)
	suite.qrm.AssertExpectations(suite.T())
}

func (suite *QuestionUseCaseTestSuite) TestAnswerQuestionSuccess() {
	targetID := "some-id"

	answerPayload := dto.QuestionRequestDto{
		Answer: "Updated Answer to the question",
	}

	answeredQuestion := model.Question{
		Answer: answerPayload.Answer,
	}

	suite.qrm.On("Answer", mock.AnythingOfType("model.Question"), targetID).Return(answeredQuestion, nil)

	resultQuestion, err := suite.qu.Answer(answerPayload, targetID)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), answeredQuestion, resultQuestion)
	suite.qrm.AssertExpectations(suite.T())
}

func (suite *QuestionUseCaseTestSuite) TestAnswerQuestionError() {
	targetID := "some-id"

	answerPayload := dto.QuestionRequestDto{
		Answer: "Updated Answer to the question",
	}

	expectedError := errors.New("failed to answer question")
	suite.qrm.On("Answer", mock.AnythingOfType("model.Question"), targetID).Return(model.Question{}, expectedError)

	resultQuestion, err := suite.qu.Answer(answerPayload, targetID)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.Question{}, resultQuestion)
	suite.qrm.AssertExpectations(suite.T())
}
