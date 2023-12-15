package dto

type StudentRequestDto struct {
	StudentID   string `json:"student_id"`
	Fullname    string `json:"fullname"`
	BirthDate   string `json:"birth_date"`
	BirthPlace  string `json:"birth_place"`
	Address     string `json:"address"`
	Education   string `json:"education"`
	Institution string `json:"institution"`
	Job         string `json:"job"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsDeleted   bool   `json:"is_deleted"`
}
