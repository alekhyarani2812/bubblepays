package models

// Usages model
type Usages struct {
	ID        int     `json:"id"`
	CardID    string  `json:"card_id"`
	ServiceID string  `json:"service_id"`
	Monthnum  string  `json:"monthnum"`
	Duration  float64 `json:"duration"`
}
