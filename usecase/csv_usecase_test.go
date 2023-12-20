package usecase

import (
	"errors"
	"final-project-kelompok-1/mock/repo_mock"
	"final-project-kelompok-1/mock/usecase_mock"

	"final-project-kelompok-1/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CsvUseCaseTestSuite struct {
	suite.Suite
	seMock *usecase_mock.MockSessionUseCase
	atMock *usecase_mock.MockAttendanceUseCase
	snMock *usecase_mock.MockStudentUseCase
	tnMock *usecase_mock.MockUserUseCase
	quMock *usecase_mock.MockQuestionUseCase
	cnMock *usecase_mock.MockCourseUseCase
	csMock *usecase_mock.MockCvsCommon
	repo   *repo_mock.CsvRepositoryMock
	cu     CsvUseCase
}

func (suite *CsvUseCaseTestSuite) SetupTest() {
	suite.seMock = new(usecase_mock.MockSessionUseCase)
	suite.atMock = new(usecase_mock.MockAttendanceUseCase)
	suite.snMock = new(usecase_mock.MockStudentUseCase)
	suite.tnMock = new(usecase_mock.MockUserUseCase)
	suite.quMock = new(usecase_mock.MockQuestionUseCase)
	suite.cnMock = new(usecase_mock.MockCourseUseCase)
	suite.csMock = new(usecase_mock.MockCvsCommon)
	suite.repo = new(repo_mock.CsvRepositoryMock)
	suite.cu = NewCsvUsecase(suite.seMock, suite.atMock, suite.snMock, suite.tnMock, suite.quMock, suite.cnMock, suite.csMock, suite.repo)
}

func TestCsvUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(CsvUseCaseTestSuite))
}

func (suite *CsvUseCaseTestSuite) TestWriteCsv_Success() {
	startPoint := []string{"sessionID1", "sessionID1"}

	suite.repo.On("CsvStart").Return(startPoint, nil)
	suite.seMock.On("FindSessionById", "sessionID1").Return(model.Session{}, nil)
	suite.tnMock.On("FindUserByID", "").Return(model.Users{}, errors.New("trainer not found"))
	// suite.atMock.On("FindAttendanceBySessionId", "").Return(model.Attendance{}, errors.New("attendance not found"))
	// suite.snMock.On("FindStudentByID", "").Return(model.Student{}, errors.New("student not found"))
	// suite.quMock.On("FindQuestionByStudentId", "").Return(model.Question{}, errors.New("question not found"))
	// suite.cnMock.On("FindCourseByID", "").Return(model.Course{}, errors.New("course not found"))
	suite.csMock.On("CreateFile").Return(nil)
	// suite.csMock.On("WriterData", "").Return(nil)

	result, err := suite.cu.WriteCsv()

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), startPoint, result)

	suite.repo.AssertExpectations(suite.T())
	suite.seMock.AssertExpectations(suite.T())
	suite.tnMock.AssertExpectations(suite.T())
	suite.atMock.AssertExpectations(suite.T())
	suite.snMock.AssertExpectations(suite.T())
	suite.quMock.AssertExpectations(suite.T())
	suite.cnMock.AssertExpectations(suite.T())
	suite.csMock.AssertExpectations(suite.T())
}

func (suite *CsvUseCaseTestSuite) TearDownTest() {
	suite.seMock.AssertExpectations(suite.T())
	suite.atMock.AssertExpectations(suite.T())
	suite.snMock.AssertExpectations(suite.T())
	suite.tnMock.AssertExpectations(suite.T())
	suite.quMock.AssertExpectations(suite.T())
	suite.cnMock.AssertExpectations(suite.T())
	suite.csMock.AssertExpectations(suite.T())
	suite.repo.AssertExpectations(suite.T())
}
