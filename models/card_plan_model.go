package models

type Card_Plan struct {
	Id     int    `json:"id"`
	CardId string `json:"card_id"`
	PlanId int    `json:"plan_id"`
}
