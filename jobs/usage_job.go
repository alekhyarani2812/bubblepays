package jobs

import (
	"fmt"
	"log"
	"github.com/bubblepays/database"
	"github.com/bubblepays/models"
)

func usage() (error){
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Printf("err %v ", err.Error())
		return err
	}
	// Retrieve consumption data from the database
	var consumptions []models.Consumption
	if err := dbConnection.Find(&consumptions).Error; err != nil {
		log.Printf("failed to retrieve consumption data: %v", err)
		return err
	}
	log.Printf("SIZE of consumption: size:: %v ", len(consumptions))
	// Iterate over the result set and print the records
    for _, consum := range consumptions {
        fmt.Printf("ID: %v, CardId:%v, ServiceId:%v, StartTime: %v, EndTime: %v\n", consum.Id, consum.CardId, consum.ServiceId, consum.StartTime, consum.EndTime)
		duration := consum.EndTime.Sub(*consum.StartTime)
		fmt.Println("duration :", duration)
		//duration.Seconds()
    }
	
	return nil
}