package models

import "time"

type Consumption struct {
	ID              int        `json:"id"`
	DeviceID        string     `json:"device_id"`
	CardID          string     `json:"card_id"`
	ConsumptionDate string     `json:"consumption_date"`
	StartTime       *time.Time `json:"start_time"`
	EndTime         *time.Time `json:"end_time"`
	ServiceID       string     `json:"service_id"`
	CreatedDate     *time.Time `json:"created_date"`
	CreatedBy       string     `json:"created_by"`

	// Mention the table name explicitly  because gorm takes plural table name by default
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (Consumption) TableName() string {
	return "consumption"
}
