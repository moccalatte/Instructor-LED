package manager

import "final-project-kelompok-1/usecase"

type UseCaseManager interface {
	StudentUseCase() usecase.StudentUseCase
	UserUseCase() usecase.UserUseCase
	CourseCase() usecase.CourseUseCase
	CourseDetailUseCase() usecase.CourseDetailUseCase
	QuestionUseCase() usecase.QuestionUseCase
	SessionCaseUseCase() usecase.SessionUseCase
	AttendanceUseCase() usecase.AttendanceUseCase
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

func (u *useCaseManager) CourseDetailUseCase() usecase.CourseDetailUseCase {
	return usecase.NewCourseDetailUseCase(u.repo.CourseDetailRepo())
}

func (u *useCaseManager) QuestionUseCase() usecase.QuestionUseCase {
	return usecase.NewQuestion(u.repo.Question())
}

func (u *useCaseManager) AttendanceUseCase() usecase.AttendanceUseCase {
	return usecase.NewAttendanceUseCase(u.repo.Attendance())
}

func (u *useCaseManager) SessionCaseUseCase() usecase.SessionUseCase {
	return usecase.NewSession(u.repo.Session())
}

func NewUseCaseManager(repo RepoManager) UseCaseManager {
	return &useCaseManager{repo: repo}
}
