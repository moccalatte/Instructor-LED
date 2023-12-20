package model

import (
	"time"
)

type Session struct {
	SessionID   string    `json:"session_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	SessionDate string    `json:"session_date"`
	SessionTime string    `json:"session_time"`
	SessionLink string    `json:"session_link"`
	TrainerID   string    `json:"trainer_id"`
	IsDeleted   bool      `json:"is_deleted"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Note        string    `json:"note"`
}
