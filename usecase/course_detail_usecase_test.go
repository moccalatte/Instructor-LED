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

type CourseDetailUseCaseTestSuite struct {
	suite.Suite
	cdrm *repomock.CourseDetailRepoMock
	cdu  CourseDetailUseCase
}

func (suite *CourseDetailUseCaseTestSuite) SetupTest() {
	suite.cdrm = new(repomock.CourseDetailRepoMock)
	suite.cdu = NewCourseDetailUseCase(suite.cdrm)
}

func TestCourseDetailUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(CourseDetailUseCaseTestSuite))
}

func (suite *CourseDetailUseCaseTestSuite) TestAddCourseDetailSuccess() {
	payload := dto.CourseDetailRequestDto{
		CourseId:      "course-1",
		CourseChapter: "Chapter 1",
		CourseContent: "Content of Chapter 1",
	}

	expectedCourseDetail := model.CourseDetail{
		CourseID:      payload.CourseId,
		CourseChapter: payload.CourseChapter,
		CourseContent: payload.CourseContent,
	}

	suite.cdrm.On("Create", mock.AnythingOfType("model.CourseDetail")).Return(expectedCourseDetail, nil)

	addedCourseDetail, err := suite.cdu.AddCourse(payload)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedCourseDetail, addedCourseDetail)
	suite.cdrm.AssertExpectations(suite.T())
}

func (suite *CourseDetailUseCaseTestSuite) TestAddCourseDetailError() {
	payload := dto.CourseDetailRequestDto{
		CourseId:      "course-1",
		CourseChapter: "Chapter 1",
		CourseContent: "Content of Chapter 1",
	}

	expectedError := errors.New("some error")

	suite.cdrm.On("Create", mock.AnythingOfType("model.CourseDetail")).Return(model.CourseDetail{}, expectedError)

	addedCourseDetail, err := suite.cdu.AddCourse(payload)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.CourseDetail{}, addedCourseDetail)
	suite.cdrm.AssertExpectations(suite.T())
}

func (suite *CourseDetailUseCaseTestSuite) TestFindCourseDetailByIDSuccess() {
	targetID := "some-id"

	expectedCourseDetail := model.CourseDetail{
		CourseID:      "course-1",
		CourseChapter: "Chapter 1",
		CourseContent: "Content of Chapter 1",
	}

	suite.cdrm.On("GetById", targetID).Return(expectedCourseDetail, nil)

	resultCourseDetail, err := suite.cdu.FindCourseDetailByID(targetID)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedCourseDetail, resultCourseDetail)
	suite.cdrm.AssertExpectations(suite.T())
}

func (suite *CourseDetailUseCaseTestSuite) TestFindCourseDetailByIDError() {
	invalidID := "non-existent-id"

	expectedError := errors.New("data not found")
	suite.cdrm.On("GetById", invalidID).Return(model.CourseDetail{}, expectedError)

	resultCourseDetail, err := suite.cdu.FindCourseDetailByID(invalidID)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.CourseDetail{}, resultCourseDetail)
	suite.cdrm.AssertExpectations(suite.T())
}

func (suite *CourseDetailUseCaseTestSuite) TestUpdateCourseDetailSuccess() {
	targetID := "some-id"

	updatePayload := dto.CourseDetailRequestDto{
		CourseId:      "updated-course-1",
		CourseChapter: "Updated Chapter 1",
		CourseContent: "Updated Content of Chapter 1",
	}

	updatedCourseDetail := model.CourseDetail{
		CourseID:      updatePayload.CourseId,
		CourseChapter: updatePayload.CourseChapter,
		CourseContent: updatePayload.CourseContent,
	}

	suite.cdrm.On("Update", mock.AnythingOfType("model.CourseDetail"), targetID).Return(updatedCourseDetail, nil)

	resultCourseDetail, err := suite.cdu.UpdateAttendance(updatePayload, targetID)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), updatedCourseDetail, resultCourseDetail)
	suite.cdrm.AssertExpectations(suite.T())
}

func (suite *CourseDetailUseCaseTestSuite) TestUpdateCourseDetailError() {
	targetID := "some-id"

	updatePayload := dto.CourseDetailRequestDto{
		CourseId: "updated-course-1",
	}

	expectedError := errors.New("failed update data by id : SomeError")
	suite.cdrm.On("Update", mock.AnythingOfType("model.CourseDetail"), targetID).Return(model.CourseDetail{}, expectedError)

	_, err := suite.cdu.UpdateAttendance(updatePayload, targetID)

	assert.Error(suite.T(), err)
	assert.EqualError(suite.T(), err, "failed to Update Course Detail : failed update data by id : SomeError")
	suite.cdrm.AssertExpectations(suite.T())
}

func (suite *CourseDetailUseCaseTestSuite) TestDeleteCourseDetailSuccess() {
	targetID := "some-id"

	deletedCourseDetail := model.CourseDetail{
		CourseID:      "course-1",
		CourseChapter: "Chapter 1",
		CourseContent: "Content of Chapter 1",
	}

	suite.cdrm.On("Delete", targetID).Return(deletedCourseDetail, nil)

	resultCourseDetail, err := suite.cdu.Delete(targetID)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), deletedCourseDetail, resultCourseDetail)
	suite.cdrm.AssertExpectations(suite.T())
}

func (suite *CourseDetailUseCaseTestSuite) TestDeleteCourseDetailError() {
	targetID := "some-id"

	expectedError := errors.New("failed to delete data")
	suite.cdrm.On("Delete", targetID).Return(model.CourseDetail{}, expectedError)

	resultCourseDetail, err := suite.cdu.Delete(targetID)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.CourseDetail{}, resultCourseDetail)
	suite.cdrm.AssertExpectations(suite.T())
}
