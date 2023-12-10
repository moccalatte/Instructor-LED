package dto

type AuthRequestDto struct {
	Fullname string `json:"fullname"`
	Password string `json:"password"`
}

type AuthResponseDto struct {
	Token string `json:"token"`
}
