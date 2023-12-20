package usecase

import (
	"final-project-kelompok-1/repository"
	"final-project-kelompok-1/utils/common"
	"fmt"
	"strconv"
)

type CsvUseCase interface {
	WriteCsv() ([]string, error)
}

type csvUseCase struct {
	se   SessionUseCase
	at   AttendanceUseCase
	sn   StudentUseCase
	tn   UserUseCase
	qu   QuestionUseCase
	cn   CourseUseCase
	cs   common.CvsCommon
	repo repository.CsvRepository
}

func (c *csvUseCase) WriteCsv() ([]string, error) {
	startPoint, err := c.repo.CsvStart()
	if err != nil {
		return nil, fmt.Errorf("failed in usecase: %s", err.Error())
	}

	// Membuat file CSV
	c.cs.CreateFile()

	for _, data := range startPoint {
		dataSession, err := c.se.FindSessionById(data)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		trainerName, err := c.tn.FindUserByID(dataSession.TrainerID)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		attandance, err := c.at.FindAttendanceBySessionId(dataSession.SessionID)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		attendanceBool := strconv.FormatBool(attandance.AttendanceStudent)

		studentName, err := c.sn.FindStudentByID(attandance.StudentID)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		question, err := c.qu.FindQuestionByStudentId(studentName.StudentID)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		course, err := c.cn.FindCourseByID(question.CourseID)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		row := []string{
			dataSession.Title, dataSession.SessionDate, dataSession.SessionTime,
			trainerName.Fullname, studentName.Fullname, attendanceBool, question.Title,
			question.Status, course.CourseName,
		}
		cvString := common.ConvertArrayToString(row)

		err = c.cs.WriterData(cvString)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
	}

	fmt.Println("Data is successfully written to CSV")
	return startPoint, nil
}

func NewCsvUsecase(se SessionUseCase, at AttendanceUseCase, sn StudentUseCase, tn UserUseCase, qu QuestionUseCase, cn CourseUseCase, cs common.CvsCommon, repo repository.CsvRepository) CsvUseCase {
	return &csvUseCase{se: se, at: at, sn: sn, tn: tn, qu: qu, cn: cn, cs: cs, repo: repo}
}
