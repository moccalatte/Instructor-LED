package dto

type UserRequestDto struct {
	Fullname  string `json:"fullname"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	IsDeleted bool   `json:"is_deleted"`
}
