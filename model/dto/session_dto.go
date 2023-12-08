package dto

type SessionRequestDto struct {
	// SessionId string `json:"session_id"`
	Date      string `json:"date"`
	TrainerId string `json:"trainer_id"`
	IsDeleted bool   `json:"is_deleted"`
}
