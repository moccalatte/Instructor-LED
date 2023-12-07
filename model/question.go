package model

import (
    "github.com/google/uuid"
)

type Question struct {
    QuestionsID uuid.UUID `json:"questions_id"`
    UserID      uuid.UUID `json:"user_id"`
    Questionary string    `json:"questionary"`
    Status      bool      `json:"status"`
    IsDeleted   bool      `json:"is_deleted"`
}
