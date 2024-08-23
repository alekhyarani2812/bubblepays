package consumption

import (
	"context"

	bcontext "github.com/bubblepays/context"
	"github.com/bubblepays/log"
	"github.com/bubblepays/models"
)

func SaveJobInfo(ctx context.Context, jobData *models.Jobs) error {
	//get dbconnection from contex
	dbConnection, err := bcontext.GetDBConnection(ctx)
	if err != nil {
		log.Errorft(ctx, "error in db connection %v ", err)
		return err
	}
	log.Infof("**************** Save job info")
	ormDB := dbConnection.Create(jobData)
	if ormDB.Error != nil {
		log.Infof("Error inserting job info: %v ", ormDB.Error)
		return ormDB.Error
	}
	log.Infof("job inserted successfully  ormDB : %v ", ormDB)
	return nil
}

func UpdateJobInfo(ctx context.Context, jobData *models.Jobs, filename string) error {
	//get dbconnection from contex
	dbConnection, err := bcontext.GetDBConnection(ctx)
	if err != nil {
		log.Errorft(ctx, "error in db connection %v ", err)
		return err
	}
	log.Infof("*************Update job info************")
	jobRecord := &models.Jobs{}
	ormDB := dbConnection.Model(jobRecord).Where("file_name = ?", filename).Updates(&jobData)
	if ormDB.Error != nil {
		log.Errorft(ctx, "Error Update job info: %v\n ", ormDB.Error)
		return ormDB.Error
	} else {
		log.Info("job Updated successfully")
		return nil
	}
}
