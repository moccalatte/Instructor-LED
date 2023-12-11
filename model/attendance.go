package model

import "time"

type Attendance struct {
	AttendanceID      string    `json:"attendance_id"`
	SessionID         string    `json:"session_id"`
	StudentID         string    `json:"student_id"`
	AttendanceStudent bool      `json:"attendance_student"`
	IsDeleted         bool      `json:"is_deleted"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}
