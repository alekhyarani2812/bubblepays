package service

import (
	"log"

	"github.com/bubblepays/database"
	"github.com/bubblepays/models"
)

// Read services
func GetServices() ([]models.Services, error) {
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Println("err in dbConnection :****")
		return nil, err
	}
	var service []models.Services
	if err := dbConnection.Find(&service).Error; 
	err != nil {
		log.Printf("failed to retrieve consumption data: %v", err)
		return nil, err
	}
	log.Printf("SIZE of consumption: size:: %v ", len(service))
	return service, nil
}

func SaveService(service models.Services) (*models.Services, error) {
	log.Println("serviceData in service repo", service)
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Println("err in dbConnection :****")
		return nil, err
	}
	ormDB := dbConnection.Create(&service)
	if ormDB.Error != nil {
		log.Println("Error inserting service:", ormDB.Error)
		return nil, ormDB.Error
	}
	log.Println("Employee inserted successfully")
	return &service, nil
}

func DeleteServiceById(pkId int) (int64, error) {
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Println("error in db connecton  ")
		return 0, err
	}
	var service models.Services
	ormDb := dbConnection.Where("id=?", pkId).Delete(&service)
	rowsEffected := ormDb.RowsAffected
	return rowsEffected, nil
}
