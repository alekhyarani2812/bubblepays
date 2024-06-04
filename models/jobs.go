package models

import "time"

type Jobs struct {
	ID         int        `json:"id"`
	Filename   string     `gorm:"column:file_name" json:"file_name"`
	Date       *time.Time `json:"date"`
	JobMessage string     `gorm:"column:job_message" json:"job_message"`
	Status     string     `json:"status"`
}

func (Jobs) TableName() string {
	return "job"
}
