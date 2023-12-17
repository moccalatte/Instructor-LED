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

// func (suite *StudentRepositoryTestSuite) TestCreateSucces() {
// 	dummyPayload := model.Student{
// 		StudentID:   "12345678910",
// 		Fullname:    "ARIA",
// 		BirthDate:   "1991-12-12",
// 		BirthPlace:  "Mesir",
// 		Address:     "Mexico",
// 		Education:   "S1 teknik tambang",
// 		Institution: "ITB",
// 		Job:         "Cto",
// 		Email:       "ARIA@gmail.com",
// 		Password:    "irfanandikarizkiAndre",
// 		CreatedAt:   time.Now(),
// 		UpdatedAt:   time.Now(),
// 		IsDeleted:   false,
// 		Role:        "Boss",
// 	}

// 	suite.sqlmock.ExpectBegin()
// 	rows:=sqlmock.NewRows([]string{"student_id", "fullname", "birth_date", "birth_place", "address", "education", "institution", "job", "email", "password", "created_at", "updated_at", "is_deleted", "role"}).AddRow(
// 		dummyPayload.StudentID,
// 		dummyPayload.Fullname,
// 		dummyPayload.BirthDate,
// 		dummyPayload.BirthPlace,
// 		dummyPayload.Address,
// 		dummyPayload.Education,
// 		dummyPayload.Institution,
// 		dummyPayload.Job,
// 		dummyPayload.Email,
// 		dummyPayload.Password,
// 		dummyPayload.CreatedAt,
// 		dummyPayload.UpdatedAt,
// 		dummyPayload.IsDeleted,
// 		dummyPayload.Role,
// 	)

// 	suite.sqlmock.ExpectQuery("insert into student").WillReturnRows(rows)
// 	suite.sqlmock.ExpectCommit()
// 	actual, err := suite.repo.Create(dummyPayload)

// 	assert.Nil(suite.T(), err, "Error should be nill")
// 	assert.Equal(suite.T(),dummyPayload.StudentID,actual.StudentID,"StudentId should match")
// 	assert.Equal(suite.T(),dummyPayload.Fullname,actual.Fullname,"Fullname should match")
// 	assert.Equal(suite.T(),dummyPayload.BirthDate,actual.BirthDate,"BirthDate should match")
// 	assert.Equal(suite.T(),dummyPayload.BirthPlace,actual.BirthPlace,"BirthPlace should match")
// 	assert.Equal(suite.T(),dummyPayload.Address,actual.Address,"Address should match")
// 	assert.Equal(suite.T(),dummyPayload.Education,actual.Education,"Education should match")
// 	assert.Equal(suite.T(),dummyPayload.Institution,actual.Institution,"Institution should match")
// 	assert.Equal(suite.T(),dummyPayload.Job,actual.Job,"Job should match")
// 	assert.Equal(suite.T(),dummyPayload.Email,actual.Email,"Email should match")
// 	assert.Equal(suite.T(),dummyPayload.Password,actual.Password,"Password should match")
// 	assert.Equal(suite.T(),dummyPayload.CreatedAt,actual.CreatedAt,"CreatedAt should match")
// 	assert.Equal(suite.T(),dummyPayload.UpdatedAt,actual.UpdatedAt,"UpdatedAt should match")
// 	assert.Equal(suite.T(),dummyPayload.IsDeleted,actual.IsDeleted,"IsDeleted should match")
// 	assert.Equal(suite.T(),dummyPayload.Role,actual.Role,"IsDeleted should match")

// }

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

func (suite *StudentRepositoryTestSuite) TestCreateStudent_Success() {
	dummy := model.Student{
		StudentID:   "sdsdsdsdsdsdsd",
		Fullname:    "Joko Santoso",
		BirthDate:   "2000-01-01",
		BirthPlace:  "Jakarta",
		Address:     "Jl. ABC No. 123",
		Education:   "Bachelor",
		Institution: "University XYZ",
		Job:         "Student",
		Email:       "chril@example.com",
		Password:    "230104",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now().Add(3 * 24 * time.Hour),
		Role:        "student",
		IsDeleted:   false,
	}

	suite.sqlmock.ExpectBegin()

	rows := sqlmock.NewRows([]string{"student_id", "fullname", "birth_date", "birth_place", "address", "education", "institution", "job", "email", "password", "created_at", "updated_at", "role", "is_deleted"}).
		AddRow(dummy.StudentID, dummy.Fullname, dummy.BirthDate, dummy.BirthPlace, dummy.Address, dummy.Education, dummy.Institution, dummy.Job, dummy.Email, dummy.Password, dummy.CreatedAt, dummy.UpdatedAt, dummy.Role, dummy.IsDeleted)
	suite.sqlmock.ExpectQuery("insert into student").WithArgs(dummy.Fullname, dummy.BirthDate, dummy.BirthPlace, dummy.Address, dummy.Education, dummy.Institution, dummy.Job, dummy.Email, dummy.Password, sqlmock.AnyArg(), sqlmock.AnyArg(), dummy.Role, dummy.IsDeleted).
		WillReturnRows(rows)

	suite.sqlmock.ExpectCommit()

	actual, err := suite.repo.Create(dummy)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), dummy.StudentID, actual.StudentID)
}

func (suite *StudentRepositoryTestSuite) TestGetAllStudents_Success() {
	dummyStudents := []model.Student{
		{
			StudentID:   "1",
			Fullname:    "John Doe",
			BirthDate:   "2000-01-01",
			BirthPlace:  "City A",
			Address:     "Street 123",
			Education:   "Bachelor",
			Institution: "University X",
			Job:         "Student",
			Email:       "john.doe@example.com",
			Password:    "hashed_password_1",
			IsDeleted:   false,
		},
		{
			StudentID:   "2",
			Fullname:    "Jane Doe",
			BirthDate:   "1998-05-15",
			BirthPlace:  "City B",
			Address:     "Street 456",
			Education:   "Master",
			Institution: "University Y",
			Job:         "Student",
			Email:       "jane.doe@example.com",
			Password:    "hashed_password_2",
			IsDeleted:   false,
		},
	}

	rows := sqlmock.NewRows([]string{
		"student_id", "fullname", "birth_date", "birth_place", "address",
		"education", "institution", "job", "email", "password", "is_deleted",
	}).AddRow(
		dummyStudents[0].StudentID, dummyStudents[0].Fullname, dummyStudents[0].BirthDate,
		dummyStudents[0].BirthPlace, dummyStudents[0].Address, dummyStudents[0].Education,
		dummyStudents[0].Institution, dummyStudents[0].Job, dummyStudents[0].Email,
		dummyStudents[0].Password, dummyStudents[0].IsDeleted,
	).AddRow(
		dummyStudents[1].StudentID, dummyStudents[1].Fullname, dummyStudents[1].BirthDate,
		dummyStudents[1].BirthPlace, dummyStudents[1].Address, dummyStudents[1].Education,
		dummyStudents[1].Institution, dummyStudents[1].Job, dummyStudents[1].Email,
		dummyStudents[1].Password, dummyStudents[1].IsDeleted,
	)

	suite.sqlmock.ExpectQuery(`select \* from student where is_deleted \= \$1;`).WillReturnRows(rows)

	actualStudents, err := suite.repo.GetAll()

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), dummyStudents, actualStudents)
}

func (suite *StudentRepositoryTestSuite) TestUpdateStudent_Success() {
	dummy := model.Student{
		StudentID:   "sdsdsdsdsdsdsd",
		Fullname:    "Joko Santoso",
		BirthDate:   "2000-01-01",
		BirthPlace:  "Jakarta",
		Address:     "Jl. ABC No. 123",
		Education:   "Bachelor",
		Institution: "University XYZ",
		Job:         "Student",
		Email:       "chril@example.com",
		Password:    "230104",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now().Add(3 * 24 * time.Hour),
		Role:        "student",
		IsDeleted:   false,
	}

	suite.sqlmock.ExpectBegin()

	rows := sqlmock.NewRows([]string{
		"student_id", "fullname", "birth_date", "birth_place", "address",
		"education", "institution", "job", "email", "password", "created_at",
		"updated_at", "is_deleted", "role",
	}).AddRow(
		dummy.StudentID, dummy.Fullname, dummy.BirthDate, dummy.BirthPlace, dummy.Address,
		dummy.Education, dummy.Institution, dummy.Job, dummy.Email, dummy.Password, dummy.CreatedAt,
		dummy.UpdatedAt, dummy.IsDeleted, dummy.Role,
	)

	suite.sqlmock.ExpectQuery(`update student set fullname \= \$1, birth_date \= \$2, birth_place \= \$3, address \= \$4, education \= \$5, institution \= \$6, job \= \$7, email \= \$8, password \= \$9, updated_at \= \$10, is_deleted \= \$11, role \= \$12 where student_id \= \$13 returning student_id, fullname, birth_date, birth_place, address, education, institution, job, email, password, created_at, updated_at, is_deleted, role`).WithArgs(
		dummy.Fullname, dummy.BirthDate, dummy.BirthPlace, dummy.Address, dummy.Education,
		dummy.Institution, dummy.Job, dummy.Email, dummy.Password, dummy.CreatedAt, dummy.UpdatedAt,
		dummy.Role, dummy.StudentID,
	).WillReturnRows(rows)

	suite.sqlmock.ExpectCommit()

	actual, err := suite.repo.Update(dummy, dummy.StudentID)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), dummy.StudentID, actual.StudentID)
	assert.Equal(suite.T(), dummy.Fullname, actual.Fullname)
	assert.Equal(suite.T(), dummy.BirthDate, actual.BirthDate)
	assert.Equal(suite.T(), dummy.BirthPlace, actual.BirthPlace)
	assert.Equal(suite.T(), dummy.Address, actual.Address)
	assert.Equal(suite.T(), dummy.Education, actual.Education)
	assert.Equal(suite.T(), dummy.Institution, actual.Institution)
	assert.Equal(suite.T(), dummy.Job, actual.Job)
	assert.Equal(suite.T(), dummy.Email, actual.Email)
	assert.Equal(suite.T(), dummy.Password, actual.Password)
	assert.WithinDuration(suite.T(), dummy.CreatedAt, actual.CreatedAt, time.Second)
	assert.WithinDuration(suite.T(), dummy.UpdatedAt, actual.UpdatedAt, time.Second)
	assert.Equal(suite.T(), dummy.IsDeleted, actual.IsDeleted)
	assert.Equal(suite.T(), dummy.Role, actual.Role)
}

func (suite *StudentRepositoryTestSuite) TestDelete() {

}

func (suite *StudentRepositoryTestSuite) TestGetStudentbyEmail() {

}
