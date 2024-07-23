package repository

import (
	"github.com/bubblepays/log"
	"github.com/bubblepays/models"
	"gorm.io/gorm"
)

func SaveUsage(usageData models.Usages, dbConnection *gorm.DB) {
	var usages []models.Usages
	res := dbConnection.Where("card_id = ?", usageData.CardID)
	res.Where("service_id = ?", usageData.ServiceID)
	res.Where("monthnum = ?", usageData.Monthnum)
	res.Find(&usages)

	if res.Error != nil {
		log.Infof("error in this block.... %v ", res.Error)
		return
	}
	if len(usages) > 0 {
		//should run update query
		//log.Info( "******update query********records found len(usages) : %v ", len(usages))
		usage := usages[0]
		// log.Info( "usage : ............", usage.Duration)
		// log.Info( "usageData.Duration : ......", usageData.Duration)
		dbConnection.Model(&usage).Update("duration", usageData.Duration)
		dbConnection.Find(&usage)
	} else {
		//should run insert query
		ormDB := dbConnection.Create(&usageData)
		log.Infof("**************ormDB usages table rows Inserted %v ", ormDB.RowsAffected)
	}
}
