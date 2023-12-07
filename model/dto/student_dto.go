package dto

import (
	"time"
)

type StudentRequestDto struct {
	Fullname    string    `json:"fullname"`
	Shortname   string    `json:"shortname"`
	BirthDate   time.Time `json:"birth_date"`
	BirthPlace  string    `json:"birth_place"`
	Address     string    `json:"address"`
	Education   string    `json:"education"`
	Institution string    `json:"institution"`
	Job         string    `json:"job"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	IsDeleted   bool      `json:"is_deleted"`
}
