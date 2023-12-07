package repository

import (
	"database/sql"
	"final-project-kelompok-1/model"
)

type AdminTrainerRepository interface {
	Create(payload model.AdminTrainer) (model.AdminTrainer, error)
	GetById(id int) (model.AdminTrainer, error)
	Update(id int) (model.AdminTrainer, error)
	Delete(id int) (model.AdminTrainer, error)
}

type adminTrainerRepository struct {
	db *sql.DB
}

func (a *adminTrainerRepository) Create(payload model.AdminTrainer) (model.AdminTrainer, error) {
}

func (a *adminTrainerRepository) GetById(id int) (model.AdminTrainer, error) {
}

func (a *adminTrainerRepository) Update(id int) (model.AdminTrainer, error) {
}

func (a *adminTrainerRepository) Delete(id int) (model.AdminTrainer, error) {
}

func NewAdminTrainerRepository(db *sql.DB) AdminTrainerRepository {
	return &adminTrainerRepository{db: db}
}
