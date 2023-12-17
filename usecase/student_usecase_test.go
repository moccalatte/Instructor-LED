package usecase

import (
	"errors"
	repomock "final-project-kelompok-1/mock/repo_mock"
	"final-project-kelompok-1/model"
	"final-project-kelompok-1/model/dto"
	"fmt"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type StudentUseCaseTestSuite struct {
	suite.Suite
	srm *repomock.StudentRepoMock
	su  StudentUseCase
}

func (suite *StudentUseCaseTestSuite) SetupTest() {
	suite.srm = new(repomock.StudentRepoMock)
	suite.su = NewStudentUseCase(suite.srm)
}

func TestStudentUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(StudentUseCaseTestSuite))
}

func (suite *StudentUseCaseTestSuite) TestAddStudentSuccess() {
	payload := dto.StudentRequestDto{
		Fullname:    "John Doe",
		BirthDate:   "1990-01-01",
		BirthPlace:  "City",
		Address:     "Street",
		Education:   "Bachelor",
		Institution: "University",
		Job:         "Engineer",
		Email:       "john@example.com",
		Password:    "password123",
	}

	expectedStudent := model.Student{
		Fullname:    payload.Fullname,
		BirthDate:   payload.BirthDate,
		BirthPlace:  payload.BirthPlace,
		Address:     payload.Address,
		Education:   payload.Education,
		Institution: payload.Institution,
		Job:         payload.Job,
		Email:       payload.Email,
		Password:    payload.Password,
	}

	suite.srm.On("Create", mock.AnythingOfType("model.Student")).Return(expectedStudent, nil)

	createdStudent, err := suite.su.AddStudent(payload)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedStudent, createdStudent)
	suite.srm.AssertExpectations(suite.T())

}

func (suite *StudentUseCaseTestSuite) TestAddStudentNegative() {
	invalidPayload := dto.StudentRequestDto{
		Fullname:    "John Doe",
		BirthDate:   "1990-01-01",
		BirthPlace:  "City",
		Address:     "Street",
		Education:   "Bachelor",
		Institution: "University",
		Job:         "Engineer",
		Email:       "invalid-email",
		Password:    "password123",
	}

	suite.srm.On("Create", mock.AnythingOfType("model.Student")).Return(model.Student{}, fmt.Errorf("Create should not be called")).Once()

	createdStudent, err := suite.su.AddStudent(invalidPayload)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.Student{}, createdStudent)

	suite.srm.AssertExpectations(suite.T())
}

func (suite *StudentUseCaseTestSuite) TestFindStudentByIDPositive() {
	targetID := "some-id"

	expectedStudent := model.Student{
		Fullname:    "John Doe",
		BirthDate:   "1990-01-01",
		BirthPlace:  "City",
		Address:     "Street",
		Education:   "Bachelor",
		Institution: "University",
		Job:         "Engineer",
		Email:       "invalid-email",
		Password:    "password123",
	}

	suite.srm.On("GetById", targetID).Return(expectedStudent, nil)

	resultStudent, err := suite.su.FindStudentByID(targetID)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedStudent, resultStudent)

	suite.srm.AssertExpectations(suite.T())
}

func (suite *StudentUseCaseTestSuite) TestFindStudentByIDNegative() {
	invalidID := "non-existent-id"

	expectedError := errors.New("data not found")
	suite.srm.On("GetById", invalidID).Return(model.Student{}, expectedError)

	resultStudent, err := suite.su.FindStudentByID(invalidID)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.Student{}, resultStudent)

	suite.srm.AssertExpectations(suite.T())
}

func (suite *StudentUseCaseTestSuite) TestGetAllStudent_Success() {
	// Mock repository response
	expectedStudents := []model.Student{
		{
			StudentID:   "1",
			Fullname:    "John Doe",
			BirthDate:   "2000-01-01",
			BirthPlace:  "City",
			Address:     "Street 123",
			Education:   "Bachelor",
			Institution: "University",
			Job:         "Developer",
			Email:       "john.doe@example.com",
			Password:    "hashed_password",
		},
	}
	suite.srm.On("GetAll").Return(expectedStudents, nil)

	// Call the method being tested
	students, err := suite.su.GetAllStudent()

	// Assertions
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedStudents, students)

	// Assert that the expected method was called
	suite.srm.AssertExpectations(suite.T())
}

func (suite *StudentUseCaseTestSuite) TestGetAllStudent_ErrorFromRepository() {
	// Mock repository response with an error
	suite.srm.On("GetAll").Return([]model.Student{}, errors.New("repository error"))

	// Call the method being tested
	students, err := suite.su.GetAllStudent()

	// Assertions
	assert.Error(suite.T(), err)
	assert.Empty(suite.T(), students)

	// Assert that the expected method was called
	suite.srm.AssertExpectations(suite.T())
}

func (suite *StudentUseCaseTestSuite) TestUpdateStudentSuccess() {
	targetID := "some-id"

	updatePayload := dto.StudentRequestDto{
		Fullname:    "John Doe",
		BirthDate:   "1990-01-01",
		BirthPlace:  "City",
		Address:     "Street",
		Education:   "Bachelor",
		Institution: "University",
		Job:         "Engineer",
		Email:       "invalid-email",
		Password:    "password123",
	}

	updatedStudent := model.Student{
		Fullname:    "John Doe",
		BirthDate:   "1990-01-01",
		BirthPlace:  "City",
		Address:     "Street",
		Education:   "Bachelor",
		Institution: "University",
		Job:         "Engineer",
		Email:       "invalid-email",
		Password:    "password123",
	}

	suite.srm.On("Update", mock.AnythingOfType("model.Student"), targetID).Return(updatedStudent, nil)

	resultStudent, err := suite.su.UpdateStudent(updatePayload, targetID)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), updatedStudent, resultStudent)

	suite.srm.AssertExpectations(suite.T())
}

func (suite *StudentUseCaseTestSuite) TestUpdateStudentNegative() {
	targetID := "some-id"

	updatePayload := dto.StudentRequestDto{
		Fullname:  "Updated Name",
		BirthDate: "1990-01-01",
	}

	expectedError := errors.New("failed update data by id : SomeError")
	suite.srm.On("Update", mock.AnythingOfType("model.Student"), targetID).Return(model.Student{}, expectedError)

	_, err := suite.su.UpdateStudent(updatePayload, targetID)

	assert.Error(suite.T(), err)
	assert.EqualError(suite.T(), err, "failed update data by id : failed update data by id : SomeError")

	suite.srm.AssertExpectations(suite.T())

}

func (suite *StudentUseCaseTestSuite) TestDeleteStudentSuccess() {
	targetID := "some-id"

	deletedStudent := model.Student{
		Fullname:    "John Doe",
		BirthDate:   "1990-01-01",
		BirthPlace:  "City",
		Address:     "Street",
		Education:   "Bachelor",
		Institution: "University",
		Job:         "Engineer",
		Email:       "invalid-email",
		Password:    "password123",
	}

	suite.srm.On("Delete", targetID).Return(deletedStudent, nil)

	resultStudent, err := suite.su.DeleteStudent(targetID)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), deletedStudent, resultStudent)

	suite.srm.AssertExpectations(suite.T())
}

func (suite *StudentUseCaseTestSuite) TestDeleteStudentNegative() {
	targetID := "some-id"

	expectedError := errors.New("failed to delete data")
	suite.srm.On("Delete", targetID).Return(model.Student{}, expectedError)

	resultStudent, err := suite.su.DeleteStudent(targetID)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.Student{}, resultStudent)

	suite.srm.AssertExpectations(suite.T())
}
