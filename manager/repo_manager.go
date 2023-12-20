package manager

import "final-project-kelompok-1/repository"

type RepoManager interface {
	StudentRepo() repository.StudentRepository
	UserRepo() repository.UserRepository
	CourseRepo() repository.CourseRepository
	CourseDetailRepo() repository.CourseDetailRepository
	Question() repository.QuestionRepository
	Session() repository.SessionRepository
	Attendance() repository.AttendanceRepository
	CsvRepo() repository.CsvRepository
}

type repoManager struct {
	infra InfraManager
}

func (r *repoManager) StudentRepo() repository.StudentRepository {
	return repository.NewStudentRepository(r.infra.Conn())
}

func (r *repoManager) UserRepo() repository.UserRepository {
	return repository.NewUserRepository(r.infra.Conn())
}

func (r *repoManager) CourseRepo() repository.CourseRepository {
	return repository.NewCourseRepository(r.infra.Conn())
}

func (r *repoManager) CourseDetailRepo() repository.CourseDetailRepository {
	return repository.NewCourseDetailRepository(r.infra.Conn())
}

func (r *repoManager) Question() repository.QuestionRepository {
	return repository.NewQuestionRepository(r.infra.Conn())
}

func (r *repoManager) Session() repository.SessionRepository {
	return repository.NewSessionRepository(r.infra.Conn())
}

func (r *repoManager) Attendance() repository.AttendanceRepository {
	return repository.NewAttendanceRepository(r.infra.Conn())
}

func (r *repoManager) CsvRepo() repository.CsvRepository {
	return repository.NewCsv(r.infra.Conn())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{infra: infra}
}
