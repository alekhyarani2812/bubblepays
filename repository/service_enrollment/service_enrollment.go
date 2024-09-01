package serviceenrollment

import (
	"context"

	bcontext "github.com/bubblepays/context"
	"github.com/bubblepays/database"
	"github.com/bubblepays/log"
	"github.com/bubblepays/models"
	"gorm.io/gorm"
)
// GetServiceEnrollment retrieves all serviceEnrollment records from the table
func GetServiceEnrollment(ctx context.Context) ([]models.ServiceEnrollment, error) {
	dbConnection, err := bcontext.GetDBConnection(ctx)
	if err != nil {
		log.Errorft(ctx, "error in db connection %v ", err)
		return nil, err
	}
	var serviceEnrollRecords []models.ServiceEnrollment
	if err := dbConnection.Find(&serviceEnrollRecords).Error; err != nil {
		log.Infof("TransactionId: %v. Failed to retrieve serviceEnrollment Records data: %v", err, bcontext.GetTransactionID(ctx))
		return nil, err
	}
	return serviceEnrollRecords, nil
}
// GetServiceEnrollmentByID Retrieve the serviceEnrollments record based on the serviceEnrollmentID
func GetServiceEnrollmentByID(ctx context.Context, pkID int) (*models.ServiceEnrollment, error) {
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Errorft(ctx, "Error : %v ", err)
		return nil, err
	}
	var serviceenroll models.ServiceEnrollment
	if err := dbConnection.Where("id = ?", pkID).First(&serviceenroll).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Infof("record not found")
		} else {
			log.Infof("Error: %v\n", err)
		}
		return nil, err
	}
	return &serviceenroll, nil
}
// SaveServiceEnrollment for inserting serviceEnrollments into table
func SaveServiceEnrollment(ctx context.Context, serviceEnrollment models.ServiceEnrollment) (*models.ServiceEnrollment, error) {
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Errorft(ctx, "Error : %v ", err)
		return nil, err
	}
	gorm := dbConnection.Create(&serviceEnrollment)
	if gorm.Error != nil {
		log.Errorf("gorm.Error :   error in inserting service enrollment error : %v ", gorm.Error)
		return nil, gorm.Error
	}
	log.Infot(ctx, "service enrollment data saved successfully")
	return &serviceEnrollment, nil
}
// UpdateServiceEnrollment function is used to update the data in the serviceEnrollment table
func UpdateServiceEnrollment(ctx context.Context, pkID int, seviceEnrRecord models.ServiceEnrollment) (*models.ServiceEnrollment, error) {
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Errorf("error to connect database %v ", err.Error())
		return nil, err
	}
	gorm := dbConnection.Model(&models.ServiceEnrollment{}).Where("id = ?", pkID).Updates(&seviceEnrRecord)
	if gorm != nil && gorm.Error != nil {
		log.Infof("error in updating data")
		return nil, err
	}
	gorm.Find(&seviceEnrRecord)
	log.Infof("seviceEnrRecord : %v ", seviceEnrRecord)
	return &seviceEnrRecord, nil
}
// DeleteServiceEnrollment for deleting serviceenrollment record based on serviceenrollmentId
func DeleteServiceEnrollment(ctx context.Context, pkID int) (int64, error) {
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Infof("error in getting DB connection %v ", err.Error())
		return 0, nil
	}
	var serviceEnrollment models.ServiceEnrollment
	gorm := dbConnection.Where("id = ?", pkID).Delete(&serviceEnrollment)
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
