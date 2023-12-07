package model

import (
    "github.com/google/uuid"
)

type Attendance struct {
    AttendanceID    uuid.UUID `db:"attendance_id"`
    StudentID       uuid.UUID `db:"student_id"`
    CourseID        uuid.UUID `db:"course_id"`
    TrainerAdminID  uuid.UUID `db:"trainer_admin_id"`
    Status          bool      `db:"status"`
    IsDeleted       bool      `db:"is_deleted"`
}

