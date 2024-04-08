package models

type Services struct {
	Id          int    `json:"id"`
	ServiceId   string `json:"service_id"`
	Active      bool   `json:"active"`
	Location    string `json:"location"`
	ServiceName string `json:"service_name"`
}
