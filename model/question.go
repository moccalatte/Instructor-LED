package model

import (
    "github.com/google/uuid"
)

type Question struct {
    QuestionID uuid.UUID `json:"question_id"`
    UserID      uuid.UUID `json:"user_id"`
    Questionary string    `json:"questionary"`
    Status      bool      `json:"status"`
    IsDeleted   bool      `json:"is_deleted"`
}
