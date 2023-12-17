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

func(suite *StudentRepositoryTestSuite)TestGetById(){
	dummy := model.Student{
		StudentID: "yweurywieyrwe",
		Fullname: "pakcik",
		BirthDate: "2003-10-20",
		BirthPlace: "Medan",
		Address: "Bengkulu",
		Education: "S10",
		Institution: "EnigmaCamp",
		Job: "profesor",
		Email: "email_eyak@gmail.com",
		Password: "23827364826342",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsDeleted: false,
		Role: "siswa",
	}
	query:= "SELECT \\* FROM student WHERE student_id = \\$1;"
	studentID := "yweurywieyrwe"

	rows:=sqlmock.NewRows([]string{"student_id", "fullname", "birth_date", "birth_place", "address", "education", "institution", "job", "email", "password", "created_at", "updated_at", "is_deleted", "role"}).AddRow(
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
	assert.Equal(suite.T(),dummy.StudentID,actual.StudentID,"StudentId should match")
	assert.Equal(suite.T(),dummy.Fullname,actual.Fullname,"Fullname should match")
	assert.Equal(suite.T(),dummy.BirthDate,actual.BirthDate,"BirthDate should match")
	assert.Equal(suite.T(),dummy.BirthPlace,actual.BirthPlace,"BirthPlace should match")
	assert.Equal(suite.T(),dummy.Address,actual.Address,"Address should match")
	assert.Equal(suite.T(),dummy.Education,actual.Education,"Education should match")
	assert.Equal(suite.T(),dummy.Institution,actual.Institution,"Institution should match")
	assert.Equal(suite.T(),dummy.Job,actual.Job,"Job should match")
	assert.Equal(suite.T(),dummy.Email,actual.Email,"Email should match")
	assert.Equal(suite.T(),dummy.Password,actual.Password,"Password should match")
	assert.Equal(suite.T(),dummy.CreatedAt,actual.CreatedAt,"CreatedAt should match")
	assert.Equal(suite.T(),dummy.UpdatedAt,actual.UpdatedAt,"UpdatedAt should match")
	assert.Equal(suite.T(),dummy.IsDeleted,actual.IsDeleted,"IsDeleted should match")
	assert.Equal(suite.T(),dummy.Role,actual.Role,"IsDeleted should match")
}

func(suite *StudentRepositoryTestSuite)TestCreate(){

}

func(suite *StudentRepositoryTestSuite)TestGetAll(){

}

func(suite *StudentRepositoryTestSuite)TestUpdate(){

}

func(suite *StudentRepositoryTestSuite)TestDelete(){

}

func(suite *StudentRepositoryTestSuite)TestGetStudentbyEmail(){
	
}