package usecase

import (
	"final-project-kelompok-1/model"
	"final-project-kelompok-1/model/dto"

	// "final-project-kelompok-1/model/dto"
	"final-project-kelompok-1/utils/common"
)

type AuthUseCase interface {
	Register(payload model.Users) (model.Users, error)
	Login(payload dto.AuthRequestDto) (dto.AuthResponseDto, error)
	LoginStudent(payload dto.AuthRequestDto) (dto.AuthResponseDto, error)
}

type authUseCase struct {
	uc       UserUseCase
	sc       StudentUseCase
	jwtToken common.JwtToken
}

func (a *authUseCase) Register(payload model.Users) (model.Users, error) {
	return a.uc.RegisterNewUser(payload)
}

func (a *authUseCase) Login(payload dto.AuthRequestDto) (dto.AuthResponseDto, error) {
	user, err := a.uc.FindByUsernamePassword(payload.Email, payload.Password)
	if err != nil {
		return dto.AuthResponseDto{}, err
	}

	token, err := a.jwtToken.GenerateToken(user)
	if err != nil {
		return dto.AuthResponseDto{}, err
	}

	return token, nil
}

func (a *authUseCase) LoginStudent(payload dto.AuthRequestDto) (dto.AuthResponseDto, error) {
	student, err := a.sc.FindByEmailPassword(payload.Email, payload.Password)
	if err != nil {
		return dto.AuthResponseDto{}, err
	}

	token, err := a.jwtToken.GenerateTokenStudent(student)
	if err != nil {
		return dto.AuthResponseDto{}, err
	}

	return token, nil
}

func NewAuthUseCase(uc UserUseCase, sc StudentUseCase, jwtToken common.JwtToken) AuthUseCase {
	return &authUseCase{uc: uc, sc: sc, jwtToken: jwtToken}
}
