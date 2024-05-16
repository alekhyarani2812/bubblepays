package consumption

import (
	"context"

	bcontext "github.com/bubblepays/context"
	"github.com/bubblepays/database"
	"github.com/bubblepays/log"
	"github.com/bubblepays/models"
	
	"gorm.io/gorm"
)

//GetConsumptions Get all consumptions data
func GetConsumptions(ctx context.Context) ([]models.Consumption, error) {
	//get dbconnection from contex
	dbConnection, err := bcontext.GetDBConnection(ctx)
	if err != nil {
		log.Errorft(ctx, "error in db connection %v ", err)
		return nil, err
	}
	var consumptions []models.Consumption
	if err := dbConnection.Find(&consumptions).Error; err != nil {
		log.Infof("failed to retrieve consumption data: %v", err)
		return nil, err
	}
	log.Infof("consumptions : %v ", consumptions)
	return consumptions, nil
}

func GetConsumptionByID(ctx context.Context, id int) ([]models.Consumption, error) {
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Infof("err %v ", err)
		return nil, err
	}
	var consumption []models.Consumption
	if err := dbConnection.Where("id = ?", id).First(&consumption).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Info("Employee not found")
		} else {
			log.Infof("Error: %v\n", err)
		}
		return nil, err
	}
	log.Infof("consumption : %v ", consumption)
	return consumption, nil
}

func SavesConsumption(ctx context.Context, consumRecord models.Consumption) (*models.Consumption, error) {
	log.Infof("consumRecord %v ", consumRecord)
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Info("err in dbConnection :****")
		return nil, err
	}
	ormDB := dbConnection.Create(&consumRecord)
	if ormDB.Error != nil {
		log.Infof("Error inserting employee: %v ", ormDB.Error)
		return nil, ormDB.Error
	}
	return &consumRecord, nil
}

func DeleteConsumption(ctx context.Context, pkID int) (int64, error) {
	log.Infof("pkID : %v ", pkID)
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Info("err in dbConnection :****")
		return 0, err
	}
	var consumption models.Consumption
	result := dbConnection.Where("id = ?", pkID).Delete(&consumption)
	if result.Error != nil {
		log.Info("failed to delete consumption")
	}
	rowsEff := result.RowsAffected
	return rowsEff, nil
}
