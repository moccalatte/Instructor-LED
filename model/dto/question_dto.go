package dto

type QuestionRequestDto struct {
	SessionID   string `json:"session_id"`
	StudentID   string `json:"student_id"`
	TrainerID   string `json:"trainer_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CourseID    string `json:"course_id"`
	ImageURL       string `json:"image_url"`
	Answer      string `json:"answer"`
	Status      string `json:"status"`
	IsDeleted   bool   `json:"is_deleted"`
	ImagePath string `json:"image_path"`
}
