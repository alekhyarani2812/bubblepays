package models

type ServiceEnrollment struct {
	ID        int    `json:"id"`
	CardID    string `json:"card_id"`
	ServiceID string `json:"service_id"`
}

func (ServiceEnrollment) TableName() string {
	return "service_enrollment"
}
