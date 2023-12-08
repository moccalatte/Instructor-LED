package model

import (
	"github.com/google/uuid"
)

type Users struct {
	UserID    uuid.UUID `json:"user_id"`
	Name      string    `json:"name"`
	Role      string    `json:"role"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	IsDeleted bool      `json:"is_deleted"`
}
