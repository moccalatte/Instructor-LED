package repository

import (
	"database/sql"
	"fmt"

	"final-project-kelompok-1/model"
	"final-project-kelompok-1/utils/common"
)

type UserRepositpry interface {
	Create(payload model.Users) (model.Users, error)
	GetById(id string) (model.Users, error)
	Update(payload model.Users, id string) (model.Users, error)
	Delete(id string) (model.Users, error)
}

type userRepository struct {
	db *sql.DB
}

func (u *userRepository) Create(payload model.Users) (model.Users, error) {
	tx, err := u.db.Begin()
	if err != nil {
		return model.Users{}, err
	}

	var user model.Users
	err = tx.QueryRow(common.CreateUser,
		payload.Fullname,
		payload.Role,
		payload.Email,
		payload.Password,
		payload.UpdatedAt,
		false,
	).Scan(
		&user.UserID,
		&user.Fullname,
		&user.Role,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.IsDeleted,
	)

	if err != nil {
		fmt.Println("Error di Repo user : ", err)
		return model.Users{}, tx.Rollback()
	}

	if err := tx.Commit(); err != nil {
		return model.Users{}, err
	}
	return user, nil

}

func (u *userRepository) GetById(id string) (model.Users, error) {
	var user model.Users
	err := u.db.QueryRow(common.GetUserById, id).Scan(
		&user.UserID,
		&user.Fullname,
		&user.Role,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.IsDeleted,
	)
	if err != nil {
		return model.Users{}, err
	}
	return user, nil
}

func (u *userRepository) Update(payload model.Users, id string) (model.Users, error) {
	tx, err := u.db.Begin()
	if err != nil {
		return model.Users{}, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	var user model.Users
	err = tx.QueryRow(common.UpdateUser,
		payload.Fullname,
		payload.Role,
		payload.Email,
		payload.Password,
		payload.UpdatedAt,
		false,
		id).Scan(
		&user.UserID,
		&user.Fullname,
		&user.Role,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.IsDeleted,
	)
	if err != nil {
		fmt.Println("Error inserting user di repo : ", err)
		return model.Users{}, tx.Rollback()
	}

	if err := tx.Commit(); err != nil {
		return model.Users{}, err
	}

	return user, nil
}

func (u *userRepository) Delete(id string) (model.Users, error) {
	tx, err := u.db.Begin()
	if err != nil {
		return model.Users{}, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	var user model.Users
	err = tx.QueryRow(common.DeleteUser,
		true,
		id).Scan(
		&user.UserID,
		&user.Fullname,
		&user.Role,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.IsDeleted,
	)
	if err != nil {
		return model.Users{}, tx.Rollback()
	}

	if err := tx.Commit(); err != nil {
		return model.Users{}, err
	}

	return user, nil
}

func NewUserRepository(db *sql.DB) UserRepositpry {
	return &userRepository{db: db}
}
