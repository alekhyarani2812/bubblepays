package monthlyinvoice

import (
	"context"

	bcontext "github.com/bubblepays/context"
	"github.com/bubblepays/log"
	"github.com/bubblepays/models"
)

func GetMonthlyInvoice(ctx context.Context, monthNumQparam string) ([]*models.MonthlyInvoice, error) {

	log.Infoft(ctx, "This is GetMonthlyInvoice monthlyBillQparam: %v", monthNumQparam)
	log.Infof("*************** hello world")
	dbConnection, err := bcontext.GetDBConnection(ctx)
	if err != nil {
		log.Errort(ctx, "this is error")
		return nil, err
	}
	var usages []models.Usages
	var monthlyInvoices []*models.MonthlyInvoice
	gorm := dbConnection.Joins("JOIN card_plan ON usages.card_id = card_plan.card_id").
		Joins("JOIN plan_price ON card_plan.plan_id = plan_price.plan_id").
		Joins("JOIN plans ON plan_price.plan_id = plans.id").
		Joins("JOIN price ON plan_price.price_id = price.id").
		Select("usages.card_id as card_id, usages.monthnum as monthnum, usages.duration as duration, price.price as price").
		Where("usages.monthnum = ?", monthNumQparam).
		Find(&usages).Find(&monthlyInvoices)
	if gorm.Error != nil && gorm != nil {
		errors := gorm.Error
		log.Errorft(ctx, "error in getting records from usages %v", errors)
		return nil, errors
	}
	log.Infoft(ctx, "usages : %v ", usages)
	for _, monthlyInvoice := range monthlyInvoices {

		monthlyInvoice.MonthlyBill = monthlyInvoice.Duration * monthlyInvoice.Price
		log.Infoft(ctx, "monthlyInvoice.MonthlyBill %v", monthlyInvoice.MonthlyBill)
	}
	log.Infoft(ctx, "monthlyInvoices : %v", monthlyInvoices)
	return monthlyInvoices, nil
}
