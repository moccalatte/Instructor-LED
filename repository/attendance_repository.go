package repository

import (
	"database/sql"
	"final-project-kelompok-1/model"
	"final-project-kelompok-1/utils/common"
	"fmt"
)

type AttendanceRepository interface {
	GetById(id string) (model.Attendance, error)
	CreateAttedance(payload model.Attendance) (model.Attendance, error)
	UpdateAttendance(payload model.Attendance) (model.Attendance, error)
	DeleteSession(payload model.Attendance) (model.Attendance, error)
}

type attendanceRepository struct {
	db *sql.DB
}

func (a *attendanceRepository) GetById(id string) (model.Attendance, error) {
	var attendance model.Attendance

	err := a.db.QueryRow(common.GetAttendanceById, id).Scan(
		&attendance.AttendanceID,
		&attendance.StudentID,
		&attendance.CourseID,
		&attendance.TrainerAdminID,
	)

	if err != nil {
		return model.Attendance{}, err
	}

	return attendance, nil
}

func (a *attendanceRepository) CreateAttedance(payload model.Attendance) (model.Attendance, error) {
	tx, err := a.db.Begin()
	if err != nil {
		return model.Attendance{}, err
	}

	var attendance model.Attendance
	err = tx.QueryRow(common.CreateAttendance,
		payload.StudentID,
		payload.CourseID,
		payload.TrainerAdminID,
		payload.Status,
	).Scan(
		&attendance.AttendanceID,
		&attendance.StudentID,
		&attendance.CourseID,
		&attendance.TrainerAdminID,
		&attendance.Status,
		&attendance.IsDeleted,
	)

	if err != nil {
		fmt.Println("Error Insert Attendance : ", err)
		return model.Attendance{}, tx.Rollback()
	}

	if err := tx.Commit(); err != nil {
		return model.Attendance{}, err
	}

	return attendance, nil
}

func (a *attendanceRepository) UpdateAttendance(payload model.Attendance) (model.Attendance, error) {
	tx, err := a.db.Begin()
	if err != nil {
		return model.Attendance{}, err
	}

	var attendance model.Attendance
	err = tx.QueryRow(common.UpdateAttendanceById,
		payload.StudentID,
		payload.CourseID,
		payload.TrainerAdminID,
		payload.Status).Scan(
		&attendance.AttendanceID,
		&attendance.StudentID,
		&attendance.CourseID,
		&attendance.TrainerAdminID,
		&attendance.Status,
	)

	if err != nil {
		fmt.Println("Error Update Attendance : ", err)
		return model.Attendance{}, tx.Rollback()
	}

	if err := tx.Commit(); err != nil {
		return model.Attendance{}, err
	}

	return attendance, nil
}

func (a *attendanceRepository) DeleteSession(payload model.Attendance) (model.Attendance, error) {
	tx, err := a.db.Begin()

	if err != nil {
		return model.Attendance{}, err
	}

	var attendance model.Attendance
	err = tx.QueryRow(common.DeleteAttendanceById,
		1,
		payload.AttendanceID).Scan(
		&attendance.IsDeleted,
		&attendance.AttendanceID,
	)

	if err != nil {
		fmt.Println("Error Delete Attendance : ", err)
		return model.Attendance{}, err
	}

	return attendance, nil
}

func NewAttendanceRepository(db *sql.DB) AttendanceRepository {
	return &attendanceRepository{db: db}
}
