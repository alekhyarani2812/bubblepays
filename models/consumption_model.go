package models

import "time"

type Consumption struct {
	Id          int        `json:"id"`
	DeviceId    string     `json:"device_id"`
	CardId      int        `json:"card_id"`
	Date        int        `json:"date"`
	StartTime   *time.Time `json:"start_time"`
	EndTime     *time.Time `json:"end_time"`
	ServiceId   int        `json:"service_id"`
	ServiceName string     `json:"service_name"`
	CreatedDate *time.Time `json:"created_date"`
	CreatedBy   string     `json:"created_by"`

	// Mention the table name explicitly
	//TableName string `gorm:"tableName:consumption"`
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (Consumption) TableName() string {
	return "consumption"
}
