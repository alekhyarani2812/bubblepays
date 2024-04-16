package models

import "time"

type Consumption struct {
	Id              int        `json:"id"`
	DeviceId        string     `json:"device_id"`
	CardId          string     `json:"card_id"`
	ConsumptionDate string     `json:"consumption_date"`
	StartTime       *time.Time `json:"start_time"`
	EndTime         *time.Time `json:"end_time"`
	ServiceId       int        `json:"service_id"`
	ServiceName     string     `json:"service_name"`
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
