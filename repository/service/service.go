package service

import (
	"context"

	bcontext "github.com/bubblepays/context"
	"github.com/bubblepays/database"
	"github.com/bubblepays/log"
	"github.com/bubblepays/models"
)

//GetServices for Read all the services
func GetServices(ctx context.Context) ([]models.Services, error) {
	dbConnection, err := bcontext.GetDBConnection(ctx)
	if err != nil {
		log.Errorft(ctx, "error in db connection %v ", err)
		return nil, err
	}
	var service []models.Services
	if err := dbConnection.Find(&service).Error; err != nil {
		log.Infof("failed to retrieve consumption data: %v", err)
		return nil, err
	}
	log.Infof("SIZE of consumption: size:: %v ", len(service))
	return service, nil
}

func SaveService(ctx context.Context, service models.Services) (*models.Services, error) {
	log.Infof("serviceData in service repo", service)
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Infof("err in dbConnection :****")
		return nil, err
	}
	ormDB := dbConnection.Create(&service)
	if ormDB.Error != nil {
		log.Infof("Error inserting service: %v ", ormDB.Error)
		return nil, ormDB.Error
	}
	log.Infof("Employee inserted successfully")
	return &service, nil
}

func DeleteServiceByID(ctx context.Context, pkID int) (int64, error) {
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Infof("error in db connecton  ")
		return 0, err
	}
	var service models.Services
	ormDb := dbConnection.Where("id=?", pkID).Delete(&service)
	rowsEffected := ormDb.RowsAffected
	return rowsEffected, nil
}
