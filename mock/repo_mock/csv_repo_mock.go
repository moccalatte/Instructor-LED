package repo_mock

import "github.com/stretchr/testify/mock"

type CsvRepoMock struct{
	mock.Mock
}


func(c *CourseRepoMock)CsvStart()([]string, error){
	args:= c.Called()
	return args.Get(0).([]string), args.Error(1)
}