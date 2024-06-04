package consumption

import (
	"context"

	bcontext "github.com/bubblepays/context"
	"github.com/bubblepays/log"
	"github.com/bubblepays/models"
)

// func JobData(ctx context.Context) {
// 	//GetJobData(filename string, ConsumptionDate string)
// 	jobData := models.Jobs{
// 		Filename:   "",
// 		Date:       "",
// 		JobMessage: "Success",
// 		Status:     "1",
// 	}
// 	fmt.Println("jobData : ", jobData)
// 	//SaveJobInfo(ctx, jobData)
// }
func SaveJobInfo(ctx context.Context, jobData *models.Jobs) error{
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
	log.Info("job inserted successfully")
	log.Infof("ormDB  :  %v ", ormDB)
	return nil
}
