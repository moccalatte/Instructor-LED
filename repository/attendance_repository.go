package repository

import (
	"database/sql"
	"time"

	"final-project-kelompok-1/model"
	"final-project-kelompok-1/utils/common"
)

type AttendanceRepository interface {
	Create(payload model.Attendance) (model.Attendance, error)
	GetById(id string) (model.Attendance, error)
	Update(payload model.Attendance, id string) (model.Attendance, error)
	Delete(id string) (model.Attendance, error)
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
		&attendance.UpdatedAt,
		&attendance.IsDeleted,
	)
	if err != nil {
		return model.Attendance{}, tx.Rollback()
	}

	if err := tx.Commit(); err != nil {
		return model.Attendance{}, err
	}

	return attendance, nil
}

func (a *attendanceRepository) GetById(id string) (model.Attendance, error) {
	var attendance model.Attendance
	err := a.db.QueryRow(common.GetStudentByid, id).Scan(
		&attendance.AttendanceID,
		&attendance.SessionID,
		&attendance.StudentID,
		&attendance.AttendanceStudent,
		&attendance.UpdatedAt,
		&attendance.IsDeleted,
	)
	if err != nil {
		return model.Attendance{}, err
	}
	return attendance, nil

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
		&attendance.UpdatedAt,
		&attendance.IsDeleted,
	)
	if err != nil {
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
		&attendance.UpdatedAt,
		&attendance.IsDeleted,
	)
	if err != nil {
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