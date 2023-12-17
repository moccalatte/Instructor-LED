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

type CourseUseCaseTestSuite struct {
	suite.Suite
	crm *repomock.CourseRepoMock
	cu  CourseUseCase
}

func (suite *CourseUseCaseTestSuite) SetupTest() {
	suite.crm = new(repomock.CourseRepoMock)
	suite.cu = NewCourseUseCase(suite.crm)
}

// func NewCourseUseCase(courseRepoMock *repomock.CourseRepoMock) {
// 	panic("unimplemented")
// }

func TestCourseUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(CourseUseCaseTestSuite))
}

func (suite *CourseUseCaseTestSuite) TestAddCourseSuccess() {
	payload := dto.CourseRequestDto{
		CourseName:  "Math",
		Description: "Advanced Mathematics",
	}

	expectedCourse := model.Course{
		CourseName:  payload.CourseName,
		Description: payload.Description,
	}

	suite.crm.On("Create", mock.AnythingOfType("model.Course")).Return(expectedCourse, nil)

	addedCourse, err := suite.cu.AddCourse(payload)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedCourse, addedCourse)
	suite.crm.AssertExpectations(suite.T())
}

func (suite *CourseUseCaseTestSuite) TestAddCourseError() {
	payload := dto.CourseRequestDto{
		CourseName:  "Math",
		Description: "Advanced Mathematics",
	}

	expectedError := errors.New("some error")

	suite.crm.On("Create", mock.AnythingOfType("model.Course")).Return(model.Course{}, expectedError)

	addedCourse, err := suite.cu.AddCourse(payload)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.Course{}, addedCourse)
	suite.crm.AssertExpectations(suite.T())
}

func (suite *CourseUseCaseTestSuite) TestFindCourseByIDSuccess() {
	targetID := "some-id"

	expectedCourse := model.Course{
		CourseName:  "Math",
		Description: "Advanced Mathematics",
	}

	suite.crm.On("GetById", targetID).Return(expectedCourse, nil)

	resultCourse, err := suite.cu.FindCourseByID(targetID)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedCourse, resultCourse)
	suite.crm.AssertExpectations(suite.T())
}

func (suite *CourseUseCaseTestSuite) TestFindCourseByIDError() {
	invalidID := "non-existent-id"

	expectedError := errors.New("data not found")
	suite.crm.On("GetById", invalidID).Return(model.Course{}, expectedError)

	resultCourse, err := suite.cu.FindCourseByID(invalidID)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.Course{}, resultCourse)
	suite.crm.AssertExpectations(suite.T())
}

func (suite *CourseUseCaseTestSuite) TestGetAllCourse_Success() {
	expectedCourses := []model.Course{
		{
			CourseID:    "1",
			CourseName:  "Introduction to Golang",
			Description: "anu anu anu",
		},
	}
	suite.crm.On("GetAll").Return(expectedCourses, nil)

	courses, err := suite.cu.GetAllCourse()

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedCourses, courses)

	suite.crm.AssertExpectations(suite.T())
}

func (suite *CourseUseCaseTestSuite) TestGetAllCourse_ErrorFromRepository() {
	suite.crm.On("GetAll").Return([]model.Course{}, errors.New("repository error"))

	courses, err := suite.cu.GetAllCourse()

	assert.Error(suite.T(), err)
	assert.Empty(suite.T(), courses)

	suite.crm.AssertExpectations(suite.T())
}

func (suite *CourseUseCaseTestSuite) TestUpdateCourseSuccess() {
	targetID := "some-id"

	updatePayload := dto.CourseRequestDto{
		CourseName:  "Updated Math",
		Description: "Updated Advanced Mathematics",
	}

	updatedCourse := model.Course{
		CourseName:  updatePayload.CourseName,
		Description: updatePayload.Description,
	}

	suite.crm.On("Update", mock.AnythingOfType("model.Course"), targetID).Return(updatedCourse, nil)

	resultCourse, err := suite.cu.UpdateCourse(updatePayload, targetID)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), updatedCourse, resultCourse)
	suite.crm.AssertExpectations(suite.T())
}

func (suite *CourseUseCaseTestSuite) TestUpdateCourseError() {
	targetID := "some-id"

	updatePayload := dto.CourseRequestDto{
		CourseName: "Updated Math",
	}

	expectedError := errors.New("failed update data by id : SomeError")
	suite.crm.On("Update", mock.AnythingOfType("model.Course"), targetID).Return(model.Course{}, expectedError)

	_, err := suite.cu.UpdateCourse(updatePayload, targetID)

	assert.Error(suite.T(), err)
	assert.EqualError(suite.T(), err, "failed to Update Course : failed update data by id : SomeError")
	suite.crm.AssertExpectations(suite.T())
}

func (suite *CourseUseCaseTestSuite) TestDeleteCourseSuccess() {
	targetID := "some-id"

	deletedCourse := model.Course{
		CourseName:  "Math",
		Description: "Advanced Mathematics",
	}

	suite.crm.On("Delete", targetID).Return(deletedCourse, nil)

	resultCourse, err := suite.cu.DeleteCourse(targetID)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), deletedCourse, resultCourse)
	suite.crm.AssertExpectations(suite.T())
}

func (suite *CourseUseCaseTestSuite) TestDeleteCourseError() {
	targetID := "some-id"

	expectedError := errors.New("failed to delete data")
	suite.crm.On("Delete", targetID).Return(model.Course{}, expectedError)

	resultCourse, err := suite.cu.DeleteCourse(targetID)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.Course{}, resultCourse)
	suite.crm.AssertExpectations(suite.T())
}
