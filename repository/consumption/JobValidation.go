package consumption

import (
	"context"
	"fmt"
	"errors"
	bcontext "github.com/bubblepays/context"
	"github.com/bubblepays/log"
	"github.com/bubblepays/models"
)

func JobValidation(ctx context.Context, consumptionData models.Consumption) (bool, error) {
	dbConnection, err := bcontext.GetDBConnection(ctx)
	if err != nil {
		log.Errorft(ctx, "error in db connection %v ", err)
		return true, err
	}
	// validate cardId, serviceId in service enrolment table
	var serviceenrollment models.ServiceEnrollment
	serviceenrolref := &serviceenrollment
	gorm := dbConnection.Where("card_id=?", consumptionData.CardID).
		Where("service_id=?", consumptionData.ServiceID).
		Find(&serviceenrolref)
	if gorm != nil && gorm.Error != nil {
		log.Errorft(ctx, "error while getting user data %v",gorm.Error)
		return false, gorm.Error
	}
	if serviceenrolref.CardID == "" || serviceenrolref.ServiceID == "" {
		message := fmt.Sprintf("There is no service enrollment for this cardId : %v and Service Id : %v ", consumptionData.CardID, consumptionData.ServiceID) 
		errMessage := errors.New(message)
		//log.Errorft(ctx, "There is no service enrollment for this cardId : %v and Service Id : %v ", consumptionData.CardID, consumptionData.ServiceID)
		return false, errMessage
	}
	// validate deviceId, serviceId in deviceservices table
	var deviceservices models.DeviceServices
	deviceserviceref := &deviceservices
	gorms := dbConnection.Where("device_id=?", consumptionData.DeviceID).
		Where("service_id=?", consumptionData.ServiceID).
		Find(deviceserviceref)
	if gorms != nil && gorms.Error != nil {
		log.Errorft(ctx, "error while getting device data")
		return false, gorms.Error
	}
	if deviceserviceref.DeviceID == "" || deviceserviceref.ServiceID == "" {
		msg := fmt.Sprintf("There is no device services mapping for this deviceId :%v and ServiceId :%v ",
		consumptionData.DeviceID, consumptionData.ServiceID)
		log.Errorft(ctx, msg)
		errMsg := errors.New(msg)
		return false, errMsg
	}
	return true, nil
}
