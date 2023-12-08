package manager

import "final-project-kelompok-1/usecase"

type UseCaseManager interface {
	StudentUseCase() usecase.StudentUseCase
	UserUseCase() usecase.UserUseCase
	CourseCase() usecase.CourseUseCase
	SessionUseCase() usecase.SessionUseCase
}

type useCaseManager struct {
	repo RepoManager
}

func (u *useCaseManager) StudentUseCase() usecase.StudentUseCase {
	return usecase.NewStudentUseCase(u.repo.StudentRepo())
}

func (u *useCaseManager) UserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(u.repo.UserRepo())
}

func (u *useCaseManager) CourseCase() usecase.CourseUseCase {
	return usecase.NewCourseUseCase(u.repo.CourseRepo())
}

func (u *useCaseManager) SessionUseCase() usecase.SessionUseCase {
	return usecase.NewSessionUseCase(u.repo.SessionRepo(), u.UserUseCase())
}

func NewUseCaseManager(repo RepoManager) UseCaseManager {
	return &useCaseManager{repo: repo}
}
