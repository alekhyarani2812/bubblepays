package models

type Services struct {
	Id          int    `json:"id"`
	ServiceId   string `json:"service_id"`
	Active      bool   `json:"active"`
	Location    string `json:"location"`
	ServiceName string `json:"service_name"`
}

type Tables interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (Services) TableName() string {
	return "services"
}
