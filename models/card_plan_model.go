package models
// CardPlan model
type CardPlan struct {
	ID     int    `json:"id"`
	CardID string `json:"card_id"`
	PlanID int    `json:"plan_id"`
}
// TableName overrides the table name used by User to `profiles`
func (CardPlan) TableName() string {
	return "card_plan"
}
