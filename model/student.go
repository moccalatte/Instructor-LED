package model

import (
    "time"
    "github.com/google/uuid"
)

type Student struct {
    StudentID  uuid.UUID `db:"student_id"`
    Fullname   string    `db:"fullname"`
    Shortname  string    `db:"shortname"`
    BirthDate  time.Time `db:"birth_date"`
    BirthPlace string    `db:"birth_place"`
    Address    string    `db:"address"`
    Education  string    `db:"education"`
    Institution string  `db:"institution"`
    Job        string    `db:"job"`
    Email      string    `db:"email"`
    Password   string    `db:"password"`
    IsDeleted  bool      `db:"is_deleted"`
}
