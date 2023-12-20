package usecase

import (
	"errors"
	"final-project-kelompok-1/model"
	"final-project-kelompok-1/model/dto"
	"final-project-kelompok-1/repository"
	"final-project-kelompok-1/utils/common"
	"fmt"
)

type UserUseCase interface {
	AddUser(payload dto.UserRequestDto) (model.Users, error)
	FindUserByID(id string) (model.Users, error)
	GetAllUser() ([]model.Users, error)
	UpdateUser(payload dto.UserRequestDto, id string) (model.Users, error)
	DeleteUser(id string) (model.Users, error)
	RegisterNewUser(payload model.Users) (model.Users, error)
	FindByUsernamePassword(email string, password string) (model.Users, error)
	// GetByUsername(username string) (model.Users, error)
}

type userUseCase struct {
	repo repository.UserRepository
}

func (u *userUseCase) AddUser(payload dto.UserRequestDto) (model.Users, error) {
	newUser := model.Users{
		Fullname: payload.Fullname,
		Role:     payload.Role,
		Email:    payload.Email,
		Password: payload.Password,
	}

	newPassword, err := common.GeneratePasswordHash(payload.Password)
	if err != nil {
		return model.Users{}, err
	}

	newUser.Password = newPassword

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
		return model.Users{}, fmt.Errorf("failed to get User data : %s", err.Error())
	}

	return userWithId, nil

}

func (u *userUseCase) GetAllUser() ([]model.Users, error) {
	var sliceUser []model.Users
	userData, err := u.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to find all data : %s", err.Error())
	}
	return append(sliceUser, userData...), nil

}

func (u *userUseCase) UpdateUser(payload dto.UserRequestDto, id string) (model.Users, error) {
	newUser := model.Users{
		Fullname: payload.Fullname,
		Role:     payload.Role,
		Email:    payload.Email,
		Password: payload.Password,
	}

	newPassword, err := common.GeneratePasswordHash(payload.Password)
	if err != nil {
		return model.Users{}, err
	}

	newUser.Password = newPassword

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

func (u *userUseCase) RegisterNewUser(payload model.Users) (model.Users, error) {
	if !payload.IsValidRole() {
		return model.Users{}, errors.New("invalid role, role must be admin or employee")
	}

	newPassword, err := common.GeneratePasswordHash(payload.Password)
	if err != nil {
		return model.Users{}, err
	}

	payload.Password = newPassword
	return u.repo.Create(payload)
}

func (u *userUseCase) FindByUsernamePassword(email string, password string) (model.Users, error) {
	user, err := u.repo.GetByUsername(email)
	fmt.Println(user)
	if err != nil {
		fmt.Println("Error in usecase : ", err.Error())
		return model.Users{}, errors.New("invalid email or password")
	}

	if err := common.ComparePasswordHash(user.Password, password); err != nil {
		return model.Users{}, err
	}

	user.Password = ""
	return user, nil
}
func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{repo: repo}
}
