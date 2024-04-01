package models

type Usages struct {
	Id        int     `json:"id"`
	CardId    string  `json:"card_id"`
	ServiceId string  `json:"service_id"`
	Monthnum  string  `json:"monthnum"`
	Duration  float64 `json:"duration"`
}
