package dto

import "final-project-kelompok-1/model"

type CourseRequestDto struct {
	CourseName   string               `json:"course_name"`
	Description  string               `json:"description"`
	IsDeleted    bool                 `json:"is_deleted"`
	CourseDetail []model.CourseDetail `json:"course_detail"`
}

type CourseDetailRequestDto struct {
	CourseId      string `json:"course_id"`
	CourseChapter string `json:"course_chapter"`
	CourseContent string `json:"course_content"`
}
