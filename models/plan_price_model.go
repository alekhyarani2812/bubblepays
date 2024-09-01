package models

// PlanPrice model
type PlanPrice struct {
	ID      int `json:"id"`
	PlanID  int `json:"plan_id"`
	PriceID int `json:"price_id"`
}
// TableName overrides the table name used by User to `profiles`
func (PlanPrice) TableName() string {
	return "plan_price"
}
