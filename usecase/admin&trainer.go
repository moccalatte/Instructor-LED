package usecase

import (
	"final-project-kelompok-1/model"
	"final-project-kelompok-1/model/dto"
	"final-project-kelompok-1/repository"
)

type AdminTrainerUseCase interface {
	AddAdminTrainer(payload dto.AdminOrTrainerRequestDto) (model.AdminTrainer, error)
	FindAdminTrainerByID(id int) (model.AdminTrainer, error)
	UpdateAdminTrainer(payload dto.AdminOrTrainerRequestDto, id int) (model.AdminTrainer, error)
	DeleteAdmintrainer(id int) (model.AdminTrainer, error)
}

type adminTrainerUseCase struct {
	repo repository.AdminTrainerRepository
}

func (a *adminTrainerUseCase) AddAdminTrainer(payload dto.AdminOrTrainerRequestDto) (model.AdminTrainer, error) {

}

func (a *adminTrainerUseCase) FindAdminTrainerByID(id int) (model.AdminTrainer, error) {

}

func (a *adminTrainerUseCase) UpdateAdminTrainer(payload dto.AdminOrTrainerRequestDto, id int) (model.AdminTrainer, error) {

}

func (a *adminTrainerUseCase) DeleteAdmintrainer(id int) (model.AdminTrainer, error) {

}

func NewAdminTrainerUseCase(repo repository.AdminTrainerRepository) AdminTrainerUseCase {
	return &adminTrainerUseCase{repo: repo}
}
