package cardplan

import (
	"context"

	bcontext "github.com/bubblepays/context"
	"github.com/bubblepays/database"
	"github.com/bubblepays/log"
	"github.com/bubblepays/models"
)

func GetCardPlan(ctx context.Context) ([]models.CardPlan, error) {
	//get dbconnection from contex
	dbConnection, err := bcontext.GetDBConnection(ctx)
	if err != nil {
		log.Infof("error in db connection %v", err.Error())
		return nil, err
	}
	var cardplan []models.CardPlan
	gorm := dbConnection.Find(&cardplan)
	if gorm.Error != nil && gorm != nil {
		errors := gorm.Error
		log.Infof("error in getting records from cardplan %v ", errors)
		return nil, errors
	}
	log.Infof("cardplan : %v ", cardplan)
	return cardplan, nil
}

func GetCardPlanByID(ctx context.Context, pkID int) (*models.CardPlan, error) {
	log.Infof("GetCardPlanByID repo ************")
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Errorft(ctx, "error in db connection %v ", err)
		return nil, err
	}
	var cardPlanRecord models.CardPlan
	gorm := dbConnection.Where("id = ?", pkID).Find(&cardPlanRecord)
	if gorm != nil && gorm.Error != nil {
		log.Infof("error in getting CardPlan data")
		return nil, err
	}
	log.Infof(" cardplan record : %v ", cardPlanRecord)
	return &cardPlanRecord, nil
}

func SaveCardPlan(ctx context.Context, cardPlanData *models.CardPlan) (*models.CardPlan, error) {
	log.Infof("this is CardPlan repo  CardPlan : %v ", cardPlanData)
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Errorft(ctx, "error in db connection %v ", err)
		return nil, err
	}
	gorm := dbConnection.Create(&cardPlanData)
	if gorm.Error != nil && gorm != nil {
		log.Infof("error in insertion of records %v", gorm.Error)
		return nil, err
	}
	log.Infof("cardPlanData : %v ", cardPlanData)
	return cardPlanData, nil
}

func UpdateCardPlan(ctx context.Context, pkID int, cardPlanRecord *models.CardPlan) (*models.CardPlan, error) {
	log.Infof("this is cardplan repo cardPlanRecord : %v ,pkID %v", cardPlanRecord, pkID)
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Infof("error in db connection ")
		return nil, err
	}
	//cardPlanRecord
	gorm := dbConnection.Model(&models.CardPlan{}).Where("id = ?", pkID).Updates(&cardPlanRecord)
	if gorm != nil && gorm.Error != nil {
		log.Infof("error in updating cardPlan Record ")
		return nil, err
	}
	gorm.Find(&cardPlanRecord)
	log.Infof("cardPlanRecord : %v ", cardPlanRecord)
	return cardPlanRecord, nil
}

func DeleteCardPlan(ctx context.Context, pkID int) (int64, error) {
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Infof("error in db connection ")
		return 0, err
	}
	var cardPlanRecord models.CardPlan
	gorm := dbConnection.Where("id = ?", pkID).Delete(&cardPlanRecord)
	if gorm.Error != nil {
		log.Infof("error in deleting card PlanRecord *****gorm.Error != nil*****")
		return 0, err
	}
	rows := gorm.RowsAffected
	if rows == 0 {
		log.Infof("error in deleting cardPlan Record")
		return 0, err
	}
	log.Infof("rows : %v ", rows)
	return rows, nil
}
