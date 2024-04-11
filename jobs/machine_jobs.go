package jobs

import (
	"fmt"
	"time"

	"github.com/bubblepays/repository/consumption"
	"gorm.io/gorm"
)

func RunMachineJob(dbConnection *gorm.DB) {
	// Execute the HTTP request every 10 seconds //new cron job
	ticker := time.NewTicker(2 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println("this fuction is executing in every 10 seconds")
			consumption.ConsumeDeviceData()
			consumption.PickDeviceData()
			//usage()
			//UsageData(dbConnection)
		}
	}
}
