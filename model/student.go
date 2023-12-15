package model

import "time"

type Student struct {
	StudentID   string    `json:"student_id"`
	Fullname    string    `json:"fullname"`
	BirthDate   string    `json:"birth_date"`
	BirthPlace  string    `json:"birth_place"`
	Address     string    `json:"address"`
	Education   string    `json:"education"`
	Institution string    `json:"institution"`
	Job         string    `json:"job"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Role        string    `json:"role"`
	IsDeleted   bool      `json:"is_deleted"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (u Users) IsValidRoleStudent() bool {
	return u.Role == "student"
}
