package model

import (
    "github.com/google/uuid"
)

type User struct {
    UserID     uuid.UUID `db:"user_id"`
    Name       string    `db:"name"`
    Role       string    `db:"role"`
    Email      string    `db:"email"`
    Password   string    `db:"password"`
    IsDeleted  bool      `db:"is_deleted"`
}
