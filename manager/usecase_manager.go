package manager

import "final-project-kelompok-1/usecase"

type UseCaseManager interface {
	StudentUseCase() usecase.StudentUseCase
	AdminTrainerUseCase() usecase.AdminTrainerUseCase
	RoleUseCase() usecase.RoleUseCase
}

type useCaseManager struct {
	repo RepoManager
}

func (u *useCaseManager) StudentUseCase() usecase.StudentUseCase {
	return usecase.NewStudentUseCase(u.repo.StudentRepo())
}

func (u *useCaseManager) AdminTrainerUseCase() usecase.AdminTrainerUseCase {
	return usecase.NewAdminTrainerUseCase(u.repo.AdminTrainerRepo())
}

func (u *useCaseManager) RoleUseCase() usecase.RoleUseCase {
	return usecase.NewRoleUseCase(u.repo.RoleRepo())
}

func NewUseCaseManager(repo RepoManager) UseCaseManager {
	return &useCaseManager{repo: repo}
}
