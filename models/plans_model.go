package models

import "time"

type Plans struct {
	ID        int        `json:"id"`
	PlanName  string     `json:"plan_name"`
	Status    string     `json:"status"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
