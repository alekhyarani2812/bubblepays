package consumption

import (
	"encoding/csv"
	"fmt"
	"github.com/bubblepays/database"
	"github.com/bubblepays/models"
	"github.com/bubblepays/utils"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const startConsumeDir = "/home/satyanarayana/Alekhya/gofilesdirectories/start_consume/"

// Read consumption data from file
func ReadConsumptionData(filePath string) (data [][]string) {
	fmt.Println("Read MachineData is called")
	// Open the CSV file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error: *****************")
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()
	// Create a new CSV reader
	reader := csv.NewReader(file)
	// Read all the records from the CSV file
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error: &&&&&&&&&&")
		fmt.Println("Error:", err)
		return
	}
	return records
}

// Insert consumption data into DB
func SaveFileConsumption(consumptionData models.Consumption, wg *sync.WaitGroup) error {
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Printf("err %v ", err.Error())
		return err
	}
	defer wg.Done()
	createdBy := "Alekhya"
	currentTime := time.Now()

	fmt.Printf("**************** Save consumption currentTime: %v", currentTime)
	ormDB := dbConnection.Create(&consumptionData)
	if ormDB.Error != nil {
		log.Println("Error inserting consumption:", ormDB.Error)
		return ormDB.Error
	}
	log.Println("consumption inserted successfully")
	log.Println("ormDB  :  ", ormDB)
	fmt.Printf("Consumption '%s' created successfully\n", createdBy)
	return nil
}

// Get Consumption data from destination dirctory and process it
func ConsumeDeviceData() {
	var wg sync.WaitGroup
	// Open the directory
	f, err := os.Open(startConsumeDir)
	if err != nil {
		fmt.Println("Error opening directory:", err)
		return
	}
	defer f.Close()
	// Read the directory contents
	files, err := f.Readdir(-1)
	if err != nil {
		fmt.Println("Error reading directory contents:", err)
		return
	}
	// Iterate over the files and print their names
	for _, file := range files {
		fmt.Println("fileName : ", file.Name())
		fileName := file.Name()
		records := ReadConsumptionData(startConsumeDir + fileName)
		fmt.Println("records :", records)
		// Create a wait group to wait for all goroutines to finish

		for _, record := range records {
			//convert cardId string to int
			cardId, err := utils.StringToInt(record[1])
			if err != nil {
				log.Println(err.Error())
				continue
			}
			// consumptionDate, err := utils.StringToInt(record[2])
			// if err != nil {
			// 	log.Println("Got an ERROR at consumptionDate: ", err.Error())
			// 	continue
			// }
			serviceId, err := utils.StringToInt(record[5])
			if err != nil {
				log.Println("Got an ERROR at serviceId: ", err.Error())
				continue
			}
			//convert startdate string to int
			startDate, err := utils.StringToTime(record[3])
			if err != nil {
				log.Println("Got an ERROR at startdate: ", err.Error())
				continue
			}
			endDate, err := utils.StringToTime(record[4])
			if err != nil {
				log.Println("Got an ERROR at enddate: ", err.Error())
				continue
			}
			createdDate := time.Now()
			consumptionData := models.Consumption{
				DeviceId:        record[0],
				CardId:          cardId,
				ConsumptionDate: record[2],
				StartTime:       startDate, //&time.Time{},
				EndTime:         endDate,   //&time.Time{},
				ServiceId:       serviceId,
				ServiceName:     record[6],
				CreatedDate:     &createdDate,
				CreatedBy:       "Alekhya",
			}
			log.Println("calling SaveFileConsumption")
			wg.Add(1)
			go SaveFileConsumption(consumptionData, &wg)
		}
	}

	wg.Wait()
	// Delete the files
	for _, file1 := range files {
		err = os.Remove(startConsumeDir + file1.Name())
		if err != nil {
			fmt.Println("Error deleting file:", err)
			continue
		}
		fmt.Println("Deleted file:", file1)
	}
}

// Move Consumption data from one directory to another
func PickDeviceData() {
	// Define source and destination directories
	sourceDir := "/home/satyanarayana/Alekhya/gofilesdirectories/machine_data/"
	destDir := "/home/satyanarayana/Alekhya/gofilesdirectories/start_consume/"
	// Get list of files in the source directory
	files, err := filepath.Glob(filepath.Join(sourceDir, "*"))
	//fmt.Println("files :", files)
	if err != nil {
		fmt.Println("Error:", err)
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
			fmt.Println("Error moving file:", err)
		} else {
			fmt.Println("Moved file:", file, "to", destPath)
		}
	}
}
