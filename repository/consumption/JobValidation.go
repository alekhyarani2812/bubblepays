package consumption

import (
	"context"

	bcontext "github.com/bubblepays/context"
	"github.com/bubblepays/log"
	"github.com/bubblepays/models"
)

func JobValidation(ctx context.Context, consumptionData models.Consumption) bool {
	dbConnection, err := bcontext.GetDBConnection(ctx)
	if err != nil {
		log.Errorft(ctx, "error in db connection %v ", err)
		return false
	}
	// validate cardId, serviceId in service enrolment table
	var serviceenrollment models.ServiceEnrollment
	serviceenrolref := &serviceenrollment
	gorm := dbConnection.Where("card_id=?", consumptionData.CardID).
		Where("service_id=?", consumptionData.ServiceID).
		Find(&serviceenrolref)
	if gorm != nil && gorm.Error != nil {
		log.Errorft(ctx, "error while getting user data")
		return false
	}
	if serviceenrolref.CardID == "" || serviceenrolref.ServiceID == "" {
		log.Errorft(ctx, "There is no service enrollment for this cardId : %v and Service Id : %v ", consumptionData.CardID, consumptionData.ServiceID)
		return false
	}
	// validate deviceId, serviceId in deviceservices table
	var deviceservices models.DeviceServices
	deviceserviceref := &deviceservices
	gorms := dbConnection.Where("device_id=?", consumptionData.DeviceID).
		Where("service_id=?", consumptionData.ServiceID).
		Find(deviceserviceref)
	if gorms != nil && gorms.Error != nil {
		log.Errorft(ctx, "error while getting device data")
		return false
	}
	if deviceserviceref.DeviceID == "" || deviceserviceref.ServiceID == "" {
		log.Errorft(ctx, "There is no device services mapping for this deviceId :%v and ServiceId :%v ",
			consumptionData.DeviceID, consumptionData.ServiceID)
		return false
	}
	return true
}
