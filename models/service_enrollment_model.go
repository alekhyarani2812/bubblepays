package models

type Service_Enrollment struct {
	Id        int    `json:"id"`
	CardId    string `json:"card_id"`
	ServiceId string `json:"service_id"`
}
