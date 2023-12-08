package manager

import "final-project-kelompok-1/repository"

type RepoManager interface {
	StudentRepo() repository.StudentRepository
	UserRepo() repository.UserRepositpry
	CourseRepo() repository.CourseRepository
	SessionRepo() repository.SessionRepository
}

type repoManager struct {
	infra InfraManager
}

func (r *repoManager) StudentRepo() repository.StudentRepository {
	return repository.NewStudentRepository(r.infra.Conn())
}

func (r *repoManager) UserRepo() repository.UserRepositpry {
	return repository.NewUserRepository(r.infra.Conn())
}

func (r *repoManager) CourseRepo() repository.CourseRepository {
	return repository.NewCourseRepository(r.infra.Conn())
}

func (r *repoManager) SessionRepo() repository.SessionRepository {
	return repository.NewSessionRepository(r.infra.Conn())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{infra: infra}
}
