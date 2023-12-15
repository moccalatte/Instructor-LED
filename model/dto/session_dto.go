package dto

type SessionRequestDto struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	SessionDate string `json:"session_date"`
	SessionTime string `json:"session_time"`
	SessionLink string `json:"session_link"`
	TrainerID   string `json:"trainer_id"`
	IsDeleted   bool   `json:"is_deleted"`
	Note        string `json:"note"`
}
