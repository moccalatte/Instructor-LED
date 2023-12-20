package usecase_mock

import (
	"final-project-kelompok-1/model"
	"final-project-kelompok-1/model/dto"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/mock"
)

type MockJwtToken struct {
	mock.Mock
}

func (m *MockJwtToken) GenerateToken(user model.Users) (dto.AuthResponseDto, error) {
	args := m.Called(user)
	return args.Get(0).(dto.AuthResponseDto), args.Error(1)
}

func (m *MockJwtToken) GenerateTokenStudent(student model.Student) (dto.AuthResponseDto, error) {
	args := m.Called(student)
	return args.Get(0).(dto.AuthResponseDto), args.Error(1)
}

func (m *MockJwtToken) VerifyToken(tokenString string) (jwt.MapClaims, error) {
	args := m.Called(tokenString)
	return args.Get(0).(jwt.MapClaims), args.Error(1)
}

// RefreshToken adalah metode untuk merefresh token JWT mock
func (m *MockJwtToken) RefreshToken(oldTokenString string) (dto.AuthResponseDto, error) {
	args := m.Called(oldTokenString)
	return args.Get(0).(dto.AuthResponseDto), args.Error(1)
}
