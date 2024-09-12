package cronejobs

import (
	"context"
	"time"

	bcontext "github.com/bubblepays/context"
	"github.com/bubblepays/log"
	"github.com/bubblepays/repository/consumption"
	"gorm.io/gorm"
)

// RunMachineJob for running jobs
func RunMachineJob(dbConnection *gorm.DB) error {
	// Create a background context
	bctx := context.Background()
	ctx := bcontext.WithTransactionID(bctx)
	ctx = bcontext.WithDBConnection(ctx)
	// Execute the HTTP request every 10 seconds //new cron job
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			log.Infot(ctx, "this fuction is executing in every 30 seconds")
			// consuming device data
			err := consumption.PickDeviceData(ctx)
			if err != nil {
				log.Errorft(ctx, "The error is %v", err)
				return err
			}
			//Start consuming data, then save it into the Job and Consumption tables.
			err1 := consumption.SaveJobConsumptionData(ctx)
			if err1 != nil {
				log.Errorft(ctx, "The error is %v", err1)
				return err1
			}
			usageDataError := UsageData(ctx, dbConnection)
			if usageDataError != nil {
				log.Errorft(ctx, "Failed to get usageData %v", usageDataError)
				return usageDataError
			}
			log.Infot(ctx, "saved Job and consumtion data successfully")
		}
	}
}
