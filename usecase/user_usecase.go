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
	UpdateUser(payload dto.UserRequestDto, id string) (model.Users, error)
	DeleteUser(id string) (model.Users, error)
	FindByFullnamePassword(fullname string, password string) (model.Users, error)

	// INI HANYA TEST
	RegisterNewUser(payload model.Users) (model.Users, error)
}

type userUseCase struct {
	repo repository.UserRepositpry
}

func (u *userUseCase) AddUser(payload dto.UserRequestDto) (model.Users, error) {
	newUser := model.Users{
		Fullname: payload.Fullname,
		Role:     payload.Role,
		Email:    payload.Email,
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
		return model.Users{}, fmt.Errorf("failed to Update data : %s", err.Error())
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

func (u *userUseCase) FindByFullnamePassword(fullname string, password string) (model.Users, error) {
	user, err := u.repo.GetByFullname(fullname)
	if err != nil {
		return model.Users{}, errors.New("Invalid fullname or password")
	}

	// compare password
	if err := common.ComparePasswordHash(user.Password, password); err != nil {
		return model.Users{}, err
	}

	user.Password = ""
	return user, nil
}

// INI HANYA TEST

func (u *userUseCase) RegisterNewUser(payload model.Users) (model.Users, error) {
	// biasanya ada pengecekan jika email / username sudah ada di database, maka tidak boleh diinsert lagi

	// cek role
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

// END

func NewUserUseCase(repo repository.UserRepositpry) UserUseCase {
	return &userUseCase{repo: repo}
}
