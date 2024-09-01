package planprice

import (
	"context"

	bcontext "github.com/bubblepays/context"
	"github.com/bubblepays/database"
	"github.com/bubblepays/log"
	"github.com/bubblepays/models"
	"gorm.io/gorm"
)
// GetPlanPrice retrieves all planPrice records from the table
func GetPlanPrice(ctx context.Context) ([]models.PlanPrice, error) {
	//get dbconnection from contex
	dbConnection, err := bcontext.GetDBConnection(ctx)
	if err != nil {
		log.Errorft(ctx, "error in db connection %v ", err)
		return nil, err
	}
	var planPrice []models.PlanPrice
	gorm := dbConnection.Find(&planPrice)
	if gorm.Error != nil {
		log.Error("error in getting records from planPrice")
		return nil, err
	}
	log.Infof("planPrice : %v", planPrice)
	return planPrice, nil
}
// GetPlanPriceByID Retrieve the planprice record based on the planpriceId
func GetPlanPriceByID(ctx context.Context, pkID int) (*models.PlanPrice, error) {
	log.Infof("this is GetPlanPriceByID repo")
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Errorft(ctx, "error in db connection %v ", err)
		return nil, err
	}
	var planPrice models.PlanPrice
	if err := dbConnection.Where("id = ?", pkID).First(&planPrice).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Info("plan not found")
		} else {
			log.Infof("Error: %v\n", err)
		}
		return nil, err
	}
	log.Infof("planPrice : %v", planPrice)
	return &planPrice, nil
}
// SavePlanPrice for inserting priceplans into table
func SavePlanPrice(ctx context.Context, planPriceData *models.PlanPrice) (*models.PlanPrice, error) {
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Errorft(ctx, "error in db connection %v ", err)
		return nil, err
	}
	gorm := dbConnection.Create(&planPriceData)
	if gorm.Error != nil && gorm != nil {
		log.Infof("error in insertion of records %v ", gorm.Error)
		return nil, err
	}
	log.Infof("planPriceData : %v", planPriceData)
	return planPriceData, nil
}
// UpdatePlanPrice function is used to update the data in the planprice table
func UpdatePlanPrice(ctx context.Context, pkID int, planPriceRecord *models.PlanPrice) (*models.PlanPrice, error) {
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Errorft(ctx, "error in db connection %v ", err)
		return nil, err
	}
	gorm := dbConnection.Model(&models.PlanPrice{}).Where("id = ?", pkID).Updates(&planPriceRecord)
	if gorm != nil && gorm.Error != nil {
		log.Infof("error in updating data ")
		return nil, err
	}
	gorm.Find(&planPriceRecord)
	log.Infof("planPriceRecord : %v ", planPriceRecord)
	return planPriceRecord, nil
}
// DeletePlanPrice for deleting planprice record based on planpriceId
func DeletePlanPrice(ctx context.Context, pkID int) (int64, error) {
	log.Infof("this is DeletePlanPrice repo")
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Errorft(ctx, "error in db connection %v ", err)
		return 0, err
	}
	var planPrice models.PlanPrice
	gorm := dbConnection.Where("id = ?", pkID).Delete(&planPrice)
	if gorm.Error != nil && gorm != nil {
		log.Infof("error in deleting PlanPrice data")
		return 0, err
	}
	rows := gorm.RowsAffected
	if rows == 0 {
		log.Infof("error in deleting PlanPrice data ")
		return 0, err
	}
	log.Infof("rows : %v  ", rows)
	return rows, nil
}
