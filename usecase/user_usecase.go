package usecase

import (
	"final-project-kelompok-1/model"
	"final-project-kelompok-1/model/dto"
	"final-project-kelompok-1/repository"
	"fmt"
)

type UserUseCase interface {
	AddUser(payload dto.UserRequestDto) (model.Users, error)
	FindUserByID(id string) (model.Users, error)
	UpdateUser(payload dto.UserRequestDto, id string) (model.Users, error)
	DeleteUser(id string) (model.Users, error)
}

type userUseCase struct {
	repo repository.UserRepositpry
}

func (u *userUseCase) AddUser(payload dto.UserRequestDto) (model.Users, error) {
	newUser := model.Users{
		Fullname: payload.Fullname,
		Role:     payload.Role,
		Email:    payload.Role,
		Password: payload.Password,
	}

	addedUser, err := u.repo.Create(newUser)

	if err != nil {
		fmt.Println("Error inserting user di usecase : ", err)
		return model.Users{}, fmt.Errorf("failed Add User : %s", err.Error())
	}

	return addedUser, nil

}

func (u *userUseCase) FindUserByID(id string) (model.Users, error) {
	userWithId, err := u.repo.GetById(id)

	if err != nil {
		fmt.Println("Error inserting user di usecafe : ", err)
		return model.Users{}, fmt.Errorf("failed to find data : %s", err.Error())
	}

	return userWithId, nil

}

func (u *userUseCase) UpdateUser(payload dto.UserRequestDto, id string) (model.Users, error) {
	newUser := model.Users{
		Fullname: payload.Fullname,
		Role:     payload.Role,
		Email:    payload.Email,
		Password: payload.Password,
	}

	updatedUser, err := u.repo.Update(newUser, id)

	if err != nil {
		return model.Users{}, fmt.Errorf("failed update data by id : %s", err.Error())
	}

	return updatedUser, nil

}

func (u *userUseCase) DeleteUser(id string) (model.Users, error) {
	deletedUser, err := u.repo.Delete(id)

	if err != nil {
		return model.Users{}, fmt.Errorf("failed to delete data : %s", err.Error())
	}

	return deletedUser, nil
}

func NewUserUseCase(repo repository.UserRepositpry) UserUseCase {
	return &userUseCase{repo: repo}
}
