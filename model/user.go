package model

import "time"

type Users struct {
	UserID    string    `json:"user_id"`
	Fullname  string    `json:"fullname"`
	Role      string    `json:"role"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	IsDeleted bool      `json:"is_deleted"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (u Users) IsValidRole() bool {
	return u.Role == "admin" || u.Role == "trainer"
}
