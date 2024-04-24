package jobs

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
	"github.com/bubblepays/models"
	repository "github.com/bubblepays/repository/usages"
	"gorm.io/gorm"
)

// creation of usage table
func getConsumptionsByMonthnum(dbConnection *gorm.DB) (map[string]float64, error) {
	// Retrieve consumption data from the database
	var consumptions []models.Consumption
	if err := dbConnection.Find(&consumptions).Error; err != nil {
		log.Printf("failed to retrieve consumption data: %v", err)
		return nil, err
	}
	log.Printf("SIZE of consumption: size:: %v ", len(consumptions))
	cardIdServiceId := make(map[string]float64)
	for _, consum := range consumptions {
		layout := "2006-01-02T15:04:05-07:00"
		// Parse the time string into a time.Time object
		var monthNumStr string
		t, err := time.Parse(layout, consum.ConsumptionDate)
		if err != nil {
			fmt.Println("Error parsing time:", err)
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
		myKey := fmt.Sprintf("%v_%v_%v", consum.CardId, consum.ServiceId, monthnum)

		//duration = endDate-startDate
		duration := consum.EndTime.Sub(*consum.StartTime)

		if dur, ok := cardIdServiceId[myKey]; ok {
			fmt.Println("myKey ::::::::::: : ", myKey)
			fmt.Println("dur is ********** :", dur)
			fmt.Println("previous duration is : ", duration.Seconds())
			totalDuration := dur + duration.Seconds()
			cardIdServiceId[myKey] = totalDuration
			fmt.Println("totalDuration is : ", totalDuration)
		} else {
			totalDuration := duration.Seconds()
			cardIdServiceId[myKey] = totalDuration
		}
	}
	// Iterate over the map using a for loop and range
	for key, value := range cardIdServiceId {
		fmt.Printf("Key: %s, Value: %v\n", key, value)
	}
	return cardIdServiceId, nil
}

func UsageData(dbConnection *gorm.DB) {
	fmt.Println("UsageData is calling************")
	usageData, err := getConsumptionsByMonthnum(dbConnection)
	if err != nil {
		log.Printf("err %v ", err.Error())
		return
	}
	//fmt.Println("usageData : ", usageData)
	for key, duration := range usageData {
		fmt.Printf("Key: %s\n, Value: %v\n", key, duration)
		// Split the string by _
		keyStr := strings.Split(key, "_")
		cardId := keyStr[0]
		serviceId := keyStr[1]
		monthNum := keyStr[2]
		fmt.Println(cardId, serviceId, monthNum)
		usageData := models.Usages{
			CardId:    cardId,
			ServiceId: serviceId,
			Monthnum:  monthNum,
			Duration:  duration,
		}
		fmt.Println("usageData : ", usageData)
		repository.SaveUsage(usageData, dbConnection)
	}
}
