package dto

type AttendanceRequestDto struct {
	SessionID         string `json:"session_id"`
	StudentID         string `json:"student_id"`
	AttendanceStudent bool   `json:"attendance_student"`
	IsDeleted         bool   `json:"is_deleted"`
}
