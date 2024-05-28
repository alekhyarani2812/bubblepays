package auth

import (
	"context"

	bcontext "github.com/bubblepays/context"
	"github.com/bubblepays/log"
	"github.com/bubblepays/models"
)

func Login(ctx context.Context, userPayload models.User) (*models.User, error) {
	dbConnection, err := bcontext.GetDBConnection(ctx)
	if err != nil {
		log.Infof("error in db connection %v", err.Error())
		return nil, err
	}
	log.Infof("payload useraname , password  %v, %v", userPayload.Username, userPayload.Password)
	var userData []*models.User
	gorm := dbConnection.Where("user_name = ?", userPayload.Username).Where(
		"password=?", userPayload.Password).Find(&userData)
	if gorm != nil && gorm.Error != nil {
		log.Errorf("error while getting user data")
		return nil, gorm.Error
	}
	log.Infof("Data from db found userData : %v ", userData[0])
	return userData[0], nil
}
