package models

// ServiceEnrollment model
type ServiceEnrollment struct {
	ID        int    `json:"id"`
	CardID    string `json:"card_id"`
	ServiceID string `json:"service_id"`
}
// TableName overrides the table name used by User to `profiles`
func (ServiceEnrollment) TableName() string {
	return "service_enrollment"
}
