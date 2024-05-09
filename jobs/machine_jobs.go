package jobs

import (
	"context"
	"fmt"
	"time"
	"github.com/bubblepays/repository/consumption"
	"github.com/bubblepays/utils"
	"gorm.io/gorm"
)

func RunMachineJob(dbConnection *gorm.DB) {
	// Create a background context
	bctx := context.Background()
	ctx := utils.WithTransactionID(bctx)
	ctx = utils.WithDBConnection(ctx)
	// Execute the HTTP request every 10 seconds //new cron job
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println("this fuction is executing in every 10 seconds")
			consumption.ConsumeDeviceData(ctx)
			consumption.PickDeviceData()
			UsageData(dbConnection)
		}
	}
}
