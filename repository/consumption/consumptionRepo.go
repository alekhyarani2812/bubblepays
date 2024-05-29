package consumption

import (
	"context"
	"encoding/csv"

	"os"
	"path/filepath"
	"sync"
	"time"

	bcontext "github.com/bubblepays/context"
	"github.com/bubblepays/log"
	"github.com/bubblepays/models"
	"github.com/bubblepays/utils"
)

const startConsumeDir = "/home/satyanarayana/Alekhya/gofilesdirectories/start_consume/"

// ConsumeDeviceData Get Consumption data from destination dirctory and process it
func ConsumeDeviceData(ctx context.Context) {
	var wg sync.WaitGroup
	// Open the directory
	f, err := os.Open(startConsumeDir)
	if err != nil {
		log.Infof("Error opening directory: %v", err)
		return
	}
	defer f.Close()
	// Read the directory contents
	files, err := f.Readdir(-1)
	if err != nil {
		log.Infof("Error reading directory contents: %v", err)
		return
	}
	// Iterate over the files and print their names
	for _, file := range files {
		log.Infof("fileName : %v ", file.Name())
		fileName := file.Name()
		records := ReadConsumptionData(startConsumeDir + fileName)
		log.Infof("records : %v ", records)
		// Create a wait group to wait for all goroutines to finish
		for _, record := range records {
			//convert startdate string to int
			startDate, err := utils.StringToTime(record[3])
			if err != nil {
				log.Infof("Got an ERROR at startdate: %v ", err.Error())
				continue
			}
			endDate, err := utils.StringToTime(record[4])
			if err != nil {
				log.Infof("Got an ERROR at enddate: %v ", err.Error())
				continue
			}
			createdDate := time.Now()
			consumptionData := models.Consumption{
				DeviceID:        record[0],
				CardID:          record[1],
				ConsumptionDate: record[2],
				StartTime:       startDate, //&time.Time{},
				EndTime:         endDate,   //&time.Time{},
				ServiceID:       record[5],
				CreatedDate:     &createdDate,
				CreatedBy:       "Alekhya",
			}
			dbConnection, err := bcontext.GetDBConnection(ctx)
			if err != nil {
				log.Errorft(ctx, "error in db connection %v ", err)
				return 
			}
			var serviceenrollment models.ServiceEnrollment
			sefRef := &serviceenrollment
			gorm := dbConnection.Where("card_id=?",consumptionData.CardID).
			Where("service_id=?",consumptionData.ServiceID).
			Find(&serviceenrollment)
			if gorm != nil && gorm.Error != nil {
				log.Errorft(ctx, "&&&&&&&&&&&&&error while getting user data")
				return 
			}
			//if sefRef == nil || 
			if sefRef.CardID == "" || sefRef.ServiceID == "" {
				log.Errorft(ctx, "There is no service enrollment for this cardId : %v and Service Id : %v ", consumptionData.CardID, consumptionData.ServiceID)
				return 
			}
			log.Infof("Data from db found serviceEnrolment : %v ", serviceenrollment)
			log.Info("calling SaveFileConsumption")
			wg.Add(1)
			go SaveFileConsumption(ctx, consumptionData, &wg)
		}
	}

	wg.Wait()
	// Delete the files
	for _, file1 := range files {
		err = os.Remove(startConsumeDir + file1.Name())
		if err != nil {
			log.Infof("Error deleting file: %v", err)
			continue
		}
		log.Infof("Deleted file: %v", file1)
	}
}

// ReadConsumptionData from file
func ReadConsumptionData(filePath string) (data [][]string) {
	log.Info("Read MachineData is called")
	// Open the CSV file
	file, err := os.Open(filePath)
	if err != nil {
		log.Info("Error: *****************")
		log.Infof("Error: %v", err)
		return
	}
	defer file.Close()
	// Create a new CSV reader
	reader := csv.NewReader(file)
	// Read all the records from the CSV file
	records, err := reader.ReadAll()
	if err != nil {
		log.Infof("Error: &&&&&&&&&&")
		log.Infof("Error: %v", err)
		return
	}
	log.Infof("records : %v ", records)
	return records
}

// SaveFileConsumption Insert consumption data into DB
func SaveFileConsumption(ctx context.Context, consumptionData models.Consumption,
	wg *sync.WaitGroup) error {
	//get dbconnection from contex
	dbConnection, err := bcontext.GetDBConnection(ctx)
	if err != nil {
		log.Errorft(ctx, "error in db connection %v ", err)
		return err
	}
	defer wg.Done()
	createdBy := "Alekhya"
	currentTime := time.Now()

	log.Infof("**************** Save consumption currentTime: %v", currentTime)
	ormDB := dbConnection.Create(&consumptionData)
	if ormDB.Error != nil {
		log.Infof("Error inserting consumption: %v ", ormDB.Error)
		return ormDB.Error
	}
	log.Info("consumption inserted successfully")
	log.Infof("ormDB  :  %v ", ormDB)
	log.Infof("Consumption '%s' created successfully\n", createdBy)
	return nil
}

// PickDeviceData Move Consumption data from one directory to another
func PickDeviceData() {
	// Define source and destination directories
	sourceDir := "/home/satyanarayana/Alekhya/gofilesdirectories/machine_data/"
	destDir := "/home/satyanarayana/Alekhya/gofilesdirectories/start_consume/"
	// Get list of files in the source directory
	files, err := filepath.Glob(filepath.Join(sourceDir, "*"))
	if err != nil {
		log.Infof("Error: %v ", err)
		return
	}
	// Move each file to the destination directory
	for _, file := range files {
		// Extract filename from path
		_, fileName := filepath.Split(file)
		// Generate destination path
		destPath := filepath.Join(destDir, fileName)
		// Move the file
		err := os.Rename(file, destPath)
		if err != nil {
			log.Infof("Error moving file: %v ", err)
		} else {
			log.Infof("Moved file: , to  %v %v ", file, destPath)
		}
	}
}
