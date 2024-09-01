package models

import "time"

// Price model
type Price struct {
	ID        int        `json:"id"`
	Price     float64    `json:"price"`
	Status    string     `json:"status"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
// TableName overrides the table name used by User to `profiles`
func (Price) TableName() string {
	return "price"
}
