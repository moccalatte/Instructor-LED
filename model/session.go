package model

import (
	"time"
    "github.com/google/uuid"
)

type Session struct {
    SessionID       uuid.UUID `db:"session_id"`
    Date            time.Time `db:"date"`
    TrainerAdminID  uuid.UUID `db:"trainer_admin_id"`
    IsDeleted       bool      `db:"is_deleted"`
}
