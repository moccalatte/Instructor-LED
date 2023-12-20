package repo_mock

import (
	"github.com/stretchr/testify/mock"
)

type CsvRepositoryMock struct {
	mock.Mock
}

func (m *CsvRepositoryMock) CsvStart() ([]string, error) {
	args := m.Called()

	// Check if the first return argument is not nil
	if arg0 := args.Get(0); arg0 != nil {
		return arg0.([]string), args.Error(1)
	}

	// If the first return argument is nil, return an empty slice of strings
	return []string{}, args.Error(1)
}
