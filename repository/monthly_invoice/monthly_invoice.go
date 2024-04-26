package monthlyinvoice

import (
	"context"
	"fmt"

	"github.com/bubblepays/models"
	"github.com/bubblepays/utils"
)

func GetMonthlyInvoice(ctx context.Context) ([]models.MonthlyInvoice, error) {
	fmt.Println("This is GetMonthlyInvoice repo")
	dbConnection, err := utils.GetDBConnection(ctx)
	if err != nil {
		fmt.Println("error in db connection ", err.Error())
		return nil, err
	}
	var usages []models.Usages
	var monthlyInvoice []models.MonthlyInvoice
	gorm := dbConnection.Joins("JOIN card_plan ON usages.card_id = card_plan.card_id").
		Joins("JOIN plan_price ON card_plan.plan_id = plan_price.plan_id").
		Joins("JOIN plans ON plan_price.plan_id = plans.id").
		Joins("JOIN price ON plan_price.price_id = price.id").
		Select("usages.card_id as card_id, usages.monthnum as monthnum, usages.duration as duration, price.price as price").
		Find(&usages).Find(&monthlyInvoice)
	if gorm.Error != nil && gorm != nil {
		errors := gorm.Error
		fmt.Println("error in getting records from usages ", errors)
		return nil, errors
	}
	fmt.Println("usages : ", usages)
	fmt.Println("monthlyInvoice : ", monthlyInvoice)
	return monthlyInvoice, nil
}
