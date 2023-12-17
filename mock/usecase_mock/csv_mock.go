package usecase_mock

import (
	"final-project-kelompok-1/config"

	"github.com/stretchr/testify/mock"
)

type MockCvsCommon struct {
	mock.Mock
}

func (m *MockCvsCommon) CreateFile() {
	m.Called()
}

func (m *MockCvsCommon) WriterData(data string) error {
	args := m.Called(data)
	return args.Error(0)
}

type MockCsvCommon struct {
	mock.Mock
}

func (m *MockCsvCommon) CreateFile() {
	m.Called()
}

func (m *MockCsvCommon) WriterData(data string) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockCsvCommon) ConvertArrayToString(array []string) string {
	args := m.Called(array)
	return args.String(0)
}

func NewMockCvsCommon() *MockCvsCommon {
	return &MockCvsCommon{}
}

func NewMockCsvCommon() *MockCsvCommon {
	return &MockCsvCommon{}
}

func NewMockCsvCommonWithConfig(cfg config.CsvFileConfig) *MockCsvCommon {
	mockCsvCommon := &MockCsvCommon{}
	mockCsvCommon.On("CreateFile").Return()
	mockCsvCommon.On("WriterData", mock.Anything).Return(nil)
	mockCsvCommon.On("ConvertArrayToString", mock.Anything).Return("")
	return mockCsvCommon
}
