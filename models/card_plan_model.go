package models

type CardPlan struct {
	ID     int    `json:"id"`
	CardID string `json:"card_id"`
	PlanID int    `json:"plan_id"`
}

func (CardPlan) TableName() string {
	return "card_plan"
}
