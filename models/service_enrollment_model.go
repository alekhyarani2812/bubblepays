package models

type ServiceEnrollment struct {
	Id        int    `json:"id"`
	CardId    string `json:"card_id"`
	ServiceId string `json:"service_id"`
}

func (ServiceEnrollment) TableName() string {
	return "service_enrollment"
}


