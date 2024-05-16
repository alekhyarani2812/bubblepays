package jobs

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bubblepays/log"
	"github.com/bubblepays/models"
	repository "github.com/bubblepays/repository/usages"
	"gorm.io/gorm"
)

func getConsumptionsByMonthnum(ctx context.Context, dbConnection *gorm.DB) (map[string]float64, error) {
	// Retrieve consumption data from the database
	var consumptions []models.Consumption
	if err := dbConnection.Find(&consumptions).Error; err != nil {
		log.Infof("failed to retrieve consumption data: %v", err)
		return nil, err
	}

	log.Infoft(ctx, "SIZE of consumption: size:: %v ", len(consumptions))
	cardIDServiceID := make(map[string]float64)
	for _, consum := range consumptions {
		layout := "2006-01-02T15:04:05-07:00"
		// Parse the time string into a time.Time object
		var monthNumStr string
		t, err := time.Parse(layout, consum.ConsumptionDate)
		if err != nil {
			log.Infof("Error parsing time: %v ", err)
		}
		year := t.Year()
		monthNumber := int(t.Month())
		if monthNumber < 10 {
			strMonth := strconv.Itoa(monthNumber)
			strZero := strconv.Itoa(0)
			monthNumStr = strZero + strMonth
		}
		// Print the parsed time
		monthnum := fmt.Sprintf("%v-%v", year, monthNumStr)
		myKey := fmt.Sprintf("%v_%v_%v", consum.CardID, consum.ServiceID, monthnum)

		//duration = endDate-startDate
		duration := consum.EndTime.Sub(*consum.StartTime)

		if dur, ok := cardIDServiceID[myKey]; ok {
			// log.Info( "previous duration is : ", duration.Seconds())
			totalDuration := dur + duration.Seconds()
			cardIDServiceID[myKey] = totalDuration
			//log.Info( "totalDuration is : ", totalDuration)
		} else {
			totalDuration := duration.Seconds()
			cardIDServiceID[myKey] = totalDuration
		}
	}
	// Iterate over the map using a for loop and range
	for key, value := range cardIDServiceID {
		log.Infoft(ctx, "Key: %s, Value: %v\n", key, value)
	}
	return cardIDServiceID, nil
}

func UsageData(ctx context.Context, dbConnection *gorm.DB) {
	usageData, err := getConsumptionsByMonthnum(ctx, dbConnection)
	if err != nil {
		log.Infof("err %v ", err.Error())
		return
	}
	for key, duration := range usageData {
		//fmt.Printf("Key: %s , Value: %v ", key, duration)
		// Split the string by _
		keyStr := strings.Split(key, "_")
		cardID := keyStr[0]
		serviceID := keyStr[1]
		monthNum := keyStr[2]
		log.Infoft(ctx, "cardID : %v, serviceID : %v, monthNum : %v ", cardID, serviceID, monthNum)
		usageData := models.Usages{
			CardID:    cardID,
			ServiceID: serviceID,
			Monthnum:  monthNum,
			Duration:  duration,
		}
		repository.SaveUsage(usageData, dbConnection)
	}
}
