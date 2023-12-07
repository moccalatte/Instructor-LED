package model

import (
	"time"
    "github.com/google/uuid"
)

type Session struct {
    SessionID       uuid.UUID `json:"session_id"`
    Date            time.Time `json:"date"`
    TrainerAdminID  uuid.UUID `json:"trainer_admin_id"`
    IsDeleted       bool      `json:"is_deleted"`
}
