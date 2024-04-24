package consumption

import (
	"context"
	"fmt"
	"log"

	"github.com/bubblepays/database"
	"github.com/bubblepays/models"
	"github.com/bubblepays/utils"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

// Get all consumptions data
func GetConsumptions(ctx context.Context) ([]models.Consumption, error) {
	//get dbconnection from contex 
	dbConnection, err :=  utils.GetDBConnection(ctx)
	if err != nil {
		fmt.Println("error in db connection ", err.Error())
		return nil, err
	}
	var consumptions []models.Consumption
	if err := dbConnection.Find(&consumptions).Error; err != nil {
		log.Printf("failed to retrieve consumption data: %v", err)
		return nil, err
	}
	log.Println("consumptions : ", consumptions)
	return consumptions, nil
}

func GetConsumptionById(ctx context.Context, id int) ([]models.Consumption, error) {
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Printf("err %v ", err)
		return nil, err
	}
	var consumption []models.Consumption
	if err := dbConnection.Where("id = ?", id).First(&consumption).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("Employee not found")
		} else {
			log.Printf("Error: %v\n", err)
		}
		return nil, err
	}
	log.Println("consumption : ", consumption)
	return consumption, nil
}

func SavesConsumption(ctx context.Context, consumRecord models.Consumption) (*models.Consumption, error) {
	log.Println("consumRecord", consumRecord)
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Println("err in dbConnection :****")
		return nil, err
	}
	ormDB := dbConnection.Create(&consumRecord)
	if ormDB.Error != nil {
		log.Println("Error inserting employee:", ormDB.Error)
		return nil, ormDB.Error
	}
	return &consumRecord, nil
}

func DeleteConsumption(ctx context.Context, pkId int) (int64, error) {
	log.Println("pkId : ", pkId)
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Println("err in dbConnection :****")
		return 0, err
	}
	var consumption models.Consumption
	result := dbConnection.Where("id = ?", pkId).Delete(&consumption)
	if result.Error != nil {
		log.Println("failed to delete consumption")
	}
	rowsEff := result.RowsAffected
	return rowsEff, nil
}
