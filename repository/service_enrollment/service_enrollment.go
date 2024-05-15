package service_enrollment

import (
	"context"

	bcontext "github.com/bubblepays/context"
	"github.com/bubblepays/database"
	"github.com/bubblepays/log"
	"github.com/bubblepays/models"
	"gorm.io/gorm"
)

func GetServiceEnrollment(ctx context.Context) ([]models.ServiceEnrollment, error) {
	dbConnection, err := bcontext.GetDBConnection(ctx)
	if err != nil {
		log.Errorft(ctx, "error in db connection %v ", err)
		return nil, err
	}
	var serviceEnrollRecords []models.ServiceEnrollment
	if err := dbConnection.Find(&serviceEnrollRecords).Error; err != nil {
		log.Infof("TransactionId: %v. Failed to retrieve consumption data: %v", err, bcontext.GetTransactionID(ctx))
		return nil, err
	}
	return serviceEnrollRecords, nil
}

func GetServiceEnrollmentById(ctx context.Context, pkId int) (*models.ServiceEnrollment, error) {
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Errorft(ctx, "Error : %v ", err)
		return nil, err
	}
	var serviceenroll models.ServiceEnrollment
	if err := dbConnection.Where("id = ?", pkId).First(&serviceenroll).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Infof("record not found")
		} else {
			log.Infof("Error: %v\n", err)
		}
		return nil, err
	}
	return &serviceenroll, nil
}

func SaveServiceEnrollment(ctx context.Context, serviceEnrollment models.ServiceEnrollment) (*models.ServiceEnrollment, error) {
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Errorft(ctx, "Error : %v ", err)
		return nil, err
	}
	gorm := dbConnection.Create(&serviceEnrollment)
	if gorm.Error != nil {
		log.Infof("gorm.Error :   error in inserting service enrollment")
	}
	return &serviceEnrollment, nil
}

func UpdateServiceEnrollment(ctx context.Context, pkId int, seviceEnrRecord models.ServiceEnrollment) (*models.ServiceEnrollment, error) {
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Infof("error to connect database", err.Error())
		return nil, err
	}
	gorm := dbConnection.Model(&models.ServiceEnrollment{}).Where("id = ?", pkId).Updates(&seviceEnrRecord)
	if gorm != nil && gorm.Error != nil {
		log.Infof("error in updating data")
		return nil, err
	}
	gorm.Find(&seviceEnrRecord)
	log.Infof("seviceEnrRecord :", seviceEnrRecord)
	return &seviceEnrRecord, nil
}

func DeleteServiceEnrollment(ctx context.Context, pkId int) (int64, error) {
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Infof("error in getting DB connection %v ", err.Error())
		return 0, nil
	}
	var serviceEnrollment models.ServiceEnrollment
	gorm := dbConnection.Where("id = ?", pkId).Delete(&serviceEnrollment)
	if gorm != nil && gorm.Error != nil {
		log.Infof("error in deleting serviceEnrollment data ")
		return 0, err
	}
	if gorm.RowsAffected == 0 {
		log.Infof("error in deleting serviceEnrollment data ")
		return 0, err
	}
	rows := gorm.RowsAffected
	return rows, nil
}
