package repository

import (
	"database/sql"
	"fmt"
	"time"

	"final-project-kelompok-1/model"
	"final-project-kelompok-1/utils/common"
)

type AttendanceRepository interface {
	Create(payload model.Attendance) (model.Attendance, error)
	GetById(id string) (model.Attendance, error)
	GetBySessionId(id string) (model.Attendance, error)
	Update(payload model.Attendance, id string) (model.Attendance, error)
	Delete(id string) (model.Attendance, error)
	GetAll() ([]model.Attendance, error)
}

type attendanceRepository struct {
	db *sql.DB
}

func (a *attendanceRepository) Create(payload model.Attendance) (model.Attendance, error) {
	tx, err := a.db.Begin()
	if err != nil {
		return model.Attendance{}, err
	}

	var attendance model.Attendance
	err = tx.QueryRow(common.CreateAttendance,
		payload.SessionID,
		payload.StudentID,
		payload.AttendanceStudent,
		time.Now(),
		false).Scan(
		&attendance.AttendanceID,
		&attendance.SessionID,
		&attendance.StudentID,
		&attendance.AttendanceStudent,
		&attendance.CreatedAt,
		&attendance.UpdatedAt,
		&attendance.IsDeleted,
	)
	fmt.Print(err, "ATTENDANCE REPO")
	if err != nil {
		fmt.Println("Error attendance in repo : ", err.Error())
		return model.Attendance{}, tx.Rollback()
	}

	if err := tx.Commit(); err != nil {
		return model.Attendance{}, err
	}

	return attendance, nil
}

func (a *attendanceRepository) GetById(id string) (model.Attendance, error) {
	fmt.Print(id, "ID NYA MASUK")
	var attendance model.Attendance
	err := a.db.QueryRow(common.GetAttendanceById, id).Scan(
		&attendance.AttendanceID,
		&attendance.SessionID,
		&attendance.StudentID,
		&attendance.AttendanceStudent,
		&attendance.CreatedAt,
		&attendance.UpdatedAt,
		&attendance.IsDeleted,
	)
	if err != nil {
		fmt.Println("Error attendance in repo : ", err.Error())
		return model.Attendance{}, err
	}
	return attendance, nil

}

func (a *attendanceRepository) GetBySessionId(id string) (model.Attendance, error) {
	var attendance model.Attendance
	err := a.db.QueryRow(common.GetAttandanceBySessionId, id).Scan(
		&attendance.AttendanceID,
		&attendance.SessionID,
		&attendance.StudentID,
		&attendance.AttendanceStudent,
		&attendance.CreatedAt,
		&attendance.UpdatedAt,
		&attendance.IsDeleted,
	)
	if err != nil {
		fmt.Println("Error in repo attendance : ", err.Error())
		return model.Attendance{}, err
	}
	return attendance, nil

}

func (a *attendanceRepository) GetAll() ([]model.Attendance, error) {
	var attendances []model.Attendance

	rows, err := a.db.Query(common.GetAllAttendance)

	if err != nil {
		return attendances, err
	}
	for rows.Next() {
		var attendance model.Attendance
		err := rows.Scan(
			&attendance.AttendanceID,
			&attendance.SessionID,
			&attendance.StudentID,
			&attendance.AttendanceStudent,
			&attendance.CreatedAt,
			&attendance.UpdatedAt,
			&attendance.IsDeleted,
		)

		if err != nil {
			fmt.Println("error in repo :", err.Error())
			return attendances, nil
		}

		attendances = append(attendances, attendance)
	}

	return attendances, nil

}

func (a *attendanceRepository) Update(payload model.Attendance, id string) (model.Attendance, error) {
	tx, err := a.db.Begin()
	if err != nil {
		return model.Attendance{}, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	var attendance model.Attendance
	err = tx.QueryRow(common.UpdateAttendanceById,
		payload.SessionID,
		payload.StudentID,
		payload.AttendanceStudent,
		time.Now(),
		false,
		id).Scan(
		&attendance.AttendanceID,
		&attendance.SessionID,
		&attendance.StudentID,
		&attendance.AttendanceStudent,
		&attendance.CreatedAt,
		&attendance.UpdatedAt,
		&attendance.IsDeleted,
	)
	fmt.Print(err, "REPO ERROR")

	if err != nil {
		fmt.Println("Error attendance in repo : ", err.Error())
		return model.Attendance{}, tx.Rollback()
	}

	if err := tx.Commit(); err != nil {
		return model.Attendance{}, err
	}
	return attendance, nil
}

func (a *attendanceRepository) Delete(id string) (model.Attendance, error) {
	tx, err := a.db.Begin()
	if err != nil {
		return model.Attendance{}, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	var attendance model.Attendance
	err = tx.QueryRow(common.DeleteAttendanceById,
		true,
		id).Scan(
		&attendance.AttendanceID,
		&attendance.SessionID,
		&attendance.StudentID,
		&attendance.AttendanceStudent,
		&attendance.CreatedAt,
		&attendance.UpdatedAt,
		&attendance.IsDeleted,
	)
	fmt.Print(err, "ERROR DI REPOSITORY")
	if err != nil {
		fmt.Println("Error attendance in repo : ", err.Error())
		return model.Attendance{}, tx.Rollback()
	}

	if err := tx.Commit(); err != nil {
		return model.Attendance{}, err
	}
	return attendance, nil
}

func NewAttendanceRepository(db *sql.DB) AttendanceRepository {
	return &attendanceRepository{db: db}
}
