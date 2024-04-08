package consumption

import (
	"log"
	"github.com/bubblepays/database"
	"github.com/bubblepays/models"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

// Get all consumptions data
func GetConsumptions(dbConnection *gorm.DB) ([]models.Consumption, error) {
	// Retrieve consumption data from the database
	var consumptions []models.Consumption
	if err := dbConnection.Find(&consumptions).Error; err != nil {
		log.Printf("failed to retrieve consumption data: %v", err)
		return nil, err
	}
	log.Printf("SIZE of consumption: size:: %v ", len(consumptions))
	return consumptions, nil
}

// Get consumption data by Id
func GetConsumptionById(id int) ([]models.Consumption, error) {
	// get orm db connection
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Printf("err %v ", err)
		return nil, err
	}
	// Retrieve consumption data from the database
	var consumption []models.Consumption
	if err := dbConnection.Where("id = ?", id).First(&consumption).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("Employee not found")
		} else {
			log.Printf("Error: %v\n", err)
		}
		return nil, err
	}
	// Employee found, print details
	log.Printf("consumption found: %+v\n", consumption)
	log.Printf("SIZE of consumption: size:: %v ", len(consumption))
	return consumption, nil
}

// Insert one consumption data
func SavesConsumption(consumRecord models.Consumption) (*models.Consumption, error) {
	log.Println("consumRecord", consumRecord)
	// get orm db connection
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Println("err in dbConnection :****")
		return nil, err
	}
	log.Println("dbConnection  :")
	log.Println("consum from repo : ", consumRecord)
	ormDB := dbConnection.Create(&consumRecord)
	if ormDB.Error != nil {
		log.Println("Error inserting employee:", ormDB.Error)
		return nil, ormDB.Error
	}
	log.Println("Employee inserted successfully")
	log.Println("ormDB  :  ", ormDB)
	return &consumRecord, nil
}

func DeleteConsumption(pkId int) (int64, error) {
	log.Println("pkId : ", pkId)
	// get orm db connection
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Println("err in dbConnection :****")
		return 0, err
	}
	log.Println("dbConnection  :")
	var consumption models.Consumption

	result := dbConnection.Where("id = ?", pkId).Delete(&consumption)
	if result.Error != nil {
		log.Println("failed to delete consumption")
	}
	log.Println("return result from repo to Handler")
	rowsEff := result.RowsAffected
	return rowsEff, nil
}
