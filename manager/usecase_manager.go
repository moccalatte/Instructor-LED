package manager

import (
	"final-project-kelompok-1/usecase"
	"final-project-kelompok-1/utils/common"
)

type UseCaseManager interface {
	StudentUseCase() usecase.StudentUseCase
	UserUseCase() usecase.UserUseCase
	CourseCase() usecase.CourseUseCase
	CourseDetailUseCase() usecase.CourseDetailUseCase
	QuestionUseCase() usecase.QuestionUseCase
	SessionCaseUseCase() usecase.SessionUseCase
	AttendanceUseCase() usecase.AttendanceUseCase
	CsvCaseUseCase(csvService common.CvsCommon) usecase.CsvUseCase
}

type useCaseManager struct {
	repo       RepoManager
	csvService common.CvsCommon
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

func (u *useCaseManager) CsvCaseUseCase(csvService common.CvsCommon) usecase.CsvUseCase {
	return usecase.NewCsvUsecase(
		u.SessionCaseUseCase(),
		u.AttendanceUseCase(),
		u.StudentUseCase(),
		u.UserUseCase(),
		u.QuestionUseCase(),
		u.CourseCase(),
		csvService,
		u.repo.CsvRepo(),
	)
}

func NewUseCaseManager(repo RepoManager, csvService common.CvsCommon) UseCaseManager {
	return &useCaseManager{repo: repo, csvService: csvService}

}
