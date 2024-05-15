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

// Read consumption data from file
func ReadConsumptionData(filePath string) (data [][]string) {
	log.Infof("Read MachineData is called")
	// Open the CSV file
	file, err := os.Open(filePath)
	if err != nil {
		log.Infof("Error: *****************")
		log.Infof("Error:", err)
		return
	}
	defer file.Close()
	// Create a new CSV reader
	reader := csv.NewReader(file)
	// Read all the records from the CSV file
	records, err := reader.ReadAll()
	if err != nil {
		log.Infof("Error: &&&&&&&&&&")
		log.Infof("Error:", err)
		return
	}
	log.Infof("records : %v ", records)
	return records
}

// Insert consumption data into DB
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

// Get Consumption data from destination dirctory and process it
func ConsumeDeviceData(ctx context.Context) {
	var wg sync.WaitGroup
	// Open the directory
	f, err := os.Open(startConsumeDir)
	if err != nil {
		log.Infof("Error opening directory:", err)
		return
	}
	defer f.Close()
	// Read the directory contents
	files, err := f.Readdir(-1)
	if err != nil {
		log.Infof("Error reading directory contents:", err)
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
			serviceId, err := utils.StringToInt(record[5])
			if err != nil {
				log.Infof("Got an ERROR at serviceId: %v ", err.Error())
				continue
			}
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
				DeviceId:        record[0],
				CardId:          record[1],
				ConsumptionDate: record[2],
				StartTime:       startDate, //&time.Time{},
				EndTime:         endDate,   //&time.Time{},
				ServiceId:       serviceId,
				ServiceName:     record[6],
				CreatedDate:     &createdDate,
				CreatedBy:       "Alekhya",
			}
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
			log.Infof("Error deleting file:", err)
			continue
		}
		log.Infof("Deleted file:", file1)
	}
}

// Move Consumption data from one directory to another
func PickDeviceData() {
	// Define source and destination directories
	sourceDir := "/home/satyanarayana/Alekhya/gofilesdirectories/machine_data/"
	destDir := "/home/satyanarayana/Alekhya/gofilesdirectories/start_consume/"
	// Get list of files in the source directory
	files, err := filepath.Glob(filepath.Join(sourceDir, "*"))
	if err != nil {
		log.Infof("Error:", err)
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
			log.Infof("Error moving file:", err)
		} else {
			log.Infof("Moved file:", file, "to", destPath)
		}
	}
}
