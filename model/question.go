package model

import (
    "github.com/google/uuid"
)

type Question struct {
    QuestionsID uuid.UUID `db:"questions_id"`
    UserID      uuid.UUID `db:"user_id"`
    Questionary string    `db:"questionary"`
    Status      bool      `db:"status"`
    IsDeleted   bool      `db:"is_deleted"`
}
