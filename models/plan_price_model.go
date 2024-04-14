package models

type PlanPrice struct {
	Id      int `json:"id"`
	PlanId  int `json:"plan_id"`
	PriceId int `json:"price_id"`
}

func (PlanPrice) TableName() string {
	return "plan_price"
}
