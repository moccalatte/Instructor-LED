package model

import (
    "github.com/google/uuid"
)

type Course struct {
    CourseID      uuid.UUID `json:"course_id"`
    CourseName    string    `json:"course_name"`
    CourseDetailID uuid.UUID `json:"course_detail_id"`
    IsDeleted     bool      `json:"is_deleted"`
}

type CourseDetail struct {
    CourseDetailID uuid.UUID `json:"course_detail_id"`
    Chapter       string    `json:"chapter"`
    CourseID      uuid.UUID `json:"course_id"`
    IsDeleted     bool      `json:"is_deleted"`
}
