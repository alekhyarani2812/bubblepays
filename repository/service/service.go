package service

import (
	"log"
	"todoproject/database"
	"todoproject/models"
)
//Read services
func GetServices() ([]models.Services, error){
	log.Println("This is GetServices repo")
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Println("err in dbConnection :****")
		return nil,err
	}
	log.Println("dbConnection  :")
	// Retrieve consumption data from the database
	var service []models.Services
	if err := dbConnection.Find(&service).Error; err != nil {
		log.Printf("failed to retrieve consumption data: %v", err)
		return nil, err
	}
	log.Printf("SIZE of consumption: size:: %v ", len(service))
	return service, nil
}

//Insert services
func SaveService(service models.Services) (*models.Services, error){
	 log.Println("serviceData in service repo", service)
	 // get orm db connection
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Println("err in dbConnection :****")
		return nil,err
	}
	log.Println("dbConnection  :")
	//var serve models.Service
	 log.Println("serviceRecord from repo : ", service)
	 ormDB := dbConnection.Create(&service)
	if ormDB.Error != nil {
		log.Println("Error inserting service:", ormDB.Error)
		return nil, ormDB.Error
	}
	log.Println("Employee inserted successfully")
	log.Println("ormDB  :  ", ormDB)
	return &service, nil
}

//Delete service By Id
func DeleteServiceById(){
	log.Println("This is DeleteServiceById handler")
}
