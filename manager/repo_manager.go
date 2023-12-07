package manager

import "final-project-kelompok-1/repository"

type RepoManager interface {
	StudentRepo() repository.StudentRepository
	AdminTrainerRepo() repository.AdminTrainerRepository
	RoleRepo() repository.RoleRepository
}

type repoManager struct {
	infra InfraManager
}

func (r *repoManager) StudentRepo() repository.StudentRepository {
	return repository.NewStudentRepository(r.infra.Conn())
}

func (r *repoManager) AdminTrainerRepo() repository.AdminTrainerRepository {
	return repository.NewAdminTrainerRepository(r.infra.Conn())
}

func (r *repoManager) RoleRepo() repository.RoleRepository {
	return repository.NewRoleRepository(r.infra.Conn())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{infra: infra}
}
