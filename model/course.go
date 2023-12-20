package model

import (
	"time"
)

type Course struct {
	CourseID    string    `json:"course_id"`
	CourseName  string    `json:"course_name"`
	Description string    `json:"description"`
	IsDeleted   bool      `json:"is_deleted"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type CourseDetail struct {
	CourseDetailID string    `json:"course_detail_id"`
	CourseID       string    `json:"course_id"`
	CourseChapter  string    `json:"course_chapter"`
	CourseContent  string    `json:"course_content"`
	IsDeleted      bool      `json:"is_deleted"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
