package models

import "time"

type Price struct {
	ID        int        `json:"id"`
	Price     float64    `json:"price"`
	Status    string     `json:"status"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func (Price) TableName() string {
	return "price"
}
