package usecase

import (
	"errors"
	repomock "final-project-kelompok-1/mock/repo_mock"
	"final-project-kelompok-1/model"
	"final-project-kelompok-1/model/dto"

	"fmt"
	"time"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserUseCaseTestSuite struct {
	suite.Suite
	urm *repomock.UserRepoMock
	uu  UserUseCase
}

func (suite *UserUseCaseTestSuite) SetupTest() {
	suite.urm = new(repomock.UserRepoMock)
	suite.uu = NewUserUseCase(suite.urm)
}

// func NewUserUseCase(userRepoMock *repomock.UserRepoMock) {
// 	panic("unimplemented")
// }

func TestUserUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUseCaseTestSuite))
}

func (suite *UserUseCaseTestSuite) TestAddUserSuccess() {
	payload := dto.UserRequestDto{
		Fullname: "yanto",
		Role:     "Admin",
		Email:    "yanto@gmail.com",
		Password: "12345",
	}

	expectedUser := model.Users{
		Fullname: payload.Fullname,
		Role:     payload.Role,
		Email:    payload.Email,
		Password: payload.Password,
	}

	suite.urm.On("Create", mock.AnythingOfType("model.Users")).Return(expectedUser, nil)

	createdUser, err := suite.uu.AddUser(payload)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedUser, createdUser)
	suite.urm.AssertExpectations(suite.T())

}

func (suite *UserUseCaseTestSuite) TestAddUserNegative() {
	invalidPayload := dto.UserRequestDto{
		Fullname: "yanto",
		Role:     "Admin",
		Email:    "yanto@gmail.com",
		Password: "12345",
	}

	suite.urm.On("Create", mock.AnythingOfType("model.Users")).Return(model.Users{}, fmt.Errorf("Create should not be called")).Once()

	createdUser, err := suite.uu.AddUser(invalidPayload)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.Users{}, createdUser)

	suite.urm.AssertExpectations(suite.T())
}

func (suite *UserUseCaseTestSuite) TestFindStudentByIDPositive() {
	targetID := "some-id"

	expectUser := model.Users{
		Fullname: "yanto",
		Role:     "Admin",
		Email:    "yanto@gmail.com",
		Password: "12345",
	}

	suite.urm.On("GetById", targetID).Return(expectUser, nil)

	resultUser, err := suite.uu.FindUserByID(targetID)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectUser, resultUser)

	suite.urm.AssertExpectations(suite.T())
}

func (suite *UserUseCaseTestSuite) TestFindStudentByIDNegative() {
	invalidID := "non-existent-id"

	expectedError := errors.New("data not found")
	suite.urm.On("GetById", invalidID).Return(model.Users{}, expectedError)

	resultStudent, err := suite.uu.FindUserByID(invalidID)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.Users{}, resultStudent)

	suite.urm.AssertExpectations(suite.T())
}

func (suite *UserUseCaseTestSuite) TestGetAllUser_Success() {
	expectedUsers := []model.Users{
		{
			UserID:    "1",
			Fullname:  "John Doe",
			Role:      "user",
			Email:     "john.doe@example.com",
			Password:  "hashed_password",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			IsDeleted: false,
		},
	}
	suite.urm.On("GetAll").Return(expectedUsers, nil)

	users, err := suite.uu.GetAllUser()

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedUsers, users)

	suite.urm.AssertExpectations(suite.T())
}

func (suite *UserUseCaseTestSuite) TestGetAllUser_ErrorFromRepository() {
	expectedErr := errors.New("repository error")
	suite.urm.On("GetAll").Return([]model.Users{}, expectedErr)

	users, err := suite.uu.GetAllUser()

	assert.Error(suite.T(), err)
	assert.EqualError(suite.T(), err, fmt.Sprintf("failed to find all data : %s", expectedErr.Error()))
	assert.Nil(suite.T(), users)

}

func (suite *UserUseCaseTestSuite) TestUpdateStudentSuccess() {
	targetID := "some-id"

	updatePayload := dto.UserRequestDto{
		Fullname: "yanto",
		Role:     "Admin",
		Email:    "yanto@gmail.com",
		Password: "12345",
	}

	updatedUser := model.Users{
		Fullname: "yanto",
		Role:     "Admin",
		Email:    "yanto@gmail.com",
		Password: "12345",
	}

	suite.urm.On("Update", mock.AnythingOfType("model.Users"), targetID).Return(updatedUser, nil)

	resultUser, err := suite.uu.UpdateUser(updatePayload, targetID)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), updatedUser, resultUser)

	suite.urm.AssertExpectations(suite.T())
}

func (suite *UserUseCaseTestSuite) TestUpdateStudentNegative() {
	targetID := "some-id"

	updatePayload := dto.UserRequestDto{
		Fullname: "yanto",
		Role:     "Admin",
		Email:    "yanto@gmail.com",
		Password: "12345",
	}

	expectedError := errors.New("failed update data by id : SomeError")
	suite.urm.On("Update", mock.AnythingOfType("model.Users"), targetID).Return(model.Users{}, expectedError)

	_, err := suite.uu.UpdateUser(updatePayload, targetID)

	assert.Error(suite.T(), err)
	assert.EqualError(suite.T(), err, "failed update data by id : failed update data by id : SomeError")

	suite.urm.AssertExpectations(suite.T())

}

func (suite *UserUseCaseTestSuite) TestDeleteStudentSuccess() {
	targetID := "some-id"

	deletedUser := model.Users{
		Fullname: "yanto",
		Role:     "Admin",
		Email:    "yanto@gmail.com",
		Password: "12345",
	}

	suite.urm.On("Delete", targetID).Return(deletedUser, nil)

	resultUser, err := suite.uu.DeleteUser(targetID)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), deletedUser, resultUser)

	suite.urm.AssertExpectations(suite.T())
}

func (suite *UserUseCaseTestSuite) TestDeleteStudentNegative() {
	targetID := "some-id"

	expectedError := errors.New("failed to delete data")
	suite.urm.On("Delete", targetID).Return(model.Users{}, expectedError)

	resultUser, err := suite.uu.DeleteUser(targetID)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.Users{}, resultUser)

	suite.urm.AssertExpectations(suite.T())
}

func (suite *UserUseCaseTestSuite) TestRegisterNewUser_Success() {
	payload := dto.UserRequestDto{
		Fullname: "yanto",
		Role:     "Admin",
		Email:    "yanto@gmail.com",
		Password: "12345",
	}

	expectedUser := model.Users{
		Fullname: payload.Fullname,
		Role:     payload.Role,
		Email:    payload.Email,
		Password: payload.Password,
	}

	suite.urm.On("Create", mock.AnythingOfType("model.Users")).Return(expectedUser, nil)

	createdUser, err := suite.uu.AddUser(payload)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedUser, createdUser)
	suite.urm.AssertExpectations(suite.T())
}

func (suite *UserUseCaseTestSuite) TestRegisterNewUser_InvalidRole() {
	invalidPayload := model.Users{
		Fullname: "yanto",
		Role:     "Admin",
		Email:    "yanto@gmail.com",
		Password: "12345",
	}

	suite.urm.On("Create", mock.AnythingOfType("model.Users")).Return(model.Users{}, fmt.Errorf("Create should not be called")).Once()

	createdUser, err := suite.uu.RegisterNewUser(invalidPayload)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.Users{}, createdUser)

	// suite.urm.AssertExpectations(suite.T())
}
