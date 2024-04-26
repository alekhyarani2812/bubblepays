package models

type MonthlyInvoice struct {
	Id       int     `json:"id"`
	CardId   string  `json:"card_id"`
	Monthnum string  `json:"monthnum"`
	Duration float64 `json:"duration"`
	Price    float64 `json:"price"`
}

func (MonthlyInvoice) TableName() string {
	return "monthly_invoice"
}
