package usecase

import (
	"final-project-kelompok-1/model"
	"final-project-kelompok-1/model/dto"
	"final-project-kelompok-1/repository"
)

type RoleUseCase interface {
	AddRole(payload dto.RoleRequestDto) (model.Role, error)
	FindRoleByID(id int) (model.Role, error)
	UpdateRole(payload dto.RoleRequestDto, id int) (model.Role, error)
	DeleteRole(id int) (model.Role, error)
}

type roleUseCase struct {
	repo repository.RoleRepository
}

func (r *roleUseCase) AddRole(payload dto.RoleRequestDto) (model.Role, error) {

}

func (r *roleUseCase) FindRoleByID(id int) (model.Role, error) {

}

func (r *roleUseCase) UpdateRole(payload dto.RoleRequestDto, id int) (model.Role, error) {

}

func (r *roleUseCase) DeleteRole(id int) (model.Role, error) {

}

func NewRoleUseCase(repo repository.RoleRepository) RoleUseCase {
	return &roleUseCase{repo: repo}
}
