package model

import (
    "github.com/google/uuid"
)

type Attendance struct {
    AttendanceID    uuid.UUID `json:"attendance_id"`
    StudentID       uuid.UUID `json:"student_id"`
    CourseID        uuid.UUID `json:"course_id"`
    TrainerAdminID  uuid.UUID `json:"trainer_admin_id"`
    Status          bool      `json:"status"`
    IsDeleted       bool      `json:"is_deleted"`
}

