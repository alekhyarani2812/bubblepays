package models

type CardPlan struct {
	Id     int    `json:"id"`
	CardId string `json:"card_id"`
	PlanId int    `json:"plan_id"`
}

func (CardPlan) TableName() string {
	return "card_plan"
}
