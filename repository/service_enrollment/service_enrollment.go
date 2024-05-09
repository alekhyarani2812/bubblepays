package service_enrollment

import (
	"context"
	"fmt"
	"log"
	"github.com/bubblepays/database"
	"github.com/bubblepays/models"
	"github.com/bubblepays/utils"
	"gorm.io/gorm"
)

func GetServiceEnrollment(ctx context.Context) ([]models.ServiceEnrollment, error) {
	dbConnection, err :=  utils.GetDBConnection(ctx)
	if err != nil {
		fmt.Println("error in db connection ", err.Error())
		return nil, err
	}
	var serviceEnrollRecords []models.ServiceEnrollment
	if err := dbConnection.Find(&serviceEnrollRecords).Error; err != nil {
		log.Printf("TransactionId: %v. Failed to retrieve consumption data: %v", err, utils.GetTransactionID(ctx))
		return nil, err
	}
	return serviceEnrollRecords, nil
}

func GetServiceEnrollmentById(ctx context.Context, pkId int) (*models.ServiceEnrollment, error) {
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	var serviceenroll models.ServiceEnrollment
	if err := dbConnection.Where("id = ?", pkId).First(&serviceenroll).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("record not found")
		} else {
			fmt.Printf("Error: %v\n", err)
		}
		return nil, err
	}
	return &serviceenroll, nil
}

func SaveServiceEnrollment(ctx context.Context, serviceEnrollment models.ServiceEnrollment) (*models.ServiceEnrollment, error) {
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	gorm := dbConnection.Create(&serviceEnrollment)
	if gorm.Error != nil {
		fmt.Println("gorm.Error :   error in inserting service enrollment")
	}
	return &serviceEnrollment, nil
}

func UpdateServiceEnrollment(ctx context.Context, pkId int, seviceEnrRecord models.ServiceEnrollment) (*models.ServiceEnrollment, error) {
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		fmt.Println("error to connect database", err.Error())
		return nil, err
	}
	gorm := dbConnection.Model(&models.ServiceEnrollment{}).Where("id = ?", pkId).Updates(&seviceEnrRecord)
	if gorm != nil && gorm.Error != nil {
		fmt.Println("error in updating data")
		return nil, err
	}
	gorm.Find(&seviceEnrRecord)
	fmt.Println("seviceEnrRecord :", seviceEnrRecord)
	return &seviceEnrRecord, nil
}

func DeleteServiceEnrollment(ctx context.Context, pkId int) (int64, error) {
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		fmt.Println("error in getting DB connection ", err.Error())
		return 0, nil
	}
	var serviceEnrollment models.ServiceEnrollment
	gorm := dbConnection.Where("id = ?", pkId).Delete(&serviceEnrollment)
	if gorm != nil && gorm.Error != nil {
		fmt.Println("error in deleting serviceEnrollment data ")
		return 0, err
	}
	if gorm.RowsAffected == 0 {
		fmt.Println("error in deleting serviceEnrollment data ")
		return 0, err
	}
	rows := gorm.RowsAffected
	return rows, nil
}
