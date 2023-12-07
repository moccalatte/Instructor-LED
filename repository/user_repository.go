package repository

import (
	"database/sql"

	"final-project-kelompok-1/model"
	"final-project-kelompok-1/utils/common"
)

type UserRepositpry interface {
	Create(payload model.Users) (model.Users, error)
	GetById(id int) (model.Users, error)
	Update(payload model.Users, id int) (model.Users, error)
	Delete(id int) (model.Users, error)
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
		payload.Name,
		payload.Role,
		payload.Email,
		payload.Password,
		true,
	).Scan(
		&user.UserID,
		&user.Name,
		&user.Role,
		&user.Email,
		&user.Password,
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

func (u *userRepository) GetById(id int) (model.Users, error) {
	var user model.Users
	err := u.db.QueryRow(common.GetUserById, id).Scan(
		&user.UserID,
		&user.Name,
		&user.Role,
		&user.Email,
		&user.Password,
		&user.IsDeleted,
	)
	if err != nil {
		return model.Users{}, err
	}
	return user, nil
}

func (u *userRepository) Update(payload model.Users, id int) (model.Users, error) {
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
		payload.Name,
		payload.Role,
		payload.Email,
		payload.Password,
		true,
		id).Scan(
		&user.UserID,
		&user.Name,
		&user.Role,
		&user.Email,
		&user.Password,
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

func (u *userRepository) Delete(id int) (model.Users, error) {
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
		false,
		id).Scan(
		&user.UserID,
		&user.Name,
		&user.Role,
		&user.Email,
		&user.Password,
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

func NewAdminTrainerRepository(db *sql.DB) UserRepositpry {
	return &userRepository{db: db}
}
