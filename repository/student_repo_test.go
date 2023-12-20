package repository

import (
	"database/sql"
	"final-project-kelompok-1/model"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type StudentRepositoryTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	sqlmock sqlmock.Sqlmock
	repo    StudentRepository
}

func (suite *StudentRepositoryTestSuite) SetupTest() {
	db, sqlmock, err := sqlmock.New()
	assert.NoError(suite.T(), err)
	suite.mockDb = db
	suite.sqlmock = sqlmock
	suite.repo = NewStudentRepository(suite.mockDb)
}

func TestStudentRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(StudentRepositoryTestSuite))
}

func (suite *StudentRepositoryTestSuite) TestGetById() {
	dummy := model.Student{
		StudentID:   "yweurywieyrwe",
		Fullname:    "pakcik",
		BirthDate:   "2003-10-20",
		BirthPlace:  "Medan",
		Address:     "Bengkulu",
		Education:   "S10",
		Institution: "EnigmaCamp",
		Job:         "profesor",
		Email:       "email_eyak@gmail.com",
		Password:    "23827364826342",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		IsDeleted:   false,
		Role:        "siswa",
	}
	query := "SELECT \\* FROM student WHERE student_id = \\$1;"
	studentID := "yweurywieyrwe"

	rows := sqlmock.NewRows([]string{"student_id", "fullname", "birth_date", "birth_place", "address", "education", "institution", "job", "email", "password", "created_at", "updated_at", "is_deleted", "role"}).AddRow(
		dummy.StudentID,
		dummy.Fullname,
		dummy.BirthDate,
		dummy.BirthPlace,
		dummy.Address,
		dummy.Education,
		dummy.Institution,
		dummy.Job,
		dummy.Email,
		dummy.Password,
		dummy.CreatedAt,
		dummy.UpdatedAt,
		dummy.IsDeleted,
		dummy.Role,
	)
	suite.sqlmock.ExpectQuery(query).WithArgs(studentID).WillReturnRows(rows)

	actual, err := suite.repo.GetById(studentID)

	assert.Nil(suite.T(), err, "Error should be nill")
	assert.Equal(suite.T(), dummy.StudentID, actual.StudentID, "StudentId should match")
	assert.Equal(suite.T(), dummy.Fullname, actual.Fullname, "Fullname should match")
	assert.Equal(suite.T(), dummy.BirthDate, actual.BirthDate, "BirthDate should match")
	assert.Equal(suite.T(), dummy.BirthPlace, actual.BirthPlace, "BirthPlace should match")
	assert.Equal(suite.T(), dummy.Address, actual.Address, "Address should match")
	assert.Equal(suite.T(), dummy.Education, actual.Education, "Education should match")
	assert.Equal(suite.T(), dummy.Institution, actual.Institution, "Institution should match")
	assert.Equal(suite.T(), dummy.Job, actual.Job, "Job should match")
	assert.Equal(suite.T(), dummy.Email, actual.Email, "Email should match")
	assert.Equal(suite.T(), dummy.Password, actual.Password, "Password should match")
	assert.Equal(suite.T(), dummy.CreatedAt, actual.CreatedAt, "CreatedAt should match")
	assert.Equal(suite.T(), dummy.UpdatedAt, actual.UpdatedAt, "UpdatedAt should match")
	assert.Equal(suite.T(), dummy.IsDeleted, actual.IsDeleted, "IsDeleted should match")
	assert.Equal(suite.T(), dummy.Role, actual.Role, "IsDeleted should match")
}

func (suite *StudentRepositoryTestSuite) TestDelete() {
	dummy := model.Student{
		StudentID:   "yweurywieyrwe",
		Fullname:    "pakcik",
		BirthDate:   "2003-10-20",
		BirthPlace:  "Medan",
		Address:     "Bengkulu",
		Education:   "S10",
		Institution: "EnigmaCamp",
		Job:         "profesor",
		Role:        "siswa",
		Email:       "email_eyak@gmail.com",
		Password:    "23827364826342",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		IsDeleted:   true,
	}

	query := "update student set is_deleted = \\$1 where student_id = \\$2 returning student_id, fullname, birth_date, birth_place, address, education, institution, job, email, password, created_at, updated_at, is_deleted, role;"

	suite.sqlmock.ExpectBegin()
	rows := sqlmock.NewRows([]string{"student_id", "fullname", "birth_date", "birth_place", "address", "education", "institution", "job", "email", "password", "created_at", "updated_at", "is_deleted"}).AddRow(
		dummy.StudentID,
		dummy.Fullname,
		dummy.BirthDate,
		dummy.BirthPlace,
		dummy.Address,
		dummy.Education,
		dummy.Institution,
		dummy.Job,
		dummy.Email,
		dummy.Password,
		dummy.CreatedAt,
		dummy.UpdatedAt,
		dummy.IsDeleted,
	)
	suite.sqlmock.ExpectQuery(query).WithArgs(
		true,
		dummy.StudentID,
	).WillReturnRows(rows)

	suite.sqlmock.ExpectCommit()

	actual, err := suite.repo.Delete(dummy.StudentID)

	assert.Nil(suite.T(), err, "Error should be nill")
	assert.Equal(suite.T(), dummy.StudentID, actual.StudentID, "StudentId should match")
	assert.Equal(suite.T(), dummy.Fullname, actual.Fullname, "Fullname should match")
	assert.Equal(suite.T(), dummy.BirthDate, actual.BirthDate, "BirthDate should match")
	assert.Equal(suite.T(), dummy.BirthPlace, actual.BirthPlace, "BirthPlace should match")
	assert.Equal(suite.T(), dummy.Address, actual.Address, "Address should match")
	assert.Equal(suite.T(), dummy.Education, actual.Education, "Education should match")
	assert.Equal(suite.T(), dummy.Institution, actual.Institution, "Institution should match")
	assert.Equal(suite.T(), dummy.Job, actual.Job, "Job should match")
	assert.Equal(suite.T(), dummy.Email, actual.Email, "Email should match")
	assert.Equal(suite.T(), dummy.Password, actual.Password, "Password should match")
	assert.Equal(suite.T(), dummy.CreatedAt, actual.CreatedAt, "CreatedAt should match")
	assert.Equal(suite.T(), dummy.UpdatedAt, actual.UpdatedAt, "UpdatedAt should match")
	assert.Equal(suite.T(), dummy.IsDeleted, actual.IsDeleted, "IsDeleted should match")

}

func (suite *StudentRepositoryTestSuite) TestGetStudentbyEmail() {
	dummy := model.Student{
		StudentID:   "yweurywieyrwe",
		Fullname:    "pakcik",
		BirthDate:   "2003-10-20",
		BirthPlace:  "Medan",
		Address:     "Bengkulu",
		Education:   "S10",
		Institution: "EnigmaCamp",
		Job:         "profesor",
		Role:        "siswa",
		Email:       "email_eyak@gmail.com",
		Password:    "23827364826342",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		IsDeleted:   false,
	}
	query := "select student_id, fullname, birth_date, birth_place, address, education, institution, job, role,email, password, created_at, updated_at, is_deleted from student where fullname = \\$1 OR email = \\$1;"
	email := "email_eyak@gmail.com"

	rows := sqlmock.NewRows([]string{"student_id", "fullname", "birth_date", "birth_place", "address", "education", "institution", "job", "role", "email", "password", "created_at", "updated_at", "is_deleted"}).AddRow(
		dummy.StudentID,
		dummy.Fullname,
		dummy.BirthDate,
		dummy.BirthPlace,
		dummy.Address,
		dummy.Education,
		dummy.Institution,
		dummy.Job,
		dummy.Role,
		dummy.Email,
		dummy.Password,
		dummy.CreatedAt,
		dummy.UpdatedAt,
		dummy.IsDeleted,
	)
	suite.sqlmock.ExpectQuery(query).WithArgs(email).WillReturnRows(rows)

	actual, err := suite.repo.GetByStudentEmail(email)

	assert.Nil(suite.T(), err, "Error should be nill")
	assert.Equal(suite.T(), dummy.StudentID, actual.StudentID, "StudentId should match")
	assert.Equal(suite.T(), dummy.Fullname, actual.Fullname, "Fullname should match")
	assert.Equal(suite.T(), dummy.BirthDate, actual.BirthDate, "BirthDate should match")
	assert.Equal(suite.T(), dummy.BirthPlace, actual.BirthPlace, "BirthPlace should match")
	assert.Equal(suite.T(), dummy.Address, actual.Address, "Address should match")
	assert.Equal(suite.T(), dummy.Education, actual.Education, "Education should match")
	assert.Equal(suite.T(), dummy.Institution, actual.Institution, "Institution should match")
	assert.Equal(suite.T(), dummy.Job, actual.Job, "Job should match")
	assert.Equal(suite.T(), dummy.Email, actual.Email, "Email should match")
	assert.Equal(suite.T(), dummy.Password, actual.Password, "Password should match")
	assert.Equal(suite.T(), dummy.CreatedAt, actual.CreatedAt, "CreatedAt should match")
	assert.Equal(suite.T(), dummy.UpdatedAt, actual.UpdatedAt, "UpdatedAt should match")
	assert.Equal(suite.T(), dummy.IsDeleted, actual.IsDeleted, "IsDeleted should match")
	assert.Equal(suite.T(), dummy.Role, actual.Role, "IsDeleted should match")
}

func (suite *StudentRepositoryTestSuite) TestGetall() {
	dummyResult := []model.Student{
		{
			StudentID:   "yweurywieyrwe",
			Fullname:    "pakcik",
			BirthDate:   "2003-10-20",
			BirthPlace:  "Medan",
			Address:     "Bengkulu",
			Education:   "S10",
			Institution: "EnigmaCamp",
			Job:         "profesor",
			Email:       "email_eyak@gmail.com",
			Password:    "23827364826342",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			IsDeleted:   false,
			Role:        "siswa",
		},
		{
			StudentID:   "yweurywieyrwe",
			Fullname:    "pakcik",
			BirthDate:   "2003-10-20",
			BirthPlace:  "Medan",
			Address:     "Bengkulu",
			Education:   "S10",
			Institution: "EnigmaCamp",
			Job:         "profesor",
			Email:       "email_eyak@gmail.com",
			Password:    "23827364826342",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			IsDeleted:   false,
			Role:        "siswa"},
	}

	query := "select \\* from student where is_deleted = \\$1;"
	rows := sqlmock.NewRows([]string{"student_id", "fullname", "birth_date", "birth_place", "address", "education", "institution", "job", "email", "password", "is_deleted"})

	for _, result := range dummyResult {
		rows.AddRow(
			result.StudentID,
			result.Fullname,
			result.BirthDate,
			result.BirthPlace,
			result.Address,
			result.Education,
			result.Institution,
			result.Job,
			result.Email,
			result.Password,
			result.IsDeleted,
		)
	}
	suite.sqlmock.ExpectQuery(query).WithArgs().WillReturnRows(rows)

	actual, err := suite.repo.GetAll()

	assert.Nil(suite.T(), err, "error should be nill")
	assert.Len(suite.T(), actual, len(dummyResult))

	for i, expected := range dummyResult {
		assert.Equal(suite.T(), expected.StudentID, actual[i].StudentID)
		assert.Equal(suite.T(), expected.Fullname, actual[i].Fullname)
		assert.Equal(suite.T(), expected.BirthDate, actual[i].BirthDate)
		assert.Equal(suite.T(), expected.BirthPlace, actual[i].BirthPlace)

	}
}

// func (suite *StudentRepositoryTestSuite) TestCreate() {
// 	dummy := model.Student{
// 		StudentID:   "yweurywieyrwe",
// 		Fullname:    "pakcik",
// 		BirthDate:   "2003-10-20",
// 		BirthPlace:  "Medan",
// 		Address:     "Bengkulu",
// 		Education:   "S10",
// 		Institution: "EnigmaCamp",
// 		Job:         "profesor",
// 		Email:       "email_eyak@gmail.com",
// 		Password:    "23827364826342",
// 		CreatedAt:   time.Now(),
// 		UpdatedAt:   time.Now(),
// 		Role:        "siswa",
// 		IsDeleted:   false,
// 	}
// 	suite.sqlmock.ExpectBegin()
// 	query := "insert into student"

// 	rows := sqlmock.NewRows([]string{"student_id", "fullname", "birth_date", "birth_place", "address", "education", "institution", "job", "email", "password", "created_at", "updated_at", "role", "is_deleted"}).AddRow(
// 		dummy.StudentID,
// 		dummy.Fullname,
// 		dummy.BirthDate,
// 		dummy.BirthPlace,
// 		dummy.Address,
// 		dummy.Education,
// 		dummy.Institution,
// 		dummy.Job,
// 		dummy.Email,
// 		dummy.Password,
// 		dummy.CreatedAt,
// 		dummy.UpdatedAt,
// 		dummy.Role,
// 		dummy.IsDeleted,
// 	)

// 	suite.sqlmock.ExpectQuery(query).WithArgs(
// 		dummy.StudentID,
// 		dummy.Fullname,
// 		dummy.BirthDate,
// 		dummy.BirthPlace,
// 		dummy.Address,
// 		dummy.Education,
// 		dummy.Institution,
// 		dummy.Job,
// 		dummy.Email,
// 		dummy.Password,
// 		dummy.CreatedAt,
// 		dummy.UpdatedAt,
// 		dummy.Role,
// 		dummy.IsDeleted,
// 	).WillReturnRows(rows)
// 	suite.sqlmock.ExpectCommit()

// 	actual, err := suite.repo.Create(dummy)

// 	assert.Nil(suite.T(), err, "Error should be nill")
// 	assert.Equal(suite.T(), dummy.StudentID, actual.StudentID, "StudentId should match")

// }

// func (suite *StudentRepositoryTestSuite) TestCreateStudent_Success() {
// 	dummy := model.Student{
// 		StudentID:   "yweurywieyrwe",
// 		Fullname:    "pakcik",
// 		BirthDate:   "2003-10-20",
// 		BirthPlace:  "Medan",
// 		Address:     "Bengkulu",
// 		Education:   "S10",
// 		Institution: "EnigmaCamp",
// 		Job:         "profesor",
// 		Email:       "email_eyak@gmail.com",
// 		Password:    "23827364826342",
// 		CreatedAt:   time.Now(),
// 		UpdatedAt:   time.Now().Local(),
// 		Role:        "siswa",
// 		IsDeleted:   false,
// 	}

// 	suite.sqlmock.ExpectBegin()

// 	rows := sqlmock.NewRows([]string{"student_id", "fullname", "birth_date", "birth_place", "address", "education", "institution", "job", "email", "password", "created_at", "updated_at", "role", "is_deleted"}).AddRow(
// 		dummy.StudentID,
// 		dummy.Fullname,
// 		dummy.BirthDate,
// 		dummy.BirthPlace,
// 		dummy.Address,
// 		dummy.Education,
// 		dummy.Institution,
// 		dummy.Job,
// 		dummy.Email,
// 		dummy.Password,
// 		dummy.CreatedAt,
// 		dummy.UpdatedAt,
// 		dummy.Role,
// 		dummy.IsDeleted,
// 	)

// 	suite.sqlmock.ExpectQuery("insert into student").WithArgs(
// 		dummy.Fullname,
// 		dummy.BirthDate,
// 		dummy.BirthPlace,
// 		dummy.Address,
// 		dummy.Education,
// 		dummy.Institution,
// 		dummy.Job,
// 		dummy.Email,
// 		dummy.Password,
// 		dummy.UpdatedAt,
// 		dummy.Role,
// 		dummy.IsDeleted,
// 	).WillReturnRows(rows)

// 	suite.sqlmock.ExpectCommit()

// 	actual, err := suite.repo.Create(dummy)
// 	assert.Nil(suite.T(), err)
// 	assert.NoError(suite.T(), err)
// 	assert.Equal(suite.T(), dummy.StudentID, actual.StudentID)
// 	assert.WithinDuration(suite.T(),dummy.UpdatedAt,actual.UpdatedAt,time.Second)
// }
