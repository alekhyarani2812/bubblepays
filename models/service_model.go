package models

// Services model
type Services struct {
	ID          int    `json:"id"`
	ServiceID   string `json:"service_id"`
	Active      bool   `json:"active"`
	Location    string `json:"location"`
	ServiceName string `json:"service_name"`
}
