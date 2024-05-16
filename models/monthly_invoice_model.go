package models

//MonthlyInvoice swagger:response MonthlyInvoice
type MonthlyInvoice struct {
	ID          int     `json:"-"`
	CardID      string  `json:"card_id"`
	Monthnum    string  `json:"monthnum"`
	Duration    float64 `json:"-"`
	Price       float64 `json:"-"`
	MonthlyBill float64 `json:"monthlybill"`
}

func (MonthlyInvoice) TableName() string {
	return "monthly_invoice"
}
