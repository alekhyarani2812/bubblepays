package repository

import (
	"fmt"

	"github.com/bubblepays/models"
	"gorm.io/gorm"
)

func SaveUsage(usageData models.Usages, dbConnection *gorm.DB) {
	fmt.Println("usageDtata : **************", usageData)
	//usageDtata models.Usages
	var usages []models.Usages
	res := dbConnection.Where("card_id = ?", usageData.CardId)
	res.Where("service_id = ?", usageData.ServiceId)
	res.Where("monthnum = ?", usageData.Monthnum)
	res.Find(&usages)

	if res.Error != nil {
		fmt.Println("error in this block............", res.Error)
		return
	}
	if len(usages) > 0 {
		//should run update query
		fmt.Println("******update query********records found len(usages) : ", len(usages))
		usage := usages[0]
		fmt.Println("usage : ............", usage.Duration)
		fmt.Println("usageData.Duration : ......", usageData.Duration)
		dbConnection.Model(&usage).Update("duration", usageData.Duration)
	} else {
		//should run insert query
		fmt.Println("**************records not found should run insert query ")
		ormDB := dbConnection.Create(&usageData)
		fmt.Println("**************ormDB usages table rows Inserted ", ormDB.RowsAffected)
	}
}
