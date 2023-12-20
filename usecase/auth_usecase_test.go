package usecase

import (
	"errors"
	usecaseMock "final-project-kelompok-1/mock/usecase_mock"
	"final-project-kelompok-1/model"
	"final-project-kelompok-1/model/dto"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AuthUseCaseTestSuite struct {
	suite.Suite
	ucm     *usecaseMock.MockUserUseCase
	scm     *usecaseMock.MockStudentUseCase
	jwtMock *usecaseMock.MockJwtToken
	au      AuthUseCase
}

func (suite *AuthUseCaseTestSuite) SetupTest() {
	suite.ucm = new(usecaseMock.MockUserUseCase)
	suite.scm = new(usecaseMock.MockStudentUseCase)
	suite.jwtMock = new(usecaseMock.MockJwtToken)
	suite.au = NewAuthUseCase(suite.ucm, suite.scm, suite.jwtMock)
}

func TestAuthUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(AuthUseCaseTestSuite))
}

func (suite *AuthUseCaseTestSuite) TestRegister_Success() {
	userPayload := model.Users{
		Fullname: "yanto_balap",
		Role:     "admin",
		Email:    "cuankibalap@gmail.com",
		Password: "12345",
	}

	suite.ucm.On("RegisterNewUser", userPayload).Return(userPayload, nil)

	result, err := suite.au.Register(userPayload)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), userPayload, result)

	suite.ucm.AssertExpectations(suite.T())
}

func (suite *AuthUseCaseTestSuite) TestLogin_Success() {
	authRequest := dto.AuthRequestDto{
		Email:    "test@example.com",
		Password: "password123",
	}

	expectedUser := model.Users{
		Fullname: "yanto_balap",
		Role:     "admin",
		Email:    "cuankibalap@gmail.com",
		Password: "12345",
	}

	expectedToken := dto.AuthResponseDto{
		Token: `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.
		eyJpc3MiOiJhZG1pbjFAZ21haWwuY29tIiwiZXhwIjoxNzAyNzUxODU2LCJpYXQiOjE3MDI3NDgyNTYsInVzZXJJZCI6IjA2OTA3YTQwL
		TJjNmYtNGE5YS1hNzU3LWRlNjUxNmIzYTdmZiIsInJvbGUiOiJzdHVkZW50Iiwic2VydmljZXMiOm51bGx9.
		KPSuzAm7iqlNbe65T5bogYEDGCqx8btn4DD3aljfKpA`,
	}

	suite.ucm.On("FindByUsernamePassword", authRequest.Email, authRequest.Password).Return(expectedUser, nil)
	suite.jwtMock.On("GenerateToken", expectedUser).Return(expectedToken, nil)

	result, err := suite.au.Login(authRequest)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedToken, result)

	suite.ucm.AssertExpectations(suite.T())
	suite.jwtMock.AssertExpectations(suite.T())
}

func (suite *AuthUseCaseTestSuite) TestLogin_ErrorFromUserRepo() {
	authRequest := dto.AuthRequestDto{
		Email:    "test@example.com",
		Password: "password123",
	}

	expectedError := errors.New("user repo error")

	suite.ucm.On("FindByUsernamePassword", authRequest.Email, authRequest.Password).Return(model.Users{}, expectedError)

	result, err := suite.au.Login(authRequest)

	assert.Error(suite.T(), err)
	assert.EqualError(suite.T(), err, expectedError.Error())
	assert.Equal(suite.T(), dto.AuthResponseDto{}, result)

	suite.ucm.AssertExpectations(suite.T())
}

func (suite *AuthUseCaseTestSuite) TestLoginStudent_Success() {
	authRequest := dto.AuthRequestDto{
		Email:    "test@example.com",
		Password: "password123",
	}

	expectedStudent := model.Student{
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

	expectedToken := dto.AuthResponseDto{
		Token: `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.
		eyJpc3MiOiJhZG1pbjFAZ21haWwuY29tIiwiZXhwIjoxNzAyNzUxODU2LCJpYXQiOjE3MDI3NDgyNTYs
		InVzZXJJZCI6IjA2OTA3YTQwLTJjNmYtNGE5YS1hNzU3LWRlNjUxNmIzYTdmZiIsInJvbGUiOiJzdHVkZW50Iiwic2VydmljZXMiOm51bGx9.
		KPSuzAm7iqlNbe65T5bogYEDGCqx8btn4DD3aljfKpA`,
	}

	suite.scm.On("FindByEmailPassword", authRequest.Email, authRequest.Password).Return(expectedStudent, nil)
	suite.jwtMock.On("GenerateTokenStudent", expectedStudent).Return(expectedToken, nil)

	result, err := suite.au.LoginStudent(authRequest)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedToken, result)

	suite.scm.AssertExpectations(suite.T())
	suite.jwtMock.AssertExpectations(suite.T())
}

func (suite *AuthUseCaseTestSuite) TestLoginStudent_ErrorFromStudentRepo() {
	authRequest := dto.AuthRequestDto{
		Email:    "test@example.com",
		Password: "password123",
	}

	expectedError := errors.New("student repo error")

	suite.scm.On("FindByEmailPassword", authRequest.Email, authRequest.Password).Return(model.Student{}, expectedError)

	result, err := suite.au.LoginStudent(authRequest)

	assert.Error(suite.T(), err)
	assert.EqualError(suite.T(), err, expectedError.Error())
	assert.Equal(suite.T(), dto.AuthResponseDto{}, result)

	suite.scm.AssertExpectations(suite.T())
}

func (suite *AuthUseCaseTestSuite) TearDownTest() {
	suite.ucm.AssertExpectations(suite.T())
	suite.scm.AssertExpectations(suite.T())
	suite.jwtMock.AssertExpectations(suite.T())
}
