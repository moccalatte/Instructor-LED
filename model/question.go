package model

import (
	"time"
)

type Question struct {
	QuestionID  string    `json:"question_id"`
	SessionID   Session   `json:"session_id"`
	StudentID   Student   `json:"student_id"`
	TrainerID   Users     `json:"trainer_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CourseID    Course    `json:"course_id"`
	Image       string    `json:"image"`
	Answer      string    `json:"answer"`
	Status      string    `json:"status"`
	IsDeleted   bool      `json:"is_deleted"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}