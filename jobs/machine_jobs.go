package jobs

import (
	"context"
	"time"

	bcontext "github.com/bubblepays/context"
	"github.com/bubblepays/repository/consumption"
	"github.com/bubblepays/log"
	"gorm.io/gorm"
)

func RunMachineJob(dbConnection *gorm.DB) {
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
			consumption.PickDeviceData()
			consumption.ConsumeDeviceData(ctx)
			//UsageData(ctx, dbConnection)
		}
	}
}
