package dto

import "github.com/google/uuid"

type CourseRequestDto struct {
	CourseName     string    `json:"course_name"`
	CourseDetailID uuid.UUID `json:"course_detail_id"`
}
