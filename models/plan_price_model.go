package models

type PlanPrice struct {
	ID      int `json:"id"`
	PlanID  int `json:"plan_id"`
	PriceID int `json:"price_id"`
}

func (PlanPrice) TableName() string {
	return "plan_price"
}
