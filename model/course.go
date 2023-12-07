package model

import (
    "github.com/google/uuid"
)

type Course struct {
    CourseID      uuid.UUID `db:"course_id"`
    CourseName    string    `db:"course_name"`
    CourseDetailID uuid.UUID `db:"course_detail_id"`
    IsDeleted     bool      `db:"is_deleted"`
}

type CourseDetail struct {
    CourseDetailID uuid.UUID `db:"course_detail_id"`
    Chapter       string    `db:"chapter"`
    CourseID      uuid.UUID `db:"course_id"`
    IsDeleted     bool      `db:"is_deleted"`
}
