package consumption

import (
	"context"
	"encoding/csv"

	"os"
	"sync"
	"time"

	bcontext "github.com/bubblepays/context"
	"github.com/bubblepays/log"
	"github.com/bubblepays/models"
	"github.com/bubblepays/utils"
)

const startConsumeDir = "/home/satyanarayana/Alekhya/gofilesdirectories/start_consume/"

func GetJobRecords(ctx context.Context) {
	var wg sync.WaitGroup
	dbConnection, err := bcontext.GetDBConnection(ctx)
	if err != nil {
		log.Errorft(ctx, "error in db connection %v ", err)
		return
	}
	var jobs []models.Jobs
	dbConnection.Where("status=?", "IN_PROGRESS").Find(&jobs)
	log.Infoft(ctx, "in_progress jobs are : %v\n", jobs)
	for _, job := range jobs {
		filename := job.Filename
		log.Infof("file name from DB job table %v", filename)
		// ConsumeDeviceData from destination dirctory and process it
		records := ReadConsumptionData(startConsumeDir + filename)
		log.Infof("records : %v ", records)
		for _, record := range records {
			//convert startdate string to int
			startDate, err := utils.StringToTime(record[3])
			if err != nil {
				log.Infof("Got an ERROR at startdate: %v ", err.Error())
				return
			}
			endDate, err := utils.StringToTime(record[4])
			if err != nil {
				log.Errorft(ctx, "Got an ERROR at enddate: %v ", err.Error())
				return
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
			jobdata := models.Jobs{
				JobMessage: "Job completed",
				Status:     utils.JobSuccess,
			}
			//validate cardId, serviceId in service enrolment table
			bool := JobValidation(ctx, consumptionData)
			if !bool {
				jobdata := models.Jobs{
					JobMessage: "Incorrect Job data",
					Status:     utils.JobFailed,
				}
				wg.Add(1)
				go UpdateJobInfo(ctx, &jobdata, filename, &wg)
				return
			}
			jobdata.Status = utils.JobSuccess
			wg.Add(1)
			go UpdateJobInfo(ctx, &jobdata, filename, &wg)
			wg.Add(1)
			log.Info("calling SaveFileConsumption")
			go SaveFileConsumption(ctx, consumptionData, &wg)
		}
		wg.Wait()
		RemoveFile(filename)
	}
}
func RemoveFile(filename string) {
	var err error
	err = os.Remove(startConsumeDir + filename)
	if err != nil {
		log.Infof("Error deleting file: %v", err)
		return
	}
	log.Infof("Deleted file: from startConsumeDir %v", filename)
}

// ReadConsumptionData from file
func ReadConsumptionData(filePath string) (data [][]string) {
	log.Info("Read MachineData is called")
	// Open the CSV file
	file, err := os.Open(filePath)
	if err != nil {
		log.Errorf("Error: %v", err)
		return nil
	}
	defer file.Close()
	// Create a new CSV reader
	reader := csv.NewReader(file)
	// Read all the records from the CSV file
	records, err := reader.ReadAll()
	if err != nil {
		log.Errorf("Error: %v", err)
		return nil
	}
	log.Infof("records : %v ", records)
	return records
}

// SaveFileConsumption Insert consumption data into DB
func SaveFileConsumption(ctx context.Context, consumptionData models.Consumption,
	wg *sync.WaitGroup) error {
	defer wg.Done()
	//get dbconnection from contex
	dbConnection, err := bcontext.GetDBConnection(ctx)
	if err != nil {
		log.Errorft(ctx, "error in db connection %v ", err)
		return err
	}
	createdBy := "Alekhya"
	currentTime := time.Now()
	log.Infof("**************** Save consumption currentTime: %v", currentTime)
	ormDB := dbConnection.Create(&consumptionData)
	if ormDB.Error != nil {
		log.Infof("Error inserting consumption: %v ", ormDB.Error)
		return ormDB.Error
	}
	log.Infof("ormDB consumption inserted successfully :  %v ", ormDB)
	log.Infof("Consumption '%s' created successfully\n", createdBy)
	return nil
}

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
				log.Errorft(ctx, "Got an ERROR at enddate: %v ", err.Error())
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
			jobdata := models.Jobs{
				ID:         8,
				JobMessage: "Job completed",
				Status:     utils.JobSuccess,
			}
			//validate cardId, serviceId in service enrolment table
			bool := JobValidation(ctx, consumptionData)
			if !bool {
				jobdata.Status = utils.JobFailed
				UpdateJobInfo(ctx, &jobdata, fileName, &wg)
				return
			}
			jobdata.Status = utils.JobSuccess
			wg.Add(1)
			go UpdateJobInfo(ctx, &jobdata, fileName, &wg)
			wg.Add(1)
			log.Info("calling SaveFileConsumption")
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
