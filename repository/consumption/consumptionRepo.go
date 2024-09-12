package consumption

import (
	"context"
	"encoding/csv"
	"errors"

	"os"
	"time"

	bcontext "github.com/bubblepays/context"
	"github.com/bubblepays/log"
	"github.com/bubblepays/models"
	"github.com/bubblepays/utils"
)

const startConsumeDir = "/home/satyanarayana/Alekhya/gofilesdirectories/start_consume/"
const consumeSuccessDir = "/home/satyanarayana/Alekhya/gofilesdirectories/consume_success/"
const consumeFailureDir = "/home/satyanarayana/Alekhya/gofilesdirectories/consume_failure/"

// SaveJobConsumptionData Start consuming data, then save it into the Job and Consumption tables.
func SaveJobConsumptionData(ctx context.Context) error {
	dbConnection, err := bcontext.GetDBConnection(ctx)
	if err != nil {
		log.Errorft(ctx, "error in db connection %v ", err)
		return err
	}
	var jobs []models.Jobs
	dbConnection.Where("status=?", "IN_PROGRESS").Find(&jobs)
	log.Infoft(ctx, "in_progress jobs are : %v\n", jobs)
	for _, job := range jobs {
		//var jobData models.Jobs
		filename := job.Filename
		log.Infof("file name from DB job table %v", filename)
		// ConsumeDeviceData from destination dirctory and process it
		records, err := ReadConsumptionData(startConsumeDir + filename)
		if err != nil {
			log.Errorft(ctx, "Error while getting Records %v", err)
			return err
		}
		log.Infof("*****************records : %v ", records)
		for _, record := range records {
			//convert startdate string to time
			startDate, err := utils.StringToTime(record[3])
			if err != nil {
				log.Infof("Got an ERROR at startdate: %v ", err.Error())
				return err
			}
			endDate, err := utils.StringToTime(record[4])
			if err != nil {
				log.Errorft(ctx, "Got an ERROR at enddate: %v ", err.Error())
				return err
			}
			//Preparing consumption data
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
			//Validate deviceId as string
			if !(utils.IsAlphanumeric(consumptionData.DeviceID)) {
				log.Infoft(ctx, "This is invalid deviceid %v", consumptionData.DeviceID)
				return errors.New("invalid userid")
			}
			//Validate cardId as string
			if !(utils.IsAlphanumeric(consumptionData.CardID)) {
				log.Infoft(ctx, "This is invalid cardId %v", consumptionData.CardID)
				return errors.New("invalid CardID")
			}
			////Validate ServiceID as string
			if !(utils.IsAlphanumeric(consumptionData.ServiceID)) {
				log.Infoft(ctx, "This is invalid ServiceID %v", consumptionData.ServiceID)
				return errors.New("invalid ServiceID")
			}
			//validate cardId, serviceId in service enrolment table
			boole, err := JobValidation(ctx, consumptionData)
			if err != nil {
				log.Infof("Got an ERROR in Job validation: %v ", err.Error())
				return err
			}
			if boole {
				//validation is success
				log.Infot(ctx, "Job validation done. Now calling SaveFileConsumption")
				err := SaveFileConsumption(ctx, consumptionData)
				// If we get error while saving consumption data
				if err != nil {
					log.Errorft(ctx, "Error while saving consumption data %v", err)
					return err
				}
				// consumption data has been saved
				log.Infot(ctx, "consumption data has been saved in consumption table")
				// Preparing job data
				jobData := models.Jobs{
					JobMessage: "Job completed",
					Status:     utils.JobSuccess,
				}
				// If we get an error while updating job data
				err1 := UpdateJobInfo(ctx, &jobData, filename)
				if err1 != nil {
					log.Errorft(ctx, "Error while updating Job data %v", err1)
					return err1
				}
				// Job data updated successfully
				log.Infot(ctx, "Job data updated successfully")
				//Move file to destination dir
				err2 := MoveTOSuccessDir(filename)
				if err2 != nil {
					// Error while while moving file
					log.Errorft(ctx, "######### Error while moving file :%v from source dir to destination dir", filename)
					return err2
				}
				log.Infoft(ctx, "######### successfully file %v moved from %v to %v ", filename, consumeSuccessDir, consumeFailureDir)
			} else {
				// If validation fails
				jobData := models.Jobs{
					JobMessage: "Incorrect Job data",
					Status:     utils.JobFailed,
				}
				err := UpdateJobInfo(ctx, &jobData, filename)
				if err != nil {
					log.Errorft(ctx, "Error while updating Job data %v", err)
					return err
				}
				err2 := MoveTOFailureDir(filename)
				if err2 != nil {
					// Error while while moving file
					log.Errorft(ctx, "######### If got Error in validation the file :%v moves from consume dir to failure dir", filename)
					return err2
				}
				log.Infoft(ctx, "######### successfully file %v moved from %v to %v ", filename, consumeSuccessDir, consumeFailureDir)
			}
		}
	}
	return nil
}

// MoveTOSuccessDir Move file from consume dir to success dir
func MoveTOSuccessDir(filename string) error {
	sourcePath := startConsumeDir + filename
	destinationPath := consumeSuccessDir + filename
	err := os.Rename(sourcePath, destinationPath)
	if err != nil {
		log.Infof("Error moving the file: %v", err)
		return err
	}
	log.Infof("File moved successfully %v", filename)
	return nil
}

// MoveTOFailureDir Move file from consume dir to failure dir
func MoveTOFailureDir(filename string) error {
	sourcePath := startConsumeDir + filename
	destinationPath := consumeFailureDir + filename
	err := os.Rename(sourcePath, destinationPath)
	if err != nil {
		log.Infof("Error moving the file: %v", err)
		return err
	}
	log.Infof("File moved successfully %v", filename)
	return nil
}

// ReadConsumptionData from complete path of single file
func ReadConsumptionData(filePath string) (data [][]string, err error) {
	log.Info("Read MachineData is called")
	// Open the CSV file
	file, err := os.Open(filePath)
	if err != nil {
		log.Errorf("Error: %v", err)
		return nil, err
	}
	defer file.Close()
	// Create a new CSV reader
	reader := csv.NewReader(file)
	// Read all the records from the CSV file
	records, err := reader.ReadAll()
	if err != nil {
		log.Errorf("Error: %v", err)
		return nil, err
	}
	log.Infof("records : %v ", records)
	return records, nil
}

// SaveFileConsumption Insert consumption data into DB
func SaveFileConsumption(ctx context.Context, consumptionData models.Consumption) error {
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
